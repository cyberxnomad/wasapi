package mmdevice

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/cyberxnomad/wasapi/com"
	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("D666063F-1587-4E43-81F1-B948E807363F")
var _IID_IMMDevice = windows.GUID{Data1: 0xD666063F, Data2: 0x1587, Data3: 0x4E43, Data4: [8]byte{0x81, 0xF1, 0xB9, 0x48, 0xE8, 0x07, 0x36, 0x3F}}

func IID_IMMDevice() windows.GUID {
	return _IID_IMMDevice
}

type IMMDevice struct {
	vtbl *_IMMDeviceVtbl
}

type _IMMDeviceVtbl struct {
	com.IUnknownVtbl
	Activate          uintptr
	OpenPropertyStore uintptr
	GetId             uintptr
	GetState          uintptr
}

func (device *IMMDevice) Release() (err error) {
	r, _, _ := syscall.SyscallN(device.vtbl.Release, uintptr(unsafe.Pointer(device)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDevice::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// Activate 方法创建具有指定接口的 COM 对象。
// 不需要时需主动调用 Release 方法。
func (device *IMMDevice) Activate(iid windows.GUID, clsCtx uint32, activationParams *com.PROPVARIANT) (ppInterface unsafe.Pointer, err error) {
	r, _, _ := syscall.SyscallN(device.vtbl.Activate, uintptr(unsafe.Pointer(device)),
		uintptr(unsafe.Pointer(&iid)),
		uintptr(clsCtx),
		uintptr(unsafe.Pointer(activationParams)),
		uintptr(unsafe.Pointer(&ppInterface)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDevice::Activate failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetId 方法检索标识音频终结点设备的终结点 ID 字符串。
func (device *IMMDevice) GetId() (id string, err error) {
	var utf16ptr *uint16
	r, _, _ := syscall.SyscallN(device.vtbl.GetId, uintptr(unsafe.Pointer(device)),
		uintptr(unsafe.Pointer(&utf16ptr)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDevice::GetId failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	id = windows.UTF16PtrToString(utf16ptr)

	// 释放存储
	com.CoTaskMemFree(unsafe.Pointer(utf16ptr))

	return
}

// GetState 方法检索当前设备状态。
func (device *IMMDevice) GetState() (state uint32, err error) {
	r, _, _ := syscall.SyscallN(device.vtbl.GetState, uintptr(unsafe.Pointer(device)),
		uintptr(unsafe.Pointer(&state)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDevice::GetState failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// OpenPropertyStore 方法检索设备属性存储的接口。
// 不需要时需主动调用 Release 方法。
func (device *IMMDevice) OpenPropertyStore(stgmAccess uint32) (properties *com.IPropertyStore, err error) {
	r, _, _ := syscall.SyscallN(device.vtbl.OpenPropertyStore, uintptr(unsafe.Pointer(device)),
		uintptr(stgmAccess),
		uintptr(unsafe.Pointer(&properties)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDevice::OpenPropertyStore failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

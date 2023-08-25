package mmdevice

import (
	"fmt"
	"syscall"
	"unsafe"
	"wasapi/com"

	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("A95664D2-9614-4F35-A746-DE8DB63617E6")
var _IID_IMMDeviceEnumerator = windows.GUID{Data1: 0xA95664D2, Data2: 0x9614, Data3: 0x4F35, Data4: [8]byte{0xA7, 0x46, 0xDE, 0x8D, 0xB6, 0x36, 0x17, 0xE6}}

func IID_IMMDeviceEnumerator() windows.GUID {
	return _IID_IMMDeviceEnumerator
}

type IMMDeviceEnumerator struct {
	vtbl *_IMMDeviceEnumeratorVtbl
}

type _IMMDeviceEnumeratorVtbl struct {
	com.IUnknownVtbl
	EnumAudioEndpoints                     uintptr
	GetDefaultAudioEndpoint                uintptr
	GetDevice                              uintptr
	RegisterEndpointNotificationCallback   uintptr
	UnregisterEndpointNotificationCallback uintptr
}

func (self *IMMDeviceEnumerator) Release() (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Release, uintptr(unsafe.Pointer(self)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDeviceEnumerator::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// EnumAudioEndpoints 方法生成符合指定条件的音频终结点设备的集合。
// 不需要时需主动调用 Release 方法。
func (self *IMMDeviceEnumerator) EnumAudioEndpoints(dataFlow EDataFlow, stateMask uint32) (devices *IMMDeviceCollection, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.EnumAudioEndpoints, uintptr(unsafe.Pointer(self)),
		uintptr(dataFlow),
		uintptr(stateMask),
		uintptr(unsafe.Pointer(&devices)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDeviceEnumerator::EnumAudioEndpoints failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetDefaultAudioEndpoint 方法检索指定数据流方向和角色的默认音频终结点。
func (self *IMMDeviceEnumerator) GetDefaultAudioEndpoint(dataFlow EDataFlow, role ERole) (endpoint *IMMDevice, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetDefaultAudioEndpoint, uintptr(unsafe.Pointer(self)),
		uintptr(dataFlow),
		uintptr(role),
		uintptr(unsafe.Pointer(&endpoint)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDeviceEnumerator::GetDefaultAudioEndpoint failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetDevice 方法检索由终结点 ID 字符串标识的音频终结点设备。
func (self *IMMDeviceEnumerator) GetDevice(id string) (device *IMMDevice, err error) {
	utf16ptr := windows.StringToUTF16Ptr(id)

	r, _, _ := syscall.SyscallN(self.vtbl.GetDevice, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(utf16ptr)),
		uintptr(unsafe.Pointer(&device)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDeviceEnumerator::GetDevice failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// RegisterEndpointNotificationCallback 方法注册客户端的通知回调接口。
func (self *IMMDeviceEnumerator) RegisterEndpointNotificationCallback(client *IMMNotificationClient) (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetDevice, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(client)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDeviceEnumerator::RegisterEndpointNotificationCallback failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// UnregisterEndpointNotificationCallback 方法删除客户端在对
// IMMDeviceEnumerator::RegisterEndpointNotificationCallback 方法的调用中注册的通知接口的注册。
func (self *IMMDeviceEnumerator) UnregisterEndpointNotificationCallback(client *IMMNotificationClient) (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetDevice, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(client)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDeviceEnumerator::UnregisterEndpointNotificationCallback failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

package mmdevice

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"
	"wasapi/com"

	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("7991EEC9-7E89-4D85-8390-6C703CEC60C0")
var _IID_IMMNotificationClient = windows.GUID{Data1: 0x7991EEC9, Data2: 0x7E89, Data3: 0x4D85, Data4: [8]byte{0x83, 0x90, 0x6C, 0x70, 0x3C, 0xEC, 0x60, 0xC0}}

func IID_IMMNotificationClient() windows.GUID {
	return _IID_IMMNotificationClient
}

type IMMNotificationClient struct {
	vtbl *_IMMNotificationClientVtbl
}

type _IMMNotificationClientVtbl struct {
	com.IUnknownVtbl
	OnDeviceStateChanged   uintptr
	OnDeviceAdded          uintptr
	OnDeviceRemoved        uintptr
	OnDefaultDeviceChanged uintptr
	OnPropertyValueChanged uintptr
}

func (self *IMMNotificationClient) Release() (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Release, uintptr(unsafe.Pointer(self)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDevice::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// OnDefaultDeviceChanged 方法通知客户端特定设备的默认音频终结点设备已更改。
func (self *IMMNotificationClient) OnDefaultDeviceChanged(flow EDataFlow, role ERole, defaultDeviceId string) (err error) {
	utf16ptr := windows.StringToUTF16Ptr(defaultDeviceId)
	r, _, _ := syscall.SyscallN(self.vtbl.OnDefaultDeviceChanged, uintptr(unsafe.Pointer(self)),
		uintptr(flow),
		uintptr(role),
		uintptr(unsafe.Pointer(utf16ptr)),
	)
	runtime.KeepAlive(utf16ptr)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDevice::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// OnDeviceAdded 方法指示已添加新的音频终结点设备。
func (self *IMMNotificationClient) OnDeviceAdded(deviceId string) (err error) {
	utf16ptr := windows.StringToUTF16Ptr(deviceId)
	r, _, _ := syscall.SyscallN(self.vtbl.OnDefaultDeviceChanged, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(utf16ptr)),
	)
	runtime.KeepAlive(utf16ptr)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDevice::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// OnDeviceRemoved 方法指示已删除音频终结点设备。
func (self *IMMNotificationClient) OnDeviceRemoved(deviceId string) (err error) {
	utf16ptr := windows.StringToUTF16Ptr(deviceId)
	r, _, _ := syscall.SyscallN(self.vtbl.OnDefaultDeviceChanged, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(utf16ptr)),
	)
	runtime.KeepAlive(utf16ptr)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDevice::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// OnDeviceStateChanged 方法指示音频终结点设备的状态已更改。
func (self *IMMNotificationClient) OnDeviceStateChanged(deviceId string, newState uint32) (err error) {
	utf16ptr := windows.StringToUTF16Ptr(deviceId)
	r, _, _ := syscall.SyscallN(self.vtbl.OnDefaultDeviceChanged, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(utf16ptr)),
		uintptr(newState),
	)
	runtime.KeepAlive(utf16ptr)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDevice::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// OnPropertyValueChanged 方法指示属于音频终结点设备的属性的值已更改。
func (self *IMMNotificationClient) OnPropertyValueChanged(deviceId string, key com.PROPERTYKEY) (err error) {
	utf16ptr := windows.StringToUTF16Ptr(deviceId)
	r, _, _ := syscall.SyscallN(self.vtbl.OnDefaultDeviceChanged, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(utf16ptr)),
		uintptr(unsafe.Pointer(&key)),
	)
	runtime.KeepAlive(utf16ptr)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDevice::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

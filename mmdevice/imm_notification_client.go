package mmdevice

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/xray-bit/wasapi/com"
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

func (client *IMMNotificationClient) Release() (err error) {
	r, _, _ := syscall.SyscallN(client.vtbl.Release, uintptr(unsafe.Pointer(client)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDevice::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// OnDefaultDeviceChanged 方法通知客户端特定设备的默认音频终结点设备已更改。
func (client *IMMNotificationClient) OnDefaultDeviceChanged(flow EDataFlow, role ERole, defaultDeviceId string) (err error) {
	var utf16ptr *uint16

	if utf16ptr, err = windows.UTF16PtrFromString(defaultDeviceId); err != nil {
		return
	}

	r, _, _ := syscall.SyscallN(client.vtbl.OnDefaultDeviceChanged, uintptr(unsafe.Pointer(client)),
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
func (client *IMMNotificationClient) OnDeviceAdded(deviceId string) (err error) {
	var utf16ptr *uint16

	if utf16ptr, err = windows.UTF16PtrFromString(deviceId); err != nil {
		return
	}

	r, _, _ := syscall.SyscallN(client.vtbl.OnDefaultDeviceChanged, uintptr(unsafe.Pointer(client)),
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
func (client *IMMNotificationClient) OnDeviceRemoved(deviceId string) (err error) {
	var utf16ptr *uint16

	if utf16ptr, err = windows.UTF16PtrFromString(deviceId); err != nil {
		return
	}

	r, _, _ := syscall.SyscallN(client.vtbl.OnDefaultDeviceChanged, uintptr(unsafe.Pointer(client)),
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
func (client *IMMNotificationClient) OnDeviceStateChanged(deviceId string, newState uint32) (err error) {
	var utf16ptr *uint16

	if utf16ptr, err = windows.UTF16PtrFromString(deviceId); err != nil {
		return
	}

	r, _, _ := syscall.SyscallN(client.vtbl.OnDefaultDeviceChanged, uintptr(unsafe.Pointer(client)),
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
func (client *IMMNotificationClient) OnPropertyValueChanged(deviceId string, key com.PROPERTYKEY) (err error) {
	var utf16ptr *uint16

	if utf16ptr, err = windows.UTF16PtrFromString(deviceId); err != nil {
		return
	}

	r, _, _ := syscall.SyscallN(client.vtbl.OnDefaultDeviceChanged, uintptr(unsafe.Pointer(client)),
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

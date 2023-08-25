package mmdevice

import (
	"fmt"
	"syscall"
	"unsafe"
	"wasapi/com"

	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("0BD7A1BE-7A1A-44DB-8397-CC5392387B5E")
var _IID_IMMDeviceCollection = windows.GUID{Data1: 0x0BD7A1BE, Data2: 0x7A1A, Data3: 0x8397, Data4: [8]byte{0x83, 0x97, 0xCC, 0x53, 0x92, 0x38, 0x7B, 0x5E}}

func IID_IMMDeviceCollection() windows.GUID {
	return _IID_IMMDeviceCollection
}

type IMMDeviceCollection struct {
	vtbl *_IMMDeviceCollectionVtbl
}

type _IMMDeviceCollectionVtbl struct {
	com.IUnknownVtbl
	GetCount uintptr
	Item     uintptr
}

func (self *IMMDeviceCollection) Release() (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Release, uintptr(unsafe.Pointer(self)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDeviceCollection::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetCount 方法检索设备集合中的设备计数。
func (self *IMMDeviceCollection) GetCount() (pcDevices uint, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetCount, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&pcDevices)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDeviceCollection::GetCount failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// Item 方法检索指向设备集合中指定项的指针。
// 不需要时需主动调用 Release 方法。
func (self *IMMDeviceCollection) Item(nDevice uint32) (device *IMMDevice, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Item, uintptr(unsafe.Pointer(self)),
		uintptr(nDevice),
		uintptr(unsafe.Pointer(&device)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMDeviceCollection::Item failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

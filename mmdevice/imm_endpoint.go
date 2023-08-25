package mmdevice

import (
	"fmt"
	"syscall"
	"unsafe"
	"wasapi/com"

	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("1BE09788-6894-4089-8586-9A2A6C265AC5")
var _IID_IMMEndpoint = windows.GUID{Data1: 0x1BE09788, Data2: 0x6894, Data3: 0x4089, Data4: [8]byte{0x85, 0x86, 0x9A, 0x2A, 0x6C, 0x26, 0x5A, 0xC5}}

func IID_IMMEndpoint() windows.GUID {
	return _IID_IMMEndpoint
}

type IMMEndpoint struct {
	vtbl *_IMMEndpointVtbl
}

type _IMMEndpointVtbl struct {
	com.IUnknownVtbl
	GetDataFlow uintptr
}

func (self *IMMEndpoint) Release() (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Release, uintptr(unsafe.Pointer(self)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMEndpoint::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetDataFlow 方法指示音频端点设备是呈现设备还是捕获设备。
func (self *IMMEndpoint) GetDataFlow() (dataFlow EDataFlow, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetDataFlow, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&dataFlow)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IMMEndpoint::GetDataFlow failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

package audioclient

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/xray-bit/wasapi/com"
	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("CD63314F-3FBA-4a1b-812C-EF96358728E7")
var _IID_IAudioClock = windows.GUID{Data1: 0xCD63314F, Data2: 0x3FBA, Data3: 0x4a1b, Data4: [8]byte{0x81, 0x2C, 0xEF, 0x96, 0x35, 0x87, 0x28, 0xE7}}

func IID_IAudioClock() windows.GUID {
	return _IID_IAudioClock
}

type IAudioClock struct {
	vtbl *_IAudioClockVtbl
}

type _IAudioClockVtbl struct {
	com.IUnknownVtbl
	GetFrequency       uintptr
	GetPosition        uintptr
	GetCharacteristics uintptr
}

func (self *IAudioClock) Release() (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Release, uintptr(unsafe.Pointer(self)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClock::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetFrequency 方法获取设备频率。
func (self *IAudioClock) GetFrequency() (frequency uint64, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetFrequency, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&frequency)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClock::GetFrequency failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetPosition 方法获取当前设备位置。
func (self *IAudioClock) GetPosition() (position uint64, QPCPosition uint64, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetPosition, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&position)),
		uintptr(unsafe.Pointer(&QPCPosition)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClock::GetPosition failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetCharacteristics 方法保留供将来使用。
func (self *IAudioClock) GetCharacteristics() (characteristics uint32, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetCharacteristics, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&characteristics)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClock::GetCharacteristics failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

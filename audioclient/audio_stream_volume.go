package audioclient

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/xray-bit/wasapi/com"
	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("93014887-242D-4068-8A15-CF5E93B90FE3")
var _IID_IAudioStreamVolume = windows.GUID{Data1: 0x93014887, Data2: 0x242D, Data3: 0x4068, Data4: [8]byte{0x8A, 0x15, 0xCF, 0x5E, 0x93, 0xB9, 0x0F, 0xE3}}

func IID_IAudioStreamVolume() windows.GUID {
	return _IID_IAudioStreamVolume
}

type IAudioStreamVolume struct {
	vtbl *_IAudioStreamVolumeVtbl
}

type _IAudioStreamVolumeVtbl struct {
	com.IUnknownVtbl
	GetChannelCount  uintptr
	SetChannelVolume uintptr
	GetChannelVolume uintptr
	SetAllVolumes    uintptr
	GetAllVolumes    uintptr
}

func (self *IAudioStreamVolume) Release() (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Release, uintptr(unsafe.Pointer(self)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioStreamVolume::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetChannelCount 方法检索音频流中的通道数。
func (self *IAudioStreamVolume) GetChannelCount() (count uint32, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetChannelCount, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&count)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioStreamVolume::GetChannelCount failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// SetChannelVolume 方法设置音频流中指定通道的音量级别。
func (self *IAudioStreamVolume) SetChannelVolume(index uint32, level float32) (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.SetChannelVolume, uintptr(unsafe.Pointer(self)),
		uintptr(index),
		uintptr(level),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioStreamVolume::SetChannelVolume failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetChannelVolume 方法检索音频流中指定声道的音量级别。
func (self *IAudioStreamVolume) GetChannelVolume(index uint32) (level float32, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetChannelVolume, uintptr(unsafe.Pointer(self)),
		uintptr(index),
		uintptr(unsafe.Pointer(&level)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioStreamVolume::GetChannelVolume failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// SetAllVolumes 方法设置音频流中所有声道的单个音量级别。
func (self *IAudioStreamVolume) SetAllVolumes(count uint32, level float32) (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.SetAllVolumes, uintptr(unsafe.Pointer(self)),
		uintptr(count),
		uintptr(level),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioStreamVolume::SetAllVolumes failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetAllVolumes 方法检索音频流中所有声道的音量级别。
func (self *IAudioStreamVolume) GetAllVolumes(count uint32) (level float32, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetAllVolumes, uintptr(unsafe.Pointer(self)),
		uintptr(count),
		uintptr(unsafe.Pointer(&level)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioStreamVolume::GetAllVolumes failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

package audioclient

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/xray-bit/wasapi/com"
	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("87CE5498-68D6-44E5-9215-6DA47EF883D8")
var _IID_ISimpleAudioVolume = windows.GUID{Data1: 0x87CE5498, Data2: 0x68D6, Data3: 0x44E5, Data4: [8]byte{0x92, 0x15, 0x6D, 0xA4, 0x7E, 0xF8, 0x83, 0xD8}}

func IID_ISimpleAudioVolume() windows.GUID {
	return _IID_ISimpleAudioVolume
}

type ISimpleAudioVolume struct {
	vtbl *_ISimpleAudioVolumeVtbl
}

type _ISimpleAudioVolumeVtbl struct {
	com.IUnknownVtbl
	SetMasterVolume uintptr
	GetMasterVolume uintptr
	SetMute         uintptr
	GetMute         uintptr
}

func (self *ISimpleAudioVolume) Release() (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Release, uintptr(unsafe.Pointer(self)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("ISimpleAudioVolume::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// SetMasterVolume 方法设置音频会话的主音量级别。
func (self *ISimpleAudioVolume) SetMasterVolume(level float32, eventContext *windows.GUID) (count uint32, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.SetMasterVolume, uintptr(unsafe.Pointer(self)),
		uintptr(level),
		uintptr(unsafe.Pointer(eventContext)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("ISimpleAudioVolume::SetMasterVolume failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetMasterVolume 方法检索音频会话的客户端音量级别。
func (self *ISimpleAudioVolume) GetMasterVolume() (level float32, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetMasterVolume, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&level)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("ISimpleAudioVolume::GetMasterVolume failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// SetMute 方法设置音频会话的静音状态。
func (self *ISimpleAudioVolume) SetMute(mute bool, eventContext *windows.GUID) (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.SetMute, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&mute)),
		uintptr(unsafe.Pointer(eventContext)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("ISimpleAudioVolume::SetMute failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetMute 方法检索音频会话的当前静音状态。
func (self *ISimpleAudioVolume) GetMute() (mute bool, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetMute, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&mute)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("ISimpleAudioVolume::GetMute failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

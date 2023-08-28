package audioclient

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/xray-bit/wasapi/com"
	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("1C158861-B533-4B30-B1CF-E853E51C59B8")
var _IID_IChannelAudioVolume = windows.GUID{Data1: 0x1C158861, Data2: 0xB533, Data3: 0x4B30, Data4: [8]byte{0xB1, 0xCF, 0xE8, 0x53, 0xE5, 0x1C, 0x59, 0xB8}}

func IID_IChannelAudioVolume() windows.GUID {
	return _IID_IChannelAudioVolume
}

type IChannelAudioVolume struct {
	vtbl *_IChannelAudioVolumeVtbl
}

type _IChannelAudioVolumeVtbl struct {
	com.IUnknownVtbl
	GetChannelCount  uintptr
	SetChannelVolume uintptr
	GetChannelVolume uintptr
	SetAllVolumes    uintptr
	GetAllVolumes    uintptr
}

func (self *IChannelAudioVolume) Release() (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Release, uintptr(unsafe.Pointer(self)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IChannelAudioVolume::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetChannelCount 方法检索音频会话的流格式的通道数。
func (self *IChannelAudioVolume) GetChannelCount() (count uint32, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetChannelCount, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&count)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IChannelAudioVolume::GetChannelCount failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// SetChannelVolume 方法设置音频会话中指定声道的音量级别。
func (self *IChannelAudioVolume) SetChannelVolume(index uint32, level float32, eventContext *windows.GUID) (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.SetChannelVolume, uintptr(unsafe.Pointer(self)),
		uintptr(index),
		uintptr(level),
		uintptr(unsafe.Pointer(eventContext)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IChannelAudioVolume::SetChannelVolume failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetChannelVolume 方法检索音频会话中指定声道的音量级别。
func (self *IChannelAudioVolume) GetChannelVolume(index uint32) (level float32, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetChannelVolume, uintptr(unsafe.Pointer(self)),
		uintptr(index),
		uintptr(unsafe.Pointer(&level)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IChannelAudioVolume::GetChannelVolume failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// SetAllVolumes 方法设置音频会话中所有声道的单个音量级别。
func (self *IChannelAudioVolume) SetAllVolumes(count uint32, level []float32, eventContext *windows.GUID) (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.SetAllVolumes, uintptr(unsafe.Pointer(self)),
		uintptr(count),
		uintptr(unsafe.Pointer(&level)),
		uintptr(unsafe.Pointer(eventContext)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IChannelAudioVolume::SetAllVolumes failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetAllVolumes 方法检索音频会话中所有通道的音量级别。
func (self *IChannelAudioVolume) GetAllVolumes(count uint32) (level []float32, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetAllVolumes, uintptr(unsafe.Pointer(self)),
		uintptr(count),
		uintptr(unsafe.Pointer(&level)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IChannelAudioVolume::GetAllVolumes failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

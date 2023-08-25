package audioclient

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/xray-bit/wasapi/com"
	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("F294ACFC-3146-4483-A7BF-ADDCA7C260E2")
var _IID_IAudioRenderClient = windows.GUID{Data1: 0xF294ACFC, Data2: 0x3146, Data3: 0x4483, Data4: [8]byte{0xA7, 0xBF, 0xAD, 0xDC, 0xA7, 0xC2, 0x60, 0xE2}}

func IID_IAudioRenderClient() windows.GUID {
	return _IID_IAudioRenderClient
}

type IAudioRenderClient struct {
	vtbl *_IAudioRenderClientVtbl
}

type _IAudioRenderClientVtbl struct {
	com.IUnknownVtbl
	GetBuffer         uintptr
	ReleaseBuffer     uintptr
	GetNextPacketSize uintptr
}

func (self *IAudioRenderClient) Release() (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Release, uintptr(unsafe.Pointer(self)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioRenderClient::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// 检索指向呈现终结点缓冲区中下一个可用空间的指针，调用方可以在其中写入数据包。
func (self *IAudioRenderClient) GetBuffer(numFramesRequested uint32) (data []byte, err error) {
	var buf *uint8
	r, _, _ := syscall.SyscallN(self.vtbl.GetBuffer, uintptr(unsafe.Pointer(self)),
		uintptr(numFramesRequested),
		uintptr(unsafe.Pointer(&buf)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioRenderClient::GetBuffer failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	data = unsafe.Slice(buf, numFramesRequested)

	return
}

// ReleaseBuffer 方法释放在对 IAudioRenderClient::GetBuffer 方法的上一次调用中获取的缓冲区空间。
func (self *IAudioRenderClient) ReleaseBuffer(numFramesWritten uint32, flags uint32) (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.ReleaseBuffer, uintptr(unsafe.Pointer(self)),
		uintptr(numFramesWritten),
		uintptr(flags),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioRenderClient::ReleaseBuffer failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

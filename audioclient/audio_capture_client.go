package audioclient

import (
	"fmt"
	"syscall"
	"unsafe"
	"wasapi/com"

	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("C8ADBD64-E71E-48a0-A4DE-185C395CD317")
var _IID_IAudioCaptureClient = windows.GUID{Data1: 0xC8ADBD64, Data2: 0xE71E, Data3: 0x48a0, Data4: [8]byte{0xA4, 0xDE, 0x18, 0x5C, 0x39, 0x5C, 0xD3, 0x17}}

func IID_IAudioCaptureClient() windows.GUID {
	return _IID_IAudioCaptureClient
}

type IAudioCaptureClient struct {
	vtbl *_IAudioCaptureClientVtbl
}

type _IAudioCaptureClientVtbl struct {
	com.IUnknownVtbl
	GetBuffer         uintptr
	ReleaseBuffer     uintptr
	GetNextPacketSize uintptr
}

func (self *IAudioCaptureClient) Release() (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Release, uintptr(unsafe.Pointer(self)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// 检索指向捕获终结点缓冲区中下一个可用数据包的指针。
//
// 音频帧的大小由客户端通过调用 IAudioClient::GetMixFormat 方法获取的 WAVEFORMATEX (或 WAVEFORMATEXTENSIBLE)
// 结构的 nBlockAlign 成员指定。 音频帧的大小（以字节为单位）等于流中的通道数乘以每个通道的样本大小。 例如，
// 包含 16 位样本的立体声 (2 声道) 流的帧大小为 4 个字节。
func (self *IAudioCaptureClient) GetBuffer() (data []byte, numFramesToRead uint32, flags uint32, devicePosition uint64, QPCPosition uint64, err error) {
	var buf *uint8
	r, _, _ := syscall.SyscallN(self.vtbl.GetBuffer, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&buf)),
		uintptr(unsafe.Pointer(&numFramesToRead)),
		uintptr(unsafe.Pointer(&flags)),
		uintptr(unsafe.Pointer(&devicePosition)),
		uintptr(unsafe.Pointer(&QPCPosition)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioCaptureClient::GetBuffer failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	data = unsafe.Slice(buf, numFramesToRead)

	return
}

// ReleaseBuffer 方法释放缓冲区。
func (self *IAudioCaptureClient) ReleaseBuffer(numFramesToRead uint32) (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.ReleaseBuffer, uintptr(unsafe.Pointer(self)),
		uintptr(numFramesToRead),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioCaptureClient::ReleaseBuffer failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetNextPacketSize 方法检索捕获终结点缓冲区中下一个数据包中的帧数。
func (self *IAudioCaptureClient) GetNextPacketSize() (numFramesInNextPacket uint32, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetNextPacketSize, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&numFramesInNextPacket)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioCaptureClient::GetNextPacketSize failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

package audioclient

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/xray-bit/wasapi/com"
	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("1CB9AD4C-DBFA-4C32-B178-C2F568A703B2")
var _IID_IAudioClient = windows.GUID{Data1: 0x1CB9AD4C, Data2: 0xDBFA, Data3: 0x4C32, Data4: [8]byte{0xB1, 0x78, 0xC2, 0xF5, 0x68, 0xA7, 0x03, 0xB2}}

func IID_IAudioClient() windows.GUID {
	return _IID_IAudioClient
}

type IAudioClient struct {
	vtbl *_IAudioClientVtbl
}

type _IAudioClientVtbl struct {
	com.IUnknownVtbl
	Initialize        uintptr
	GetBufferSize     uintptr
	GetStreamLatency  uintptr
	GetCurrentPadding uintptr
	IsFormatSupported uintptr
	GetMixFormat      uintptr
	GetDevicePeriod   uintptr
	Start             uintptr
	Stop              uintptr
	Reset             uintptr
	SetEventHandle    uintptr
	GetService        uintptr
}

func (client *IAudioClient) Release() (err error) {
	r, _, _ := syscall.SyscallN(client.vtbl.Release, uintptr(unsafe.Pointer(client)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::Release failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// Initialize 方法初始化音频流。
func (client *IAudioClient) Initialize(
	shareMode AUDCLNT_SHAREMODE,
	streamFlags uint32,
	hnsBufferDuration uint64,
	hnsPeriodicity uint64,
	format *WAVEFORMATEXTENSIBLE,
	audioSessionGuid *windows.GUID,
) (err error) {
	formatBuf := format.toBytes()
	r, _, _ := syscall.SyscallN(client.vtbl.Initialize, uintptr(unsafe.Pointer(client)),
		uintptr(shareMode),
		uintptr(streamFlags),
		uintptr(hnsBufferDuration),
		uintptr(hnsPeriodicity),
		uintptr(unsafe.Pointer(&formatBuf[0])),
		uintptr(unsafe.Pointer(audioSessionGuid)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::Initialize failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetBufferSize 方法检索终结点缓冲区的大小 (最大容量) 。
func (client *IAudioClient) GetBufferSize() (numBufferFrames uint32, err error) {
	r, _, _ := syscall.SyscallN(client.vtbl.GetBufferSize, uintptr(unsafe.Pointer(client)),
		uintptr(unsafe.Pointer(&numBufferFrames)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::GetBufferSize failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetStreamLatency 方法检索当前流的最大延迟，并且可以在初始化流后随时调用。
func (client *IAudioClient) GetStreamLatency() (hnsLatency uint64, err error) {
	r, _, _ := syscall.SyscallN(client.vtbl.GetStreamLatency, uintptr(unsafe.Pointer(client)),
		uintptr(unsafe.Pointer(&hnsLatency)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::GetStreamLatency failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetCurrentPadding 方法检索终结点缓冲区中填充的帧数。
func (client *IAudioClient) GetCurrentPadding() (numPaddingFrames uint32, err error) {
	r, _, _ := syscall.SyscallN(client.vtbl.GetCurrentPadding, uintptr(unsafe.Pointer(client)),
		uintptr(unsafe.Pointer(&numPaddingFrames)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::GetCurrentPadding failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// IsFormatSupported 方法指示音频终结点设备是否支持特定的流格式。
func (client *IAudioClient) IsFormatSupported(shareMode AUDCLNT_SHAREMODE, format *WAVEFORMATEXTENSIBLE) (closestMatch WAVEFORMATEXTENSIBLE, err error) {
	var (
		closestMatchPtr *byte
		formatBuf       []byte
	)

	formatBuf = format.toBytes()

	r, _, _ := syscall.SyscallN(client.vtbl.IsFormatSupported, uintptr(unsafe.Pointer(client)),
		uintptr(shareMode),
		uintptr(unsafe.Pointer(&formatBuf[0])),
		uintptr(unsafe.Pointer(&closestMatchPtr)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::IsFormatSupported failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	if shareMode == AUDCLNT_SHAREMODE_SHARED {
		buf := unsafe.Slice(closestMatchPtr, 40)
		closestMatch.fromBytes(buf)

		// 释放内存
		com.CoTaskMemFree(unsafe.Pointer(closestMatchPtr))
	}

	return
}

// GetMixFormat 方法检索音频引擎用于内部处理共享模式流的流格式。
func (client *IAudioClient) GetMixFormat() (deviceFormat WAVEFORMATEXTENSIBLE, err error) {
	var formatPtr *byte
	r, _, _ := syscall.SyscallN(client.vtbl.GetMixFormat, uintptr(unsafe.Pointer(client)),
		uintptr(unsafe.Pointer(&formatPtr)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::GetMixFormat failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	buf := unsafe.Slice(formatPtr, 40)
	deviceFormat.fromBytes(buf)

	// 释放内存
	com.CoTaskMemFree(unsafe.Pointer(formatPtr))

	return
}

// GetDevicePeriod 方法检索音频引擎对终结点缓冲区中数据的连续处理过程分隔的周期间隔的长度。
func (client *IAudioClient) GetDevicePeriod() (hnsDefaultDevicePeriod, hnsMinimumDevicePeriod uint64, err error) {
	r, _, _ := syscall.SyscallN(client.vtbl.GetDevicePeriod, uintptr(unsafe.Pointer(client)),
		uintptr(unsafe.Pointer(&hnsDefaultDevicePeriod)),
		uintptr(unsafe.Pointer(&hnsMinimumDevicePeriod)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::GetDevicePeriod failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// Start 方法启动音频流。
func (client *IAudioClient) Start() (err error) {
	r, _, _ := syscall.SyscallN(client.vtbl.Start, uintptr(unsafe.Pointer(client)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::Start failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// Stop 方法停止音频流。
func (client *IAudioClient) Stop() (err error) {
	r, _, _ := syscall.SyscallN(client.vtbl.Stop, uintptr(unsafe.Pointer(client)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::Stop failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// Reset 方法重置音频流。
func (client *IAudioClient) Reset() (err error) {
	r, _, _ := syscall.SyscallN(client.vtbl.Reset, uintptr(unsafe.Pointer(client)))

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::Reset failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// SetEventHandle 方法设置当音频缓冲区可供客户端处理时系统发出信号的事件句柄。
func (client *IAudioClient) SetEventHandle(eventHandle windows.Handle) (err error) {
	r, _, _ := syscall.SyscallN(client.vtbl.SetEventHandle, uintptr(unsafe.Pointer(client)),
		uintptr(eventHandle),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::SetEventHandle failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

// GetService 方法从音频客户端对象访问其他服务。
func (client *IAudioClient) GetService(iid *windows.GUID) (v unsafe.Pointer, err error) {
	r, _, _ := syscall.SyscallN(client.vtbl.GetService, uintptr(unsafe.Pointer(client)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&v)),
	)

	if com.HRESULT(r) != com.HRESULT(windows.S_OK) {
		err = fmt.Errorf("IAudioClient::GetService failed with code: 0x%08X", com.HRESULT(r))
		return
	}

	return
}

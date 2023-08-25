package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/xray-bit/wasapi/audioclient"
	"github.com/xray-bit/wasapi/com"
	"github.com/xray-bit/wasapi/mmdevice"
	"golang.org/x/sys/windows"
)

const CLSCTX_ALL = windows.CLSCTX_INPROC_SERVER | windows.CLSCTX_INPROC_HANDLER | windows.CLSCTX_LOCAL_SERVER | windows.CLSCTX_REMOTE_SERVER
const (
	REFTIMES_PER_SEC      = 10000000
	REFTIMES_PER_MILLISEC = 10000
)

var wg sync.WaitGroup

func main() {
	var (
		v                 unsafe.Pointer
		enumerator        *mmdevice.IMMDeviceEnumerator
		devices           *mmdevice.IMMDeviceCollection
		device            *mmdevice.IMMDevice
		cnt               uint
		id                string
		properties        *com.IPropertyStore
		propVar           com.PROPVARIANT
		audioClient       *audioclient.IAudioClient
		captureClient     *audioclient.IAudioCaptureClient
		format            *audioclient.WAVEFORMATEXTENSIBLE
		hnsActualDuration uint64
		bufferFrameCount  uint32
		quit              chan int
		sigChan           chan os.Signal
		err               error
	)

	CLSID_MMDeviceEnumerator := mmdevice.CLSID_MMDeviceEnumerator()
	IID_IMMDeviceEnumerator := mmdevice.IID_IMMDeviceEnumerator()
	IID_IAudioCaptureClient := audioclient.IID_IAudioCaptureClient()

	// 初始化 COM 库
	if err = com.CoInitializeEx(0, windows.COINIT_APARTMENTTHREADED); err != nil {
		panic(err)
	}
	defer com.CoUninitialize()

	// 创建 IMMDeviceEnumerator 实例
	if v, err = com.CoCreateInstance(&CLSID_MMDeviceEnumerator, nil, CLSCTX_ALL, &IID_IMMDeviceEnumerator); err != nil {
		panic(err)
	}
	enumerator = mmdevice.ToType[mmdevice.IMMDeviceEnumerator](v)
	defer enumerator.Release()

	// 枚举音频端点设备
	if devices, err = enumerator.EnumAudioEndpoints(mmdevice.ERender, mmdevice.DEVICE_STATE_ACTIVE); err != nil {
		panic(err)
	}
	defer devices.Release()

	// 获取设备数量
	if cnt, err = devices.GetCount(); err != nil {
		panic(err)
	}
	fmt.Println("Device Counts:", cnt)

	if cnt == 0 {
		return
	}

	// 获取索引为 0 的设备
	if device, err = devices.Item(0); err != nil {
		panic(err)
	}
	defer device.Release()

	if id, err = device.GetId(); err != nil {
		panic(err)
	}

	fmt.Println("Device[0] ID:", id)

	if properties, err = device.OpenPropertyStore(com.STGM_READ); err != nil {
		panic(err)
	}
	defer properties.Release()

	// 获取设备友好名称
	if propVar, err = properties.GetValue(com.PKEY_Device_FriendlyName()); err != nil {
		panic(err)
	}

	fmt.Println("Device[0] Name:", windows.UTF16PtrToString(propVar.PwszVal()))

	// 创建音频客户端实例
	if v, err = device.Activate(audioclient.IID_IAudioClient(), CLSCTX_ALL, nil); err != nil {
		panic(err)
	}
	audioClient = audioclient.ToType[audioclient.IAudioClient](v)
	defer audioClient.Release()

	// 获取设备格式
	if format, err = audioClient.GetMixFormat(); err != nil {
		panic(err)
	}
	defer com.CoTaskMemFree(unsafe.Pointer(format))

	showWaveFormat(format)

	// 初始化音频流
	if err = audioClient.Initialize(
		audioclient.AUDCLNT_SHAREMODE_SHARED,
		audioclient.AUDCLNT_STREAMFLAGS_LOOPBACK,
		REFTIMES_PER_SEC, 0, format, nil,
	); err != nil {
		panic(err)
	}

	// 获取音频客户端缓冲区大小
	if bufferFrameCount, err = audioClient.GetBufferSize(); err != nil {
		panic(err)
	}

	// 访问捕捉流服务
	if v, err = audioClient.GetService(&IID_IAudioCaptureClient); err != nil {
		panic(err)
	}
	captureClient = audioclient.ToType[audioclient.IAudioCaptureClient](v)
	defer captureClient.Release()

	// 计算实际的采样时间
	hnsActualDuration = uint64(REFTIMES_PER_SEC) * uint64(bufferFrameCount) / uint64(format.Format().SamplesPerSec())

	if err = audioClient.Start(); err != nil {
		panic(err)
	}
	defer audioClient.Stop()

	quit = make(chan int)
	sigChan = make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)

	// start process
	go processCapture(captureClient, format, hnsActualDuration, quit)

	<-sigChan
	close(quit)

	// wait for quit
	wg.Wait()
}

func processCapture(client *audioclient.IAudioCaptureClient, format *audioclient.WAVEFORMATEXTENSIBLE, hnsDuration uint64, quit <-chan int) {
	var (
		packetLen          uint32
		data               []byte
		numFramesAvailable uint32
		flags              uint32
		err                error
	)

	defer wg.Done()
	defer func() {
		fmt.Println("Exit Capture")
	}()

	for {
		select {
		case <-quit:
			return
		case <-time.After(time.Duration(hnsDuration/REFTIMES_PER_MILLISEC/2) * time.Millisecond):
		}

		if packetLen, err = client.GetNextPacketSize(); err != nil {
			panic(err)
		}

		for packetLen != 0 {
			if data, numFramesAvailable, flags, _, _, err = client.GetBuffer(); err != nil {
				panic(err)
			}

			// 这里打印只是方便呈现获取到的音频内容
			// 可以在此按业务需求处理音频内容
			// 因为此处判断了静音，所以如果卡在了这里，可以尝试播放一些音乐
			if flags&audioclient.AUDCLNT_BUFFERFLAGS_SILENT == 0 &&
				format.Format().BitsPerSample() == 32 &&
				format.Format().FormatTag() == audioclient.WAVE_FORMAT_EXTENSIBLE &&
				format.SubFormatTag() == audioclient.WAVE_FORMAT_IEEE_FLOAT {
				for i := 0; i < int(numFramesAvailable/uint32(format.Format().BlockAlign()))/int(format.Format().Channels()); i++ {
					for j := 0; j < int(format.Format().Channels()); j++ {
						fmt.Printf("CH%d: %f\t", j, math.Float32frombits(binary.LittleEndian.Uint32(data[i*int(format.Format().Channels())*4:i*int(format.Format().Channels())*4+4])))
					}
					fmt.Println()
				}
			}

			if err = client.ReleaseBuffer(numFramesAvailable); err != nil {
				panic(err)
			}

			if packetLen, err = client.GetNextPacketSize(); err != nil {
				panic(err)
			}
		}
	}
}

func showWaveFormat(wf *audioclient.WAVEFORMATEXTENSIBLE) {
	fmt.Println("FormatTag:\t", wf.Format().FormatTag())
	fmt.Println("Channels:\t", wf.Format().Channels())
	fmt.Println("SamplesPerSec:\t", wf.Format().SamplesPerSec())
	fmt.Println("AvgBytesPerSec:\t", wf.Format().AvgBytesPerSec())
	fmt.Println("BlockAlign:\t", wf.Format().BlockAlign())
	fmt.Println("BitsPerSample:\t", wf.Format().BitsPerSample())
	fmt.Println("CbSize:\t\t", wf.Format().CbSize())
	if wf.Format().CbSize() == 22 {
		fmt.Println("Samples:\t", wf.Samples())
		fmt.Printf("ChannelMask:\t 0x%08X\n", wf.ChannelMask())
		fmt.Println("SubFormat:\t", wf.SubFormatTag())
	}
}

package mmdevice

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// DECLSPEC_UUID("BCDE0395-E52F-467C-8E3D-C4579291692E")
var _CLSID_MMDeviceEnumerator = windows.GUID{Data1: 0xBCDE0395, Data2: 0xE52F, Data3: 0x467C, Data4: [8]byte{0x8E, 0x3D, 0xC4, 0x57, 0x92, 0x91, 0x69, 0x2E}}

func CLSID_MMDeviceEnumerator() windows.GUID {
	return _CLSID_MMDeviceEnumerator
}

const (
	DEVICE_STATE_ACTIVE     = 0x00000001
	DEVICE_STATE_DISABLED   = 0x00000002
	DEVICE_STATE_NOTPRESENT = 0x00000004
	DEVICE_STATE_UNPLUGGED  = 0x00000008
	DEVICE_STATEMASK_ALL    = 0x0000000F
)

type EDataFlow uint32

const (
	ERender EDataFlow = iota
	ECapture
	EAll
	EDataFlow_enum_count
)

type ERole uint32

const (
	EConsole ERole = iota
	EMultimedia
	ECommunications
	ERole_enum_count
)

type EndpointFormFactor uint32

const (
	RemoteNetworkDevice EndpointFormFactor = iota
	Speakers
	LineLevel
	Headphones
	Microphone
	Headset
	Handset
	UnknownDigitalPassthrough
	SPDIF
	DigitalAudioDisplayDevice
	UnknownFormFactor
	EndpointFormFactor_enum_count
)

func ToType[T IMMDevice | IMMDeviceCollection | IMMDeviceEnumerator | IMMEndpoint | IMMNotificationClient](v unsafe.Pointer) *T {
	return (*T)(v)
}

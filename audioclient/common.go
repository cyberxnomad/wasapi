package audioclient

import (
	"encoding/binary"
	"unsafe"

	"golang.org/x/sys/windows"
)

type AUDCLNT_SHAREMODE uint32

const (
	AUDCLNT_SHAREMODE_SHARED AUDCLNT_SHAREMODE = iota
	AUDCLNT_SHAREMODE_EXCLUSIVE
)

type WAVEFORMATEX struct {
	// FormatTag      uint16
	// Channels       uint16
	// SamplesPerSec  uint32
	// AvgBytesPerSec uint32
	// BlockAlign     uint16
	// BitsPerSample  uint16
	// CbSize         uint16
	mem [18]byte
}

func (self *WAVEFORMATEX) FormatTag() uint16 {
	return binary.LittleEndian.Uint16(self.mem[0:2])
}

func (self *WAVEFORMATEX) SetFormatTag(v uint16) {
	binary.LittleEndian.PutUint16(self.mem[0:2], v)
}

func (self *WAVEFORMATEX) Channels() uint16 {
	return binary.LittleEndian.Uint16(self.mem[2:4])
}

func (self *WAVEFORMATEX) SetChannels(v uint16) {
	binary.LittleEndian.PutUint16(self.mem[2:4], v)
}

func (self *WAVEFORMATEX) SamplesPerSec() uint32 {
	return binary.LittleEndian.Uint32(self.mem[4:8])
}

func (self *WAVEFORMATEX) SetSamplesPerSec(v uint32) {
	binary.LittleEndian.PutUint32(self.mem[4:8], v)
}

func (self *WAVEFORMATEX) AvgBytesPerSec() uint32 {
	return binary.LittleEndian.Uint32(self.mem[8:12])
}

func (self *WAVEFORMATEX) SetAvgBytesPerSec(v uint32) {
	binary.LittleEndian.PutUint32(self.mem[8:12], v)
}

func (self *WAVEFORMATEX) BlockAlign() uint16 {
	return binary.LittleEndian.Uint16(self.mem[12:14])
}

func (self *WAVEFORMATEX) SetBlockAlign(v uint16) {
	binary.LittleEndian.PutUint16(self.mem[12:14], v)
}

func (self *WAVEFORMATEX) BitsPerSample() uint16 {
	return binary.LittleEndian.Uint16(self.mem[14:16])
}

func (self *WAVEFORMATEX) SetBitsPerSample(v uint16) {
	binary.LittleEndian.PutUint16(self.mem[14:16], v)
}

func (self *WAVEFORMATEX) CbSize() uint16 {
	return binary.LittleEndian.Uint16(self.mem[16:18])
}

func (self *WAVEFORMATEX) SetCbSize(v uint16) {
	binary.LittleEndian.PutUint16(self.mem[16:18], v)
}

type WAVEFORMATEXTENSIBLE struct {
	// Format      WAVEFORMATEX
	// Samples     uint16
	// ChannelMask uint32
	// SubFormat   windows.GUID
	mem [40]byte
}

func (self *WAVEFORMATEXTENSIBLE) Format() *WAVEFORMATEX {
	return (*WAVEFORMATEX)(unsafe.Pointer(&self.mem[0]))
}

func (self *WAVEFORMATEXTENSIBLE) SetFormat(v WAVEFORMATEX) {
	copy(self.mem[:18], v.mem[:])
}

func (self *WAVEFORMATEXTENSIBLE) Samples() uint16 {
	return binary.LittleEndian.Uint16(self.mem[18:20])
}

func (self *WAVEFORMATEXTENSIBLE) SetSamples(v uint16) {
	binary.LittleEndian.PutUint16(self.mem[18:20], v)
}

func (self *WAVEFORMATEXTENSIBLE) ChannelMask() uint32 {
	return binary.LittleEndian.Uint32(self.mem[20:24])
}

func (self *WAVEFORMATEXTENSIBLE) SetChannelMask(v uint32) {
	binary.LittleEndian.PutUint32(self.mem[20:24], v)
}

func (self *WAVEFORMATEXTENSIBLE) SubFormat() windows.GUID {
	return *(*windows.GUID)(unsafe.Pointer(&self.mem[24]))
}

func (self *WAVEFORMATEXTENSIBLE) SetSubFormat(v windows.GUID) {
	guid := *(*[unsafe.Sizeof(v)]byte)(unsafe.Pointer(&v))
	copy(self.mem[24:], guid[:])
}

func (self *WAVEFORMATEXTENSIBLE) SubFormatTag() uint16 {
	return uint16(self.SubFormat().Data1)
}

// WAVE form wFormatTag IDs
const (
	WAVE_FORMAT_UNKNOWN                    uint16 = 0x0000 // Microsoft Corporation
	WAVE_FORMAT_PCM                        uint16 = 0x0001 //
	WAVE_FORMAT_ADPCM                      uint16 = 0x0002 // Microsoft Corporation
	WAVE_FORMAT_IEEE_FLOAT                 uint16 = 0x0003 // Microsoft Corporation
	WAVE_FORMAT_VSELP                      uint16 = 0x0004 // Compaq Computer Corp.
	WAVE_FORMAT_IBM_CVSD                   uint16 = 0x0005 // IBM Corporation
	WAVE_FORMAT_ALAW                       uint16 = 0x0006 // Microsoft Corporation
	WAVE_FORMAT_MULAW                      uint16 = 0x0007 // Microsoft Corporation
	WAVE_FORMAT_DTS                        uint16 = 0x0008 // Microsoft Corporation
	WAVE_FORMAT_DRM                        uint16 = 0x0009 // Microsoft Corporation
	WAVE_FORMAT_WMAVOICE9                  uint16 = 0x000A // Microsoft Corporation
	WAVE_FORMAT_WMAVOICE10                 uint16 = 0x000B // Microsoft Corporation
	WAVE_FORMAT_OKI_ADPCM                  uint16 = 0x0010 // OKI
	WAVE_FORMAT_DVI_ADPCM                  uint16 = 0x0011 // Intel Corporation
	WAVE_FORMAT_IMA_ADPCM                  uint16 = 0x0011 //  Intel Corporation
	WAVE_FORMAT_MEDIASPACE_ADPCM           uint16 = 0x0012 // Videologic
	WAVE_FORMAT_SIERRA_ADPCM               uint16 = 0x0013 // Sierra Semiconductor Corp
	WAVE_FORMAT_G723_ADPCM                 uint16 = 0x0014 // Antex Electronics Corporation
	WAVE_FORMAT_DIGISTD                    uint16 = 0x0015 // DSP Solutions, Inc.
	WAVE_FORMAT_DIGIFIX                    uint16 = 0x0016 // DSP Solutions, Inc.
	WAVE_FORMAT_DIALOGIC_OKI_ADPCM         uint16 = 0x0017 // Dialogic Corporation
	WAVE_FORMAT_MEDIAVISION_ADPCM          uint16 = 0x0018 // Media Vision, Inc.
	WAVE_FORMAT_CU_CODEC                   uint16 = 0x0019 // Hewlett-Packard Company
	WAVE_FORMAT_HP_DYN_VOICE               uint16 = 0x001A // Hewlett-Packard Company
	WAVE_FORMAT_YAMAHA_ADPCM               uint16 = 0x0020 // Yamaha Corporation of America
	WAVE_FORMAT_SONARC                     uint16 = 0x0021 // Speech Compression
	WAVE_FORMAT_DSPGROUP_TRUESPEECH        uint16 = 0x0022 // DSP Group, Inc
	WAVE_FORMAT_ECHOSC1                    uint16 = 0x0023 // Echo Speech Corporation
	WAVE_FORMAT_AUDIOFILE_AF36             uint16 = 0x0024 // Virtual Music, Inc.
	WAVE_FORMAT_APTX                       uint16 = 0x0025 // Audio Processing Technology
	WAVE_FORMAT_AUDIOFILE_AF10             uint16 = 0x0026 // Virtual Music, Inc.
	WAVE_FORMAT_PROSODY_1612               uint16 = 0x0027 // Aculab plc
	WAVE_FORMAT_LRC                        uint16 = 0x0028 // Merging Technologies S.A.
	WAVE_FORMAT_DOLBY_AC2                  uint16 = 0x0030 // Dolby Laboratories
	WAVE_FORMAT_GSM610                     uint16 = 0x0031 // Microsoft Corporation
	WAVE_FORMAT_MSNAUDIO                   uint16 = 0x0032 // Microsoft Corporation
	WAVE_FORMAT_ANTEX_ADPCME               uint16 = 0x0033 // Antex Electronics Corporation
	WAVE_FORMAT_CONTROL_RES_VQLPC          uint16 = 0x0034 // Control Resources Limited
	WAVE_FORMAT_DIGIREAL                   uint16 = 0x0035 // DSP Solutions, Inc.
	WAVE_FORMAT_DIGIADPCM                  uint16 = 0x0036 // DSP Solutions, Inc.
	WAVE_FORMAT_CONTROL_RES_CR10           uint16 = 0x0037 // Control Resources Limited
	WAVE_FORMAT_NMS_VBXADPCM               uint16 = 0x0038 // Natural MicroSystems
	WAVE_FORMAT_CS_IMAADPCM                uint16 = 0x0039 // Crystal Semiconductor IMA ADPCM
	WAVE_FORMAT_ECHOSC3                    uint16 = 0x003A // Echo Speech Corporation
	WAVE_FORMAT_ROCKWELL_ADPCM             uint16 = 0x003B // Rockwell International
	WAVE_FORMAT_ROCKWELL_DIGITALK          uint16 = 0x003C // Rockwell International
	WAVE_FORMAT_XEBEC                      uint16 = 0x003D // Xebec Multimedia Solutions Limited
	WAVE_FORMAT_G721_ADPCM                 uint16 = 0x0040 // Antex Electronics Corporation
	WAVE_FORMAT_G728_CELP                  uint16 = 0x0041 // Antex Electronics Corporation
	WAVE_FORMAT_MSG723                     uint16 = 0x0042 // Microsoft Corporation
	WAVE_FORMAT_INTEL_G723_1               uint16 = 0x0043 // Intel Corp.
	WAVE_FORMAT_INTEL_G729                 uint16 = 0x0044 // Intel Corp.
	WAVE_FORMAT_SHARP_G726                 uint16 = 0x0045 // Sharp
	WAVE_FORMAT_MPEG                       uint16 = 0x0050 // Microsoft Corporation
	WAVE_FORMAT_RT24                       uint16 = 0x0052 // InSoft, Inc.
	WAVE_FORMAT_PAC                        uint16 = 0x0053 // InSoft, Inc.
	WAVE_FORMAT_MPEGLAYER3                 uint16 = 0x0055 // ISO/MPEG Layer3 Format Tag
	WAVE_FORMAT_LUCENT_G723                uint16 = 0x0059 // Lucent Technologies
	WAVE_FORMAT_CIRRUS                     uint16 = 0x0060 // Cirrus Logic
	WAVE_FORMAT_ESPCM                      uint16 = 0x0061 // ESS Technology
	WAVE_FORMAT_VOXWARE                    uint16 = 0x0062 // Voxware Inc
	WAVE_FORMAT_CANOPUS_ATRAC              uint16 = 0x0063 // Canopus, co., Ltd.
	WAVE_FORMAT_G726_ADPCM                 uint16 = 0x0064 // APICOM
	WAVE_FORMAT_G722_ADPCM                 uint16 = 0x0065 // APICOM
	WAVE_FORMAT_DSAT                       uint16 = 0x0066 // Microsoft Corporation
	WAVE_FORMAT_DSAT_DISPLAY               uint16 = 0x0067 // Microsoft Corporation
	WAVE_FORMAT_VOXWARE_BYTE_ALIGNED       uint16 = 0x0069 // Voxware Inc
	WAVE_FORMAT_VOXWARE_AC8                uint16 = 0x0070 // Voxware Inc
	WAVE_FORMAT_VOXWARE_AC10               uint16 = 0x0071 // Voxware Inc
	WAVE_FORMAT_VOXWARE_AC16               uint16 = 0x0072 // Voxware Inc
	WAVE_FORMAT_VOXWARE_AC20               uint16 = 0x0073 // Voxware Inc
	WAVE_FORMAT_VOXWARE_RT24               uint16 = 0x0074 // Voxware Inc
	WAVE_FORMAT_VOXWARE_RT29               uint16 = 0x0075 // Voxware Inc
	WAVE_FORMAT_VOXWARE_RT29HW             uint16 = 0x0076 // Voxware Inc
	WAVE_FORMAT_VOXWARE_VR12               uint16 = 0x0077 // Voxware Inc
	WAVE_FORMAT_VOXWARE_VR18               uint16 = 0x0078 // Voxware Inc
	WAVE_FORMAT_VOXWARE_TQ40               uint16 = 0x0079 // Voxware Inc
	WAVE_FORMAT_VOXWARE_SC3                uint16 = 0x007A // Voxware Inc
	WAVE_FORMAT_VOXWARE_SC3_1              uint16 = 0x007B // Voxware Inc
	WAVE_FORMAT_SOFTSOUND                  uint16 = 0x0080 // Softsound, Ltd.
	WAVE_FORMAT_VOXWARE_TQ60               uint16 = 0x0081 // Voxware Inc
	WAVE_FORMAT_MSRT24                     uint16 = 0x0082 // Microsoft Corporation
	WAVE_FORMAT_G729A                      uint16 = 0x0083 // AT&T Labs, Inc.
	WAVE_FORMAT_MVI_MVI2                   uint16 = 0x0084 // Motion Pixels
	WAVE_FORMAT_DF_G726                    uint16 = 0x0085 // DataFusion Systems (Pty) (Ltd)
	WAVE_FORMAT_DF_GSM610                  uint16 = 0x0086 // DataFusion Systems (Pty) (Ltd)
	WAVE_FORMAT_ISIAUDIO                   uint16 = 0x0088 // Iterated Systems, Inc.
	WAVE_FORMAT_ONLIVE                     uint16 = 0x0089 // OnLive! Technologies, Inc.
	WAVE_FORMAT_MULTITUDE_FT_SX20          uint16 = 0x008A // Multitude Inc.
	WAVE_FORMAT_INFOCOM_ITS_G721_ADPCM     uint16 = 0x008B // Infocom
	WAVE_FORMAT_CONVEDIA_G729              uint16 = 0x008C // Convedia Corp.
	WAVE_FORMAT_CONGRUENCY                 uint16 = 0x008D // Congruency Inc.
	WAVE_FORMAT_SBC24                      uint16 = 0x0091 // Siemens Business Communications Sys
	WAVE_FORMAT_DOLBY_AC3_SPDIF            uint16 = 0x0092 // Sonic Foundry
	WAVE_FORMAT_MEDIASONIC_G723            uint16 = 0x0093 // MediaSonic
	WAVE_FORMAT_PROSODY_8KBPS              uint16 = 0x0094 // Aculab plc
	WAVE_FORMAT_ZYXEL_ADPCM                uint16 = 0x0097 // ZyXEL Communications, Inc.
	WAVE_FORMAT_PHILIPS_LPCBB              uint16 = 0x0098 // Philips Speech Processing
	WAVE_FORMAT_PACKED                     uint16 = 0x0099 // Studer Professional Audio AG
	WAVE_FORMAT_MALDEN_PHONYTALK           uint16 = 0x00A0 // Malden Electronics Ltd.
	WAVE_FORMAT_RACAL_RECORDER_GSM         uint16 = 0x00A1 // Racal recorders
	WAVE_FORMAT_RACAL_RECORDER_G720_A      uint16 = 0x00A2 // Racal recorders
	WAVE_FORMAT_RACAL_RECORDER_G723_1      uint16 = 0x00A3 // Racal recorders
	WAVE_FORMAT_RACAL_RECORDER_TETRA_ACELP uint16 = 0x00A4 // Racal recorders
	WAVE_FORMAT_NEC_AAC                    uint16 = 0x00B0 // NEC Corp.
	WAVE_FORMAT_RAW_AAC1                   uint16 = 0x00FF // For Raw AAC, with format block AudioSpecificConfig() (as defined by MPEG-4), that follows WAVEFORMATEX
	WAVE_FORMAT_RHETOREX_ADPCM             uint16 = 0x0100 // Rhetorex Inc.
	WAVE_FORMAT_IRAT                       uint16 = 0x0101 // BeCubed Software Inc.
	WAVE_FORMAT_VIVO_G723                  uint16 = 0x0111 // Vivo Software
	WAVE_FORMAT_VIVO_SIREN                 uint16 = 0x0112 // Vivo Software
	WAVE_FORMAT_PHILIPS_CELP               uint16 = 0x0120 // Philips Speech Processing
	WAVE_FORMAT_PHILIPS_GRUNDIG            uint16 = 0x0121 // Philips Speech Processing
	WAVE_FORMAT_DIGITAL_G723               uint16 = 0x0123 // Digital Equipment Corporation
	WAVE_FORMAT_SANYO_LD_ADPCM             uint16 = 0x0125 // Sanyo Electric Co., Ltd.
	WAVE_FORMAT_SIPROLAB_ACEPLNET          uint16 = 0x0130 // Sipro Lab Telecom Inc.
	WAVE_FORMAT_SIPROLAB_ACELP4800         uint16 = 0x0131 // Sipro Lab Telecom Inc.
	WAVE_FORMAT_SIPROLAB_ACELP8V3          uint16 = 0x0132 // Sipro Lab Telecom Inc.
	WAVE_FORMAT_SIPROLAB_G729              uint16 = 0x0133 // Sipro Lab Telecom Inc.
	WAVE_FORMAT_SIPROLAB_G729A             uint16 = 0x0134 // Sipro Lab Telecom Inc.
	WAVE_FORMAT_SIPROLAB_KELVIN            uint16 = 0x0135 // Sipro Lab Telecom Inc.
	WAVE_FORMAT_VOICEAGE_AMR               uint16 = 0x0136 // VoiceAge Corp.
	WAVE_FORMAT_G726ADPCM                  uint16 = 0x0140 // Dictaphone Corporation
	WAVE_FORMAT_DICTAPHONE_CELP68          uint16 = 0x0141 // Dictaphone Corporation
	WAVE_FORMAT_DICTAPHONE_CELP54          uint16 = 0x0142 // Dictaphone Corporation
	WAVE_FORMAT_QUALCOMM_PUREVOICE         uint16 = 0x0150 // Qualcomm, Inc.
	WAVE_FORMAT_QUALCOMM_HALFRATE          uint16 = 0x0151 // Qualcomm, Inc.
	WAVE_FORMAT_TUBGSM                     uint16 = 0x0155 // Ring Zero Systems, Inc.
	WAVE_FORMAT_MSAUDIO1                   uint16 = 0x0160 // Microsoft Corporation
	WAVE_FORMAT_WMAUDIO2                   uint16 = 0x0161 // Microsoft Corporation
	WAVE_FORMAT_WMAUDIO3                   uint16 = 0x0162 // Microsoft Corporation
	WAVE_FORMAT_WMAUDIO_LOSSLESS           uint16 = 0x0163 // Microsoft Corporation
	WAVE_FORMAT_WMASPDIF                   uint16 = 0x0164 // Microsoft Corporation
	WAVE_FORMAT_UNISYS_NAP_ADPCM           uint16 = 0x0170 // Unisys Corp.
	WAVE_FORMAT_UNISYS_NAP_ULAW            uint16 = 0x0171 // Unisys Corp.
	WAVE_FORMAT_UNISYS_NAP_ALAW            uint16 = 0x0172 // Unisys Corp.
	WAVE_FORMAT_UNISYS_NAP_16K             uint16 = 0x0173 // Unisys Corp.
	WAVE_FORMAT_SYCOM_ACM_SYC008           uint16 = 0x0174 // SyCom Technologies
	WAVE_FORMAT_SYCOM_ACM_SYC701_G726L     uint16 = 0x0175 // SyCom Technologies
	WAVE_FORMAT_SYCOM_ACM_SYC701_CELP54    uint16 = 0x0176 // SyCom Technologies
	WAVE_FORMAT_SYCOM_ACM_SYC701_CELP68    uint16 = 0x0177 // SyCom Technologies
	WAVE_FORMAT_KNOWLEDGE_ADVENTURE_ADPCM  uint16 = 0x0178 // Knowledge Adventure, Inc.
	WAVE_FORMAT_FRAUNHOFER_IIS_MPEG2_AAC   uint16 = 0x0180 // Fraunhofer IIS
	WAVE_FORMAT_DTS_DS                     uint16 = 0x0190 // Digital Theatre Systems, Inc.
	WAVE_FORMAT_CREATIVE_ADPCM             uint16 = 0x0200 // Creative Labs, Inc
	WAVE_FORMAT_CREATIVE_FASTSPEECH8       uint16 = 0x0202 // Creative Labs, Inc
	WAVE_FORMAT_CREATIVE_FASTSPEECH10      uint16 = 0x0203 // Creative Labs, Inc
	WAVE_FORMAT_UHER_ADPCM                 uint16 = 0x0210 // UHER informatic GmbH
	WAVE_FORMAT_ULEAD_DV_AUDIO             uint16 = 0x0215 // Ulead Systems, Inc.
	WAVE_FORMAT_ULEAD_DV_AUDIO_1           uint16 = 0x0216 // Ulead Systems, Inc.
	WAVE_FORMAT_QUARTERDECK                uint16 = 0x0220 // Quarterdeck Corporation
	WAVE_FORMAT_ILINK_VC                   uint16 = 0x0230 // I-link Worldwide
	WAVE_FORMAT_RAW_SPORT                  uint16 = 0x0240 // Aureal Semiconductor
	WAVE_FORMAT_ESST_AC3                   uint16 = 0x0241 // ESS Technology, Inc.
	WAVE_FORMAT_GENERIC_PASSTHRU           uint16 = 0x0249 //
	WAVE_FORMAT_IPI_HSX                    uint16 = 0x0250 // Interactive Products, Inc.
	WAVE_FORMAT_IPI_RPELP                  uint16 = 0x0251 // Interactive Products, Inc.
	WAVE_FORMAT_CS2                        uint16 = 0x0260 // Consistent Software
	WAVE_FORMAT_SONY_SCX                   uint16 = 0x0270 // Sony Corp.
	WAVE_FORMAT_SONY_SCY                   uint16 = 0x0271 // Sony Corp.
	WAVE_FORMAT_SONY_ATRAC3                uint16 = 0x0272 // Sony Corp.
	WAVE_FORMAT_SONY_SPC                   uint16 = 0x0273 // Sony Corp.
	WAVE_FORMAT_TELUM_AUDIO                uint16 = 0x0280 // Telum Inc.
	WAVE_FORMAT_TELUM_IA_AUDIO             uint16 = 0x0281 // Telum Inc.
	WAVE_FORMAT_NORCOM_VOICE_SYSTEMS_ADPCM uint16 = 0x0285 // Norcom Electronics Corp.
	WAVE_FORMAT_FM_TOWNS_SND               uint16 = 0x0300 // Fujitsu Corp.
	WAVE_FORMAT_MICRONAS                   uint16 = 0x0350 // Micronas Semiconductors, Inc.
	WAVE_FORMAT_MICRONAS_CELP833           uint16 = 0x0351 // Micronas Semiconductors, Inc.
	WAVE_FORMAT_BTV_DIGITAL                uint16 = 0x0400 // Brooktree Corporation
	WAVE_FORMAT_INTEL_MUSIC_CODER          uint16 = 0x0401 // Intel Corp.
	WAVE_FORMAT_INDEO_AUDIO                uint16 = 0x0402 // Ligos
	WAVE_FORMAT_QDESIGN_MUSIC              uint16 = 0x0450 // QDesign Corporation
	WAVE_FORMAT_ON2_VP7_AUDIO              uint16 = 0x0500 // On2 Technologies
	WAVE_FORMAT_ON2_VP6_AUDIO              uint16 = 0x0501 // On2 Technologies
	WAVE_FORMAT_VME_VMPCM                  uint16 = 0x0680 // AT&T Labs, Inc.
	WAVE_FORMAT_TPC                        uint16 = 0x0681 // AT&T Labs, Inc.
	WAVE_FORMAT_LIGHTWAVE_LOSSLESS         uint16 = 0x08AE // Clearjump
	WAVE_FORMAT_OLIGSM                     uint16 = 0x1000 // Ing C. Olivetti & C., S.p.A.
	WAVE_FORMAT_OLIADPCM                   uint16 = 0x1001 // Ing C. Olivetti & C., S.p.A.
	WAVE_FORMAT_OLICELP                    uint16 = 0x1002 // Ing C. Olivetti & C., S.p.A.
	WAVE_FORMAT_OLISBC                     uint16 = 0x1003 // Ing C. Olivetti & C., S.p.A.
	WAVE_FORMAT_OLIOPR                     uint16 = 0x1004 // Ing C. Olivetti & C., S.p.A.
	WAVE_FORMAT_LH_CODEC                   uint16 = 0x1100 // Lernout & Hauspie
	WAVE_FORMAT_LH_CODEC_CELP              uint16 = 0x1101 // Lernout & Hauspie
	WAVE_FORMAT_LH_CODEC_SBC8              uint16 = 0x1102 // Lernout & Hauspie
	WAVE_FORMAT_LH_CODEC_SBC12             uint16 = 0x1103 // Lernout & Hauspie
	WAVE_FORMAT_LH_CODEC_SBC16             uint16 = 0x1104 // Lernout & Hauspie
	WAVE_FORMAT_NORRIS                     uint16 = 0x1400 // Norris Communications, Inc.
	WAVE_FORMAT_ISIAUDIO_2                 uint16 = 0x1401 // ISIAudio
	WAVE_FORMAT_SOUNDSPACE_MUSICOMPRESS    uint16 = 0x1500 // AT&T Labs, Inc.
	WAVE_FORMAT_MPEG_ADTS_AAC              uint16 = 0x1600 // Microsoft Corporation
	WAVE_FORMAT_MPEG_RAW_AAC               uint16 = 0x1601 // Microsoft Corporation
	WAVE_FORMAT_MPEG_LOAS                  uint16 = 0x1602 // Microsoft Corporation (MPEG-4 Audio Transport Streams (LOAS/LATM)
	WAVE_FORMAT_NOKIA_MPEG_ADTS_AAC        uint16 = 0x1608 // Microsoft Corporation
	WAVE_FORMAT_NOKIA_MPEG_RAW_AAC         uint16 = 0x1609 // Microsoft Corporation
	WAVE_FORMAT_VODAFONE_MPEG_ADTS_AAC     uint16 = 0x160A // Microsoft Corporation
	WAVE_FORMAT_VODAFONE_MPEG_RAW_AAC      uint16 = 0x160B // Microsoft Corporation
	WAVE_FORMAT_MPEG_HEAAC                 uint16 = 0x1610 // Microsoft Corporation (MPEG-2 AAC or MPEG-4 HE-AAC v1/v2 streams with any payload (ADTS, ADIF, LOAS/LATM, RAW). Format block includes MP4 AudioSpecificConfig() -- see HEAACWAVEFORMAT below
	WAVE_FORMAT_VOXWARE_RT24_SPEECH        uint16 = 0x181C // Voxware Inc.
	WAVE_FORMAT_SONICFOUNDRY_LOSSLESS      uint16 = 0x1971 // Sonic Foundry
	WAVE_FORMAT_INNINGS_TELECOM_ADPCM      uint16 = 0x1979 // Innings Telecom Inc.
	WAVE_FORMAT_LUCENT_SX8300P             uint16 = 0x1C07 // Lucent Technologies
	WAVE_FORMAT_LUCENT_SX5363S             uint16 = 0x1C0C // Lucent Technologies
	WAVE_FORMAT_CUSEEME                    uint16 = 0x1F03 // CUSeeMe
	WAVE_FORMAT_NTCSOFT_ALF2CM_ACM         uint16 = 0x1FC4 // NTCSoft
	WAVE_FORMAT_DVM                        uint16 = 0x2000 // FAST Multimedia AG
	WAVE_FORMAT_DTS2                       uint16 = 0x2001 //
	WAVE_FORMAT_MAKEAVIS                   uint16 = 0x3313 //
	WAVE_FORMAT_DIVIO_MPEG4_AAC            uint16 = 0x4143 // Divio, Inc.
	WAVE_FORMAT_NOKIA_ADAPTIVE_MULTIRATE   uint16 = 0x4201 // Nokia
	WAVE_FORMAT_DIVIO_G726                 uint16 = 0x4243 // Divio, Inc.
	WAVE_FORMAT_LEAD_SPEECH                uint16 = 0x434C // LEAD Technologies
	WAVE_FORMAT_LEAD_VORBIS                uint16 = 0x564C // LEAD Technologies
	WAVE_FORMAT_WAVPACK_AUDIO              uint16 = 0x5756 // xiph.org
	WAVE_FORMAT_ALAC                       uint16 = 0x6C61 // Apple Lossless
	WAVE_FORMAT_OGG_VORBIS_MODE_1          uint16 = 0x674F // Ogg Vorbis
	WAVE_FORMAT_OGG_VORBIS_MODE_2          uint16 = 0x6750 // Ogg Vorbis
	WAVE_FORMAT_OGG_VORBIS_MODE_3          uint16 = 0x6751 // Ogg Vorbis
	WAVE_FORMAT_OGG_VORBIS_MODE_1_PLUS     uint16 = 0x676F // Ogg Vorbis
	WAVE_FORMAT_OGG_VORBIS_MODE_2_PLUS     uint16 = 0x6770 // Ogg Vorbis
	WAVE_FORMAT_OGG_VORBIS_MODE_3_PLUS     uint16 = 0x6771 // Ogg Vorbis
	WAVE_FORMAT_3COM_NBX                   uint16 = 0x7000 // 3COM Corp.
	WAVE_FORMAT_OPUS                       uint16 = 0x704F // Opus
	WAVE_FORMAT_FAAD_AAC                   uint16 = 0x706D //
	WAVE_FORMAT_AMR_NB                     uint16 = 0x7361 // AMR Narrowband
	WAVE_FORMAT_AMR_WB                     uint16 = 0x7362 // AMR Wideband
	WAVE_FORMAT_AMR_WP                     uint16 = 0x7363 // AMR Wideband Plus
	WAVE_FORMAT_GSM_AMR_CBR                uint16 = 0x7A21 // GSMA/3GPP
	WAVE_FORMAT_GSM_AMR_VBR_SID            uint16 = 0x7A22 // GSMA/3GPP
	WAVE_FORMAT_COMVERSE_INFOSYS_G723_1    uint16 = 0xA100 // Comverse Infosys
	WAVE_FORMAT_COMVERSE_INFOSYS_AVQSBC    uint16 = 0xA101 // Comverse Infosys
	WAVE_FORMAT_COMVERSE_INFOSYS_SBC       uint16 = 0xA102 // Comverse Infosys
	WAVE_FORMAT_SYMBOL_G729_A              uint16 = 0xA103 // Symbol Technologies
	WAVE_FORMAT_VOICEAGE_AMR_WB            uint16 = 0xA104 // VoiceAge Corp.
	WAVE_FORMAT_INGENIENT_G726             uint16 = 0xA105 // Ingenient Technologies, Inc.
	WAVE_FORMAT_MPEG4_AAC                  uint16 = 0xA106 // ISO/MPEG-4
	WAVE_FORMAT_ENCORE_G726                uint16 = 0xA107 // Encore Software
	WAVE_FORMAT_ZOLL_ASAO                  uint16 = 0xA108 // ZOLL Medical Corp.
	WAVE_FORMAT_SPEEX_VOICE                uint16 = 0xA109 // xiph.org
	WAVE_FORMAT_VIANIX_MASC                uint16 = 0xA10A // Vianix LLC
	WAVE_FORMAT_WM9_SPECTRUM_ANALYZER      uint16 = 0xA10B // Microsoft
	WAVE_FORMAT_WMF_SPECTRUM_ANAYZER       uint16 = 0xA10C // Microsoft
	WAVE_FORMAT_GSM_610                    uint16 = 0xA10D //
	WAVE_FORMAT_GSM_620                    uint16 = 0xA10E //
	WAVE_FORMAT_GSM_660                    uint16 = 0xA10F //
	WAVE_FORMAT_GSM_690                    uint16 = 0xA110 //
	WAVE_FORMAT_GSM_ADAPTIVE_MULTIRATE_WB  uint16 = 0xA111 //
	WAVE_FORMAT_POLYCOM_G722               uint16 = 0xA112 // Polycom
	WAVE_FORMAT_POLYCOM_G728               uint16 = 0xA113 // Polycom
	WAVE_FORMAT_POLYCOM_G729_A             uint16 = 0xA114 // Polycom
	WAVE_FORMAT_POLYCOM_SIREN              uint16 = 0xA115 // Polycom
	WAVE_FORMAT_GLOBAL_IP_ILBC             uint16 = 0xA116 // Global IP
	WAVE_FORMAT_RADIOTIME_TIME_SHIFT_RADIO uint16 = 0xA117 // RadioTime
	WAVE_FORMAT_NICE_ACA                   uint16 = 0xA118 // Nice Systems
	WAVE_FORMAT_NICE_ADPCM                 uint16 = 0xA119 // Nice Systems
	WAVE_FORMAT_VOCORD_G721                uint16 = 0xA11A // Vocord Telecom
	WAVE_FORMAT_VOCORD_G726                uint16 = 0xA11B // Vocord Telecom
	WAVE_FORMAT_VOCORD_G722_1              uint16 = 0xA11C // Vocord Telecom
	WAVE_FORMAT_VOCORD_G728                uint16 = 0xA11D // Vocord Telecom
	WAVE_FORMAT_VOCORD_G729                uint16 = 0xA11E // Vocord Telecom
	WAVE_FORMAT_VOCORD_G729_A              uint16 = 0xA11F // Vocord Telecom
	WAVE_FORMAT_VOCORD_G723_1              uint16 = 0xA120 // Vocord Telecom
	WAVE_FORMAT_VOCORD_LBC                 uint16 = 0xA121 // Vocord Telecom
	WAVE_FORMAT_NICE_G728                  uint16 = 0xA122 // Nice Systems
	WAVE_FORMAT_FRACE_TELECOM_G729         uint16 = 0xA123 // France Telecom
	WAVE_FORMAT_CODIAN                     uint16 = 0xA124 // CODIAN
	WAVE_FORMAT_DOLBY_AC4                  uint16 = 0xAC40 // Dolby AC-4
	WAVE_FORMAT_FLAC                       uint16 = 0xF1AC // flac.sourceforge.net
	WAVE_FORMAT_EXTENSIBLE                 uint16 = 0xFFFE // Microsoft
	WAVE_FORMAT_DEVELOPMENT                uint16 = 0xFFFF
)

const (
	AUDCLNT_STREAMFLAGS_CROSSPROCESS        = 0x00010000
	AUDCLNT_STREAMFLAGS_LOOPBACK            = 0x00020000
	AUDCLNT_STREAMFLAGS_EVENTCALLBACK       = 0x00040000
	AUDCLNT_STREAMFLAGS_NOPERSIST           = 0x00080000
	AUDCLNT_STREAMFLAGS_RATEADJUST          = 0x00100000
	AUDCLNT_STREAMFLAGS_AUTOCONVERTPCM      = 0x80000000
	AUDCLNT_STREAMFLAGS_SRC_DEFAULT_QUALITY = 0x08000000
)

const (
	AUDCLNT_SESSIONFLAGS_EXPIREWHENUNOWNED       = 0x10000000
	AUDCLNT_SESSIONFLAGS_DISPLAY_HIDE            = 0x20000000
	AUDCLNT_SESSIONFLAGS_DISPLAY_HIDEWHENEXPIRED = 0x40000000
)

const (
	AUDCLNT_BUFFERFLAGS_DATA_DISCONTINUITY uint32 = iota
	AUDCLNT_BUFFERFLAGS_SILENT
	AUDCLNT_BUFFERFLAGS_TIMESTAMP_ERROR
)

type AUDIO_STREAM_CATEGORY uint32

const (
	AudioCategory_Other AUDIO_STREAM_CATEGORY = iota
	AudioCategory_ForegroundOnlyMedia
	AudioCategory_BackgroundCapableMedia
	AudioCategory_Communications
	AudioCategory_Alerts
	AudioCategory_SoundEffects
	AudioCategory_GameEffects
	AudioCategory_GameMedia
	AudioCategory_GameChat
	AudioCategory_Speech
	AudioCategory_Movie
	AudioCategory_Media
	AudioCategory_FarFieldSpeech
	AudioCategory_UniformSpeech
	AudioCategory_VoiceTyping
)

func ToType[T IAudioClient | IAudioCaptureClient | IAudioRenderClient](v unsafe.Pointer) *T {
	return (*T)(v)
}

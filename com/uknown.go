package com

import "golang.org/x/sys/windows"

// MIDL_INTERFACE("00000000-0000-0000-C000-000000000046")
var _IID_IUnknown = windows.GUID{Data1: 0x00000000, Data2: 0x0000, Data3: 0x0000, Data4: [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

func IID_IUnknown() windows.GUID {
	return _IID_IUnknown
}

type IUnknown struct {
	vtbl *IUnknownVtbl
}

type IUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

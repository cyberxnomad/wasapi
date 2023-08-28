package com

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type PROPVARIANT struct {
	Vt         uint16 // Value type tag.
	WReserved1 uint16
	WReserved2 uint16
	WReserved3 uint16
	Val        uint64
	_          uint64
}

func (self *PROPVARIANT) PwszVal() *uint16 {
	return *(**uint16)(unsafe.Pointer(&self.Val))
}

func (self *PROPVARIANT) PwszValString() string {
	return windows.UTF16PtrToString(self.PwszVal())
}

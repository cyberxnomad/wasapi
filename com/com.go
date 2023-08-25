package com

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	mod_ole32            = windows.NewLazyDLL("ole32.dll")
	procCoCreateInstance = mod_ole32.NewProc("CoCreateInstance")
	procCoInitializeEx   = mod_ole32.NewProc("CoInitializeEx")
	procCoUninitialize   = mod_ole32.NewProc("CoUninitialize")
	procCoTaskMemFree    = mod_ole32.NewProc("CoTaskMemFree")
)

type HRESULT = uint32

type _PROPVARIANT_Union struct {
	Union [2]uint64
}

func (self *_PROPVARIANT_Union) PwszVal() *uint16 {
	return *(**uint16)(unsafe.Pointer(self))
}

type PROPVARIANT struct {
	Vt         uint16 // Value type tag.
	WReserved1 uint16
	WReserved2 uint16
	WReserved3 uint16
	_PROPVARIANT_Union
}

type PROPERTYKEY struct {
	Fmtid windows.GUID
	Pid   uint32
}

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

// 初始化 COM 库以供调用线程使用，设置线程的并发模型，并根据需要为线程创建一个新单元。
func CoInitializeEx(reserved uintptr, coInitFlag uint32) (err error) {
	r, _, _ := syscall.SyscallN(procCoInitializeEx.Addr(),
		reserved,
		uintptr(coInitFlag),
	)

	if HRESULT(r) != HRESULT(windows.S_OK) {
		err = fmt.Errorf("com::CoInitializeEx failed with code: 0x%08X", HRESULT(r))
		return
	}

	return
}

// 释放先前通过调用 CoTaskMemAlloc 或 CoTaskMemRealloc 函数分配的任务内存块。
func CoTaskMemFree(address unsafe.Pointer) {
	syscall.SyscallN(procCoTaskMemFree.Addr(), uintptr(address))
}

// 关闭当前线程上的 COM 库，卸载线程加载的所有 DLL，释放线程维护的任何其他资源，并强制关闭线程上的所有 RPC 连接。
func CoUninitialize() {
	syscall.SyscallN(procCoUninitialize.Addr())
}

// 创建并默认初始化与指定 CLSID 关联的类的单个对象。
func CoCreateInstance(
	clsid *windows.GUID,
	unkOuter *IUnknown,
	clsContext uint32,
	iid *windows.GUID,
) (v unsafe.Pointer, err error) {
	r, _, _ := procCoCreateInstance.Call(
		uintptr(unsafe.Pointer(clsid)),
		uintptr(unsafe.Pointer(unkOuter)),
		uintptr(clsContext),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&v)),
	)

	if HRESULT(r) != HRESULT(windows.S_OK) {
		err = fmt.Errorf("com::CoCreateInstance failed with code: 0x%08X", HRESULT(r))
		return
	}

	return
}

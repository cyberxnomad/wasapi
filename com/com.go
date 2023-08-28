package com

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modole32             = windows.NewLazyDLL("ole32.dll")
	procCoCreateInstance = modole32.NewProc("CoCreateInstance")
	procCoInitializeEx   = modole32.NewProc("CoInitializeEx")
	procCoUninitialize   = modole32.NewProc("CoUninitialize")
	procCoTaskMemFree    = modole32.NewProc("CoTaskMemFree")
)

type HRESULT = uint32

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
	r, _, _ := syscall.SyscallN(procCoCreateInstance.Addr(),
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

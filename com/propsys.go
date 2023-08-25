package com

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// MIDL_INTERFACE("886d8eeb-8cf2-4446-8d02-cdba1dbdcf99")
var _IID_IPropertyStore = windows.GUID{Data1: 0x886d8eeb, Data2: 0x8cf2, Data3: 0x4446, Data4: [8]byte{0x8d, 0x02, 0xcd, 0xba, 0x1d, 0xbd, 0xcf, 0x99}}

func IID_IPropertyStore() windows.GUID {
	return _IID_IPropertyStore
}

const (
	STGM_READ      = 0x00000000
	STGM_WRITE     = 0x00000001
	STGM_READWRITE = 0x00000002
)

type IPropertyStore struct {
	vtbl *_IPropertyStoreVtbl
}

type _IPropertyStoreVtbl struct {
	IUnknownVtbl
	GetCount uintptr
	GetAt    uintptr
	GetValue uintptr
	SetValue uintptr
	Commit   uintptr
}

func (self *IPropertyStore) Release() (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Release, uintptr(unsafe.Pointer(self)))

	if HRESULT(r) != HRESULT(windows.S_OK) {
		err = fmt.Errorf("IPropertyStore::Release failed with code: 0x%08X", HRESULT(r))
		return
	}

	return
}

// 此方法返回附加到文件的属性数的计数。
func (self *IPropertyStore) GetCount() (cProps uint32, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetCount, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&cProps)),
	)

	if HRESULT(r) != HRESULT(windows.S_OK) {
		err = fmt.Errorf("IPropertyStore::GetCount failed with code: 0x%08X", HRESULT(r))
		return
	}

	return
}

// 从项的属性数组中获取属性键。
func (self *IPropertyStore) GetAt(iProp uint32) (key *PROPERTYKEY, err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.GetAt, uintptr(unsafe.Pointer(self)),
		uintptr(iProp),
		uintptr(unsafe.Pointer(&key)),
	)

	if HRESULT(r) != HRESULT(windows.S_OK) {
		err = fmt.Errorf("IPropertyStore::GetAt failed with code: 0x%08X", HRESULT(r))
		return
	}

	return
}

// 此方法检索特定属性的数据。
func (self *IPropertyStore) GetValue(key PROPERTYKEY) (propVar PROPVARIANT, err error) {

	r, _, _ := syscall.SyscallN(self.vtbl.GetValue, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&key)),
		uintptr(unsafe.Pointer(&propVar)),
	)

	if HRESULT(r) != HRESULT(windows.S_OK) && HRESULT(r) != HRESULT(windows.INPLACE_S_TRUNCATED) {
		err = fmt.Errorf("IPropertyStore::GetValue failed with code: 0x%08X", HRESULT(r))
		return
	}

	return
}

// 此方法设置属性值或替换或删除现有值。
func (self *IPropertyStore) SetValue(key PROPERTYKEY, propVar PROPVARIANT) (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.SetValue, uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&key)),
		uintptr(unsafe.Pointer(&propVar)),
	)

	if HRESULT(r) != HRESULT(windows.S_OK) && HRESULT(r) != HRESULT(windows.INPLACE_S_TRUNCATED) {
		err = fmt.Errorf("IPropertyStore::SetValue failed with code: 0x%08X", HRESULT(r))
		return
	}

	return
}

// 进行更改后，此方法将保存更改。
func (self *IPropertyStore) Commit() (err error) {
	r, _, _ := syscall.SyscallN(self.vtbl.Commit, uintptr(unsafe.Pointer(self)))

	if HRESULT(r) != HRESULT(windows.S_OK) {
		err = fmt.Errorf("IPropertyStore::Commit failed with code: 0x%08X", HRESULT(r))
		return
	}

	return
}

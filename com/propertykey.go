package com

import "golang.org/x/sys/windows"

type PROPERTYKEY struct {
	Fmtid windows.GUID
	Pid   uint32
}

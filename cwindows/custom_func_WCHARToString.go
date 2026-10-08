package cwindows

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

func WCHARToString(value []WCHAR) string {
	if len(value) == 0 {
		return ""
	}

	return windows.UTF16ToString(
		unsafe.Slice(
			(*uint16)(unsafe.Pointer(&value[0])),
			len(value),
		),
	)
}

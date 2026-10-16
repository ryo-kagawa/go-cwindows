// Code generated DO NOT EDIT

package cwindows

import (
	"unsafe"

	processtrace "github.com/ryo-kagawa/go-cwindows/dll/sechost/ProcessTrace"
)

func InitSechostProcessTrace() {
	processtrace.InitSechostProcessTrace()
}

func SechostProcessTrace(HandleArray *PROCESSTRACE_HANDLE, HandleCount ULONG, StartTime LPFILETIME, EndTime LPFILETIME) ULONG {
	r1 := processtrace.SechostProcessTrace(
		uintptr(unsafe.Pointer(HandleArray)),
		uintptr(HandleCount),
		uintptr(unsafe.Pointer(StartTime)),
		uintptr(unsafe.Pointer(EndTime)),
	)
	return ULONG(r1)
}

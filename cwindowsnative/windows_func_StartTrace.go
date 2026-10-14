// Code generated DO NOT EDIT

package cwindowsnative

import (
	"unsafe"

	starttrace "github.com/ryo-kagawa/go-cwindows/dll/sechost/StartTrace"
)

func InitSechostStartTrace() {
	starttrace.InitSechostStartTrace()
}

func StartTrace[TProperties PEVENT_TRACE_PROPERTIES | PEVENT_TRACE_PROPERTIES_V2](TraceId *CONTROLTRACE_ID, InstanceName LPCWSTR, Properties TProperties) ULONG {
	r1 := starttrace.StartTrace(
		uintptr(unsafe.Pointer(TraceId)),
		uintptr(unsafe.Pointer(InstanceName)),
		uintptr(unsafe.Pointer(Properties)),
	)
	return ULONG(r1)
}

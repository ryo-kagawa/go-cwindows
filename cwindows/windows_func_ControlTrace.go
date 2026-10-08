package cwindows

import (
	"unsafe"

	controltrace "github.com/ryo-kagawa/go-cwindows/dll/sechost/ControlTrace"
)

func InitSechostControlTrace() {
	controltrace.InitSechostControlTrace()
}

func ControlTrace[TProperties PEVENT_TRACE_PROPERTIES | PEVENT_TRACE_PROPERTIES_V2](TraceId CONTROLTRACE_ID, InstanceName LPCWSTR, Properties TProperties, ControlCode ControlTraceCode) ULONG {
	r1 := controltrace.ControlTrace(
		uintptr(TraceId),
		uintptr(unsafe.Pointer(InstanceName)),
		uintptr(unsafe.Pointer(Properties)),
		uintptr(ControlCode),
	)
	return ULONG(r1)
}

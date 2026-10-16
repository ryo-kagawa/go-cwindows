// Code generated DO NOT EDIT

package cwindows

import (
	"unsafe"

	enabletraceex2 "github.com/ryo-kagawa/go-cwindows/dll/sechost/EnableTraceEx2"
)

func InitSechostEnableTraceEx2() {
	enabletraceex2.InitSechostEnableTraceEx2()
}

func SechostEnableTraceEx2(TraceId CONTROLTRACE_ID, ProviderId LPCGUID, ControlCode EventControlCode, Level TraceLevel, MatchAnyKeyword ULONGLONG, MatchAllKeyword ULONGLONG, Timeout ULONG, EnableParameters PENABLE_TRACE_PARAMETERS) ULONG {
	r1 := enabletraceex2.SechostEnableTraceEx2(
		uintptr(TraceId),
		uintptr(unsafe.Pointer(ProviderId)),
		uintptr(ControlCode),
		uintptr(Level),
		uintptr(MatchAnyKeyword),
		uintptr(MatchAllKeyword),
		uintptr(Timeout),
		uintptr(unsafe.Pointer(EnableParameters)),
	)
	return ULONG(r1)
}

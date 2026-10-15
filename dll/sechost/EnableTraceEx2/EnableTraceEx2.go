// Code generated DO NOT EDIT

package enabletraceex2

import sechost "github.com/ryo-kagawa/go-cwindows/dll/sechost"

var impl = func(TraceId uintptr, ProviderId uintptr, ControlCode uintptr, Level uintptr, MatchAnyKeyword uintptr, MatchAllKeyword uintptr, Timeout uintptr, EnableParameters uintptr) uintptr {
	panic("SechostEnableTraceEx2 not initialized. Call InitSechostEnableTraceEx2() first.")
}

func InitSechostEnableTraceEx2() {
	proc := sechost.ModSechost.NewProc("EnableTraceEx2")
	impl = func(TraceId uintptr, ProviderId uintptr, ControlCode uintptr, Level uintptr, MatchAnyKeyword uintptr, MatchAllKeyword uintptr, Timeout uintptr, EnableParameters uintptr) uintptr {
		r1, _, _ := proc.Call(
			TraceId,
			ProviderId,
			ControlCode,
			Level,
			MatchAnyKeyword,
			MatchAllKeyword,
			Timeout,
			EnableParameters,
		)
		return r1
	}
}

func SechostEnableTraceEx2(TraceId uintptr, ProviderId uintptr, ControlCode uintptr, Level uintptr, MatchAnyKeyword uintptr, MatchAllKeyword uintptr, Timeout uintptr, EnableParameters uintptr) uintptr {
	r1 := impl(
		TraceId,
		ProviderId,
		ControlCode,
		Level,
		MatchAnyKeyword,
		MatchAllKeyword,
		Timeout,
		EnableParameters,
	)
	return r1
}

// Code generated DO NOT EDIT

package controltrace

import sechost "github.com/ryo-kagawa/go-cwindows/dll/sechost"

var impl = func(TraceId uintptr, InstanceName uintptr, Properties uintptr, ControlCode uintptr) uintptr {
	panic("SechostControlTrace not initialized. Call InitSechostControlTrace() first.")
}

func InitSechostControlTrace() {
	proc := sechost.ModSechost.NewProc("ControlTraceW")
	impl = func(TraceId uintptr, InstanceName uintptr, Properties uintptr, ControlCode uintptr) uintptr {
		r1, _, _ := proc.Call(
			TraceId,
			InstanceName,
			Properties,
			ControlCode,
		)
		return r1
	}
}

func SechostControlTrace(TraceId uintptr, InstanceName uintptr, Properties uintptr, ControlCode uintptr) uintptr {
	r1 := impl(
		TraceId,
		InstanceName,
		Properties,
		ControlCode,
	)
	return r1
}

// Code generated DO NOT EDIT

package starttrace

import sechost "github.com/ryo-kagawa/go-cwindows/dll/sechost"

var impl = func(TraceId uintptr, InstanceName uintptr, Properties uintptr) uintptr {
	panic("SechostStartTrace not initialized. Call InitSechostStartTrace() first.")
}

func InitSechostStartTrace() {
	proc := sechost.ModSechost.NewProc("StartTraceW")
	impl = func(TraceId uintptr, InstanceName uintptr, Properties uintptr) uintptr {
		r1, _, _ := proc.Call(
			TraceId,
			InstanceName,
			Properties,
		)
		return r1
	}
}

func SechostStartTrace(TraceId uintptr, InstanceName uintptr, Properties uintptr) uintptr {
	r1 := impl(
		TraceId,
		InstanceName,
		Properties,
	)
	return r1
}

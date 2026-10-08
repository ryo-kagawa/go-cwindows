package starttrace

import "github.com/ryo-kagawa/go-cwindows/dll/sechost"

var impl = func(TraceId uintptr, InstanceName uintptr, Properties uintptr) uintptr {
	panic("StartTrace not initialized. Call InitSechostStartTrace() first.")
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

func StartTrace(TraceId uintptr, InstanceName uintptr, Properties uintptr) uintptr {
	r1 := impl(
		TraceId,
		InstanceName,
		Properties,
	)
	return r1
}

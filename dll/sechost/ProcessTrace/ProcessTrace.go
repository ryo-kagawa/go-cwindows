// Code generated DO NOT EDIT

package processtrace

import sechost "github.com/ryo-kagawa/go-cwindows/dll/sechost"

var impl = func(HandleArray uintptr, HandleCount uintptr, StartTime uintptr, EndTime uintptr) uintptr {
	panic("ProcessTrace not initialized. Call InitSechostProcessTrace() first.")
}

func InitSechostProcessTrace() {
	proc := sechost.ModSechost.NewProc("ProcessTrace")
	impl = func(HandleArray uintptr, HandleCount uintptr, StartTime uintptr, EndTime uintptr) uintptr {
		r1, _, _ := proc.Call(
			HandleArray,
			HandleCount,
			StartTime,
			EndTime,
		)
		return r1
	}
}

func ProcessTrace(HandleArray uintptr, HandleCount uintptr, StartTime uintptr, EndTime uintptr) uintptr {
	r1 := impl(
		HandleArray,
		HandleCount,
		StartTime,
		EndTime,
	)
	return r1
}

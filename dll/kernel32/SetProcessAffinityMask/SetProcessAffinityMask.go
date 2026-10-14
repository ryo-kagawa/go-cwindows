// Code generated DO NOT EDIT

package setprocessaffinitymask

import kernel32 "github.com/ryo-kagawa/go-cwindows/dll/kernel32"

var impl = func(hProcess uintptr, dwProcessAffinityMask uintptr) (uintptr, error) {
	panic("SetProcessAffinityMask not initialized. Call InitKernel32SetProcessAffinityMask() first.")
}

func InitKernel32SetProcessAffinityMask() {
	proc := kernel32.ModKernel32.NewProc("SetProcessAffinityMask")
	impl = func(hProcess uintptr, dwProcessAffinityMask uintptr) (uintptr, error) {
		r1, _, lastErr := proc.Call(
			hProcess,
			dwProcessAffinityMask,
		)
		return r1, lastErr
	}
}

func SetProcessAffinityMask(hProcess uintptr, dwProcessAffinityMask uintptr) (uintptr, error) {
	r1, lastErr := impl(
		hProcess,
		dwProcessAffinityMask,
	)
	return r1, lastErr
}

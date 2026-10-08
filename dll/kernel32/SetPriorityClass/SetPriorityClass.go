package setpriorityclass

import "github.com/ryo-kagawa/go-cwindows/dll/kernel32"

var impl = func(hProcess uintptr, dwPriorityClass uintptr) (uintptr, error) {
	panic("SetPriorityClass not initialized. Call InitKernel32SetPriorityClass() first.")
}

func InitKernel32SetPriorityClass() {
	proc := kernel32.ModKernel32.NewProc("SetPriorityClass")
	impl = func(hProcess uintptr, dwPriorityClass uintptr) (uintptr, error) {
		r1, _, lastErr := proc.Call(
			hProcess,
			dwPriorityClass,
		)
		return r1, lastErr
	}
}

func SetPriorityClass(hProcess uintptr, dwPriorityClass uintptr) (uintptr, error) {
	r1, lastErr := impl(
		hProcess,
		dwPriorityClass,
	)
	return r1, lastErr
}

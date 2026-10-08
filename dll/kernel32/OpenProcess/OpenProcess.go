package openprocess

import "github.com/ryo-kagawa/go-cwindows/dll/kernel32"

var impl = func(dwDesiredAccess uintptr, bInheritHandle uintptr, dwProcessId uintptr) (uintptr, error) {
	panic("OpenProcess not initialized. Call InitKernel32OpenProcess() first.")
}

func InitKernel32OpenProcess() {
	proc := kernel32.ModKernel32.NewProc("OpenProcess")
	impl = func(dwDesiredAccess uintptr, bInheritHandle uintptr, dwProcessId uintptr) (uintptr, error) {
		r1, _, lastErr := proc.Call(
			dwDesiredAccess,
			bInheritHandle,
			dwProcessId,
		)
		return r1, lastErr
	}
}

func OpenProcess(dwDesiredAccess uintptr, bInheritHandle uintptr, dwProcessId uintptr) (uintptr, error) {
	r1, lastErr := impl(
		dwDesiredAccess,
		bInheritHandle,
		dwProcessId,
	)
	return r1, lastErr
}

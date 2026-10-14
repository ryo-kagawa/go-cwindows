// Code generated DO NOT EDIT

package createtoolhelp32snapshot

import kernel32 "github.com/ryo-kagawa/go-cwindows/dll/kernel32"

var impl = func(dwFlags uintptr, th32ProcessID uintptr) (uintptr, error) {
	panic("CreateToolhelp32Snapshot not initialized. Call InitKernel32CreateToolhelp32Snapshot() first.")
}

func InitKernel32CreateToolhelp32Snapshot() {
	proc := kernel32.ModKernel32.NewProc("CreateToolhelp32Snapshot")
	impl = func(dwFlags uintptr, th32ProcessID uintptr) (uintptr, error) {
		r1, _, lastErr := proc.Call(
			dwFlags,
			th32ProcessID,
		)
		return r1, lastErr
	}
}

func CreateToolhelp32Snapshot(dwFlags uintptr, th32ProcessID uintptr) (uintptr, error) {
	r1, lastErr := impl(
		dwFlags,
		th32ProcessID,
	)
	return r1, lastErr
}

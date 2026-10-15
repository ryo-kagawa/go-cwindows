// Code generated DO NOT EDIT

package process32next

import kernel32 "github.com/ryo-kagawa/go-cwindows/dll/kernel32"

var impl = func(hSnapshot uintptr, lppe uintptr) (uintptr, error) {
	panic("Kernel32Process32Next not initialized. Call InitKernel32Process32Next() first.")
}

func InitKernel32Process32Next() {
	proc := kernel32.ModKernel32.NewProc("Process32NextW")
	impl = func(hSnapshot uintptr, lppe uintptr) (uintptr, error) {
		r1, _, lastErr := proc.Call(
			hSnapshot,
			lppe,
		)
		return r1, lastErr
	}
}

func Kernel32Process32Next(hSnapshot uintptr, lppe uintptr) (uintptr, error) {
	r1, lastErr := impl(
		hSnapshot,
		lppe,
	)
	return r1, lastErr
}

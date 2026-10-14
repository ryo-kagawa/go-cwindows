// Code generated DO NOT EDIT

package process32first

import kernel32 "github.com/ryo-kagawa/go-cwindows/dll/kernel32"

var impl = func(hSnapshot uintptr, lppe uintptr) (uintptr, error) {
	panic("Process32First not initialized. Call InitKernel32Process32First() first.")
}

func InitKernel32Process32First() {
	proc := kernel32.ModKernel32.NewProc("Process32FirstW")
	impl = func(hSnapshot uintptr, lppe uintptr) (uintptr, error) {
		r1, _, lastErr := proc.Call(
			hSnapshot,
			lppe,
		)
		return r1, lastErr
	}
}

func Process32First(hSnapshot uintptr, lppe uintptr) (uintptr, error) {
	r1, lastErr := impl(
		hSnapshot,
		lppe,
	)
	return r1, lastErr
}

// Code generated DO NOT EDIT

package closehandle

import kernel32 "github.com/ryo-kagawa/go-cwindows/dll/kernel32"

var impl = func(hObject uintptr) (uintptr, error) {
	panic("Kernel32CloseHandle not initialized. Call InitKernel32CloseHandle() first.")
}

func InitKernel32CloseHandle() {
	proc := kernel32.ModKernel32.NewProc("CloseHandle")
	impl = func(hObject uintptr) (uintptr, error) {
		r1, _, lastErr := proc.Call(
			hObject,
		)
		return r1, lastErr
	}
}

func Kernel32CloseHandle(hObject uintptr) (uintptr, error) {
	r1, lastErr := impl(
		hObject,
	)
	return r1, lastErr
}

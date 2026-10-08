package closehandle

import "github.com/ryo-kagawa/go-cwindows/dll/kernel32"

var impl = func(hObject uintptr) (uintptr, error) {
	panic("CloseHandle not initialized. Call InitKernel32CloseHandle() first.")
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

func CloseHandle(hObject uintptr) (uintptr, error) {
	r1, lastErr := impl(
		hObject,
	)
	return r1, lastErr
}

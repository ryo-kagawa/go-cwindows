// Code generated DO NOT EDIT

package cwindows

import closehandle "github.com/ryo-kagawa/go-cwindows/dll/kernel32/CloseHandle"

func InitKernel32CloseHandle() {
	closehandle.InitKernel32CloseHandle()
}

func Kernel32CloseHandle(hObject HANDLE) (BOOL, error) {
	r1, lastErr := closehandle.Kernel32CloseHandle(
		uintptr(hObject),
	)
	return BOOL(r1), lastErr
}

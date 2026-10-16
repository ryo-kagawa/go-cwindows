// Code generated DO NOT EDIT

package cwindowsnative

import setpriorityclass "github.com/ryo-kagawa/go-cwindows/dll/kernel32/SetPriorityClass"

func InitKernel32SetPriorityClass() {
	setpriorityclass.InitKernel32SetPriorityClass()
}

func Kernel32SetPriorityClass(hProcess HANDLE, dwPriorityClass DWORD) (BOOL, error) {
	r1, lastErr := setpriorityclass.Kernel32SetPriorityClass(
		uintptr(hProcess),
		uintptr(dwPriorityClass),
	)
	return BOOL(r1), lastErr
}

// Code generated DO NOT EDIT

package cwindows

import openprocess "github.com/ryo-kagawa/go-cwindows/dll/kernel32/OpenProcess"

func InitKernel32OpenProcess() {
	openprocess.InitKernel32OpenProcess()
}

func Kernel32OpenProcess(dwDesiredAccess ProcessAccess, bInheritHandle BOOL, dwProcessId DWORD) (HANDLE, error) {
	r1, lastErr := openprocess.OpenProcess(
		uintptr(dwDesiredAccess),
		uintptr(bInheritHandle),
		uintptr(dwProcessId),
	)
	return HANDLE(r1), lastErr
}

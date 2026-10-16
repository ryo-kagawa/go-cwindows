// Code generated DO NOT EDIT

package cwindowsnative

import createtoolhelp32snapshot "github.com/ryo-kagawa/go-cwindows/dll/kernel32/CreateToolhelp32Snapshot"

func InitKernel32CreateToolhelp32Snapshot() {
	createtoolhelp32snapshot.InitKernel32CreateToolhelp32Snapshot()
}

func Kernel32CreateToolhelp32Snapshot(dwFlags DWORD, th32ProcessID DWORD) (HANDLE, error) {
	r1, lastErr := createtoolhelp32snapshot.Kernel32CreateToolhelp32Snapshot(
		uintptr(dwFlags),
		uintptr(th32ProcessID),
	)
	return HANDLE(r1), lastErr
}

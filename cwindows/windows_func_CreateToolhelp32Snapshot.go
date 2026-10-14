// Code generated DO NOT EDIT

package cwindows

import createtoolhelp32snapshot "github.com/ryo-kagawa/go-cwindows/dll/kernel32/CreateToolhelp32Snapshot"

func InitKernel32CreateToolhelp32Snapshot() {
	createtoolhelp32snapshot.InitKernel32CreateToolhelp32Snapshot()
}

func CreateToolhelp32Snapshot(dwFlags TH32CS, th32ProcessID DWORD) (HANDLE, error) {
	r1, lastErr := createtoolhelp32snapshot.CreateToolhelp32Snapshot(
		uintptr(dwFlags),
		uintptr(th32ProcessID),
	)
	return HANDLE(r1), lastErr
}

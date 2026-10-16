// Code generated DO NOT EDIT

package cwindowsnative

import (
	"unsafe"

	process32first "github.com/ryo-kagawa/go-cwindows/dll/kernel32/Process32First"
)

func InitKernel32Process32First() {
	process32first.InitKernel32Process32First()
}

func Kernel32Process32First(hSnapshot HANDLE, lppe LPPROCESSENTRY32) (BOOL, error) {
	r1, lastErr := process32first.Kernel32Process32First(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(lppe)),
	)
	return BOOL(r1), lastErr
}

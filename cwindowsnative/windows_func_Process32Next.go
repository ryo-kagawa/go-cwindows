package cwindowsnative

import (
	"unsafe"

	process32next "github.com/ryo-kagawa/go-cwindows/dll/kernel32/Process32Next"
)

func InitKernel32Process32Next() {
	process32next.InitKernel32Process32Next()
}

func Process32Next(hSnapshot HANDLE, lppe LPPROCESSENTRY32) (BOOL, error) {
	r1, lastErr := process32next.Process32Next(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(lppe)),
	)
	return BOOL(r1), lastErr
}

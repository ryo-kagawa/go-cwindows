package cwindows

import setprocessaffinitymask "github.com/ryo-kagawa/go-cwindows/dll/kernel32/SetProcessAffinityMask"

func InitKernel32SetProcessAffinityMask() {
	setprocessaffinitymask.InitKernel32SetProcessAffinityMask()
}

func SetProcessAffinityMask(hProcess HANDLE, dwProcessAffinityMask DWORD_PTR) (BOOL, error) {
	r1, lastErr := setprocessaffinitymask.SetProcessAffinityMask(
		uintptr(hProcess),
		uintptr(dwProcessAffinityMask),
	)
	return BOOL(r1), lastErr
}

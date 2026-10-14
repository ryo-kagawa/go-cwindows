// Code generated DO NOT EDIT

package cwindows

type PROCESSENTRY32 struct {
	Size            DWORD
	Usage           DWORD
	ProcessID       DWORD
	DefaultHeapID   ULONG_PTR
	ModuleID        DWORD
	Threads         DWORD
	ParentProcessID DWORD
	PriClassBase    LONG
	Flags           DWORD
	ExeFile         [MAX_PATH]WCHAR
}

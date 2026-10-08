package cwindows

type FILETIME struct {
	LowDateTime  DWORD
	HighDateTime DWORD
}
type LPFILETIME *FILETIME

package cwindowsnative

type FILETIME struct {
	LowDateTime  DWORD
	HighDateTime DWORD
}
type LPFILETIME *FILETIME

// Code generated DO NOT EDIT

package cwindows

func UTF16ToWCHAR(value []uint16) []WCHAR {
	result := make([]WCHAR, len(value))
	for i, c := range value {
		result[i] = WCHAR(c)
	}
	return result
}

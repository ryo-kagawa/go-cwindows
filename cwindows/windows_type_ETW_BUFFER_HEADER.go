// Code generated DO NOT EDIT

package cwindows

type ETW_BUFFER_HEADER struct {
	Reserved1     [4]ULONG
	TimeStamp     LARGE_INTEGER
	Reserved2     [2]ULONG
	ClientContext ETW_BUFFER_CONTEXT
	Reserved3     ULONG
	FilledBytes   ULONG
	Reserved4     [5]ULONG
}

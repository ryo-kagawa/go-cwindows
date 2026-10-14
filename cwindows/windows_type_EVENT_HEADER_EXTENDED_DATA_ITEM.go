// Code generated DO NOT EDIT

package cwindows

type EVENT_HEADER_EXTENDED_DATA_ITEM struct {
	Reserved1 USHORT
	ExtType   EventHeaderExtType
	_         [2]byte
	DataSize  USHORT
	DataPtr   ULONGLONG
}

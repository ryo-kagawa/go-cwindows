package cwindowsnative

type EVENT_HEADER_EXTENDED_DATA_ITEM struct {
	Reserved1 USHORT
	ExtType   USHORT
	_         [2]byte
	DataSize  USHORT
	DataPtr   ULONGLONG
}
type PEVENT_HEADER_EXTENDED_DATA_ITEM *EVENT_HEADER_EXTENDED_DATA_ITEM

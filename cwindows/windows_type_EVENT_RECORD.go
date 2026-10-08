package cwindows

type EVENT_RECORD struct {
	EventHeader       EVENT_HEADER
	BufferContext     ETW_BUFFER_CONTEXT
	ExtendedDataCount USHORT
	UserDataLength    USHORT
	ExtendedData      PEVENT_HEADER_EXTENDED_DATA_ITEM
	UserData          PVOID
	UserContext       PVOID
}
type PEVENT_RECORD *EVENT_RECORD

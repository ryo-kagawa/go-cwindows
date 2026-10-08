package cwindows

type EVENT_HEADER struct {
	Size            USHORT
	HeaderType      USHORT
	Flags           EventHeaderFlag
	EventProperty   EventHeaderProperty
	ThreadId        ULONG
	ProcessId       ULONG
	TimeStamp       LARGE_INTEGER
	ProviderId      GUID
	EventDescriptor EVENT_DESCRIPTOR
	DUMMYUNIONNAME  [8]byte
	ActivityId      GUID
}

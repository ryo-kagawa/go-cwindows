// Code generated DO NOT EDIT

package cwindowsnative

type EVENT_HEADER struct {
	Size            USHORT
	HeaderType      USHORT
	Flags           USHORT
	EventProperty   USHORT
	ThreadId        ULONG
	ProcessId       ULONG
	TimeStamp       LARGE_INTEGER
	ProviderId      GUID
	EventDescriptor EVENT_DESCRIPTOR
	DUMMYUNIONNAME  [8]byte
	ActivityId      GUID
}

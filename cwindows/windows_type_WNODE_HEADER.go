package cwindows

type WNODE_HEADER struct {
	BufferSize    ULONG
	ProviderId    ULONG
	union1        [8]byte
	union2        [8]byte
	Guid          GUID
	ClientContext ETWClockType
	Flags         WnodeFlag
}

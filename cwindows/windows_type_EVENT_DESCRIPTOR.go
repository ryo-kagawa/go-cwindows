package cwindows

type EVENT_DESCRIPTOR struct {
	Id      ProcessEventId
	Version UCHAR
	Channel UCHAR
	Level   WinEventLevel
	Opcode  WinEventOpCode
	Task    USHORT
	Keyword ULONGLONG
}

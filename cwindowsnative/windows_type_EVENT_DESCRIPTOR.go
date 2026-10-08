package cwindowsnative

type EVENT_DESCRIPTOR struct {
	Id      USHORT
	Version UCHAR
	Channel UCHAR
	Level   UCHAR
	Opcode  UCHAR
	Task    USHORT
	Keyword ULONGLONG
}

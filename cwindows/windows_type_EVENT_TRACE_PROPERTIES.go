// Code generated DO NOT EDIT

package cwindows

type EVENT_TRACE_PROPERTIES struct {
	Wnode               WNODE_HEADER
	BufferSize          ULONG
	MinimumBuffers      ULONG
	MaximumBuffers      ULONG
	MaximumFileSize     ULONG
	LogFileMode         LoggingMode
	FlushTimer          ULONG
	EnableFlags         EventTraceEnableFlag
	DUMMYUNIONNAME      [4]byte
	NumberOfBuffers     ULONG
	FreeBuffers         ULONG
	EventsLost          ULONG
	BuffersWritten      ULONG
	LogBuffersLost      ULONG
	RealTimeBuffersLost ULONG
	LoggerThreadId      HANDLE
	LogFileNameOffset   ULONG
	LoggerNameOffset    ULONG
}

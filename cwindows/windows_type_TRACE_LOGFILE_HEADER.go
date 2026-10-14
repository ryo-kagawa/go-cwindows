// Code generated DO NOT EDIT

package cwindows

type TRACE_LOGFILE_HEADER struct {
	BufferSize         ULONG
	DUMMYUNIONNAME     [4]byte
	ProviderVersion    ULONG
	NumberOfProcessors ULONG
	EndTime            LARGE_INTEGER
	TimerResolution    ULONG
	MaximumFileSize    ULONG
	LogFileMode        LoggingMode
	BuffersWritten     ULONG
	DUMMYUNIONNAME2    [16]byte
	LoggerName         LPWSTR
	LogFileName        LPWSTR
	TimeZone           TIME_ZONE_INFORMATION
	BootTime           LARGE_INTEGER
	PerfFreq           LARGE_INTEGER
	StartTime          LARGE_INTEGER
	ReservedFlags      ETWClockType
	BuffersLost        ULONG
}

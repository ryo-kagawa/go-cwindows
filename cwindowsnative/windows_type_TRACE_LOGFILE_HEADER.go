// Code generated DO NOT EDIT

package cwindowsnative

type TRACE_LOGFILE_HEADER struct {
	BufferSize         ULONG
	DUMMYUNIONNAME     [4]byte
	ProviderVersion    ULONG
	NumberOfProcessors ULONG
	EndTime            LARGE_INTEGER
	TimerResolution    ULONG
	MaximumFileSize    ULONG
	LogFileMode        ULONG
	BuffersWritten     ULONG
	DUMMYUNIONNAME2    [16]byte
	LoggerName         LPWSTR
	LogFileName        LPWSTR
	TimeZone           TIME_ZONE_INFORMATION
	BootTime           LARGE_INTEGER
	PerfFreq           LARGE_INTEGER
	StartTime          LARGE_INTEGER
	ReservedFlags      ULONG
	BuffersLost        ULONG
}

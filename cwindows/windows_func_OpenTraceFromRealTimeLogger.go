package cwindows

import (
	"unsafe"

	opentracefromrealtimelogger "github.com/ryo-kagawa/go-cwindows/dll/sechost/OpenTraceFromRealTimeLogger"
)

func InitSechostOpenTraceFromRealTimeLogger() {
	opentracefromrealtimelogger.InitSechostOpenTraceFromRealTimeLogger()
}

func OpenTraceFromRealTimeLogger(LoggerName PCWSTR, Options *ETW_OPEN_TRACE_OPTIONS, LogFileHeader *TRACE_LOGFILE_HEADER) (PROCESSTRACE_HANDLE, error) {
	r1, lastErr := opentracefromrealtimelogger.OpenTraceFromRealTimeLogger(
		uintptr(unsafe.Pointer(LoggerName)),
		uintptr(unsafe.Pointer(Options)),
		uintptr(unsafe.Pointer(LogFileHeader)),
	)
	return PROCESSTRACE_HANDLE(r1), lastErr
}

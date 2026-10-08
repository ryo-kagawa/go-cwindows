package opentracefromrealtimelogger

import "github.com/ryo-kagawa/go-cwindows/dll/sechost"

var impl = func(LoggerName uintptr, Options uintptr, LogFileHeader uintptr) (uintptr, error) {
	panic("OpenTraceFromRealTimeLogger not initialized. Call InitSechostOpenTraceFromRealTimeLogger() first.")
}

func InitSechostOpenTraceFromRealTimeLogger() {
	proc := sechost.ModSechost.NewProc("OpenTraceFromRealTimeLogger")
	impl = func(LoggerName uintptr, Options uintptr, LogFileHeader uintptr) (uintptr, error) {
		r1, _, lastErr := proc.Call(
			LoggerName,
			Options,
			LogFileHeader,
		)
		return r1, lastErr
	}
}

func OpenTraceFromRealTimeLogger(LoggerName uintptr, Options uintptr, LogFileHeader uintptr) (uintptr, error) {
	r1, lastErr := impl(
		LoggerName,
		Options,
		LogFileHeader,
	)
	return r1, lastErr
}

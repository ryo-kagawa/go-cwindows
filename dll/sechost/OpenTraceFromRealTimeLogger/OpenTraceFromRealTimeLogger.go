// Code generated DO NOT EDIT

package opentracefromrealtimelogger

import sechost "github.com/ryo-kagawa/go-cwindows/dll/sechost"

var impl = func(LoggerName uintptr, Options uintptr, LogFileHeader uintptr) (uintptr, error) {
	panic("SechostOpenTraceFromRealTimeLogger not initialized. Call InitSechostOpenTraceFromRealTimeLogger() first.")
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

func SechostOpenTraceFromRealTimeLogger(LoggerName uintptr, Options uintptr, LogFileHeader uintptr) (uintptr, error) {
	r1, lastErr := impl(
		LoggerName,
		Options,
		LogFileHeader,
	)
	return r1, lastErr
}

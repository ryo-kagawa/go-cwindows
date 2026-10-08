package cwindows

type ETW_BUFFER_CALLBACK_INFORMATION struct {
	TraceHandle   PROCESSTRACE_HANDLE
	LogfileHeader *TRACE_LOGFILE_HEADER
	BuffersRead   ULONG
}

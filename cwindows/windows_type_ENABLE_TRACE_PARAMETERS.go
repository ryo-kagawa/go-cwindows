package cwindows

type ENABLE_TRACE_PARAMETERS struct {
	Version          EnableTraceParametersVersion
	EnableProperty   ULONG
	ControlFlags     ULONG
	SourceId         GUID
	EnableFilterDesc PEVENT_FILTER_DESCRIPTOR
	FilterDescCount  ULONG
}
type PENABLE_TRACE_PARAMETERS *ENABLE_TRACE_PARAMETERS

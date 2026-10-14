// Code generated DO NOT EDIT

package cwindows

type ENABLE_TRACE_PARAMETERS struct {
	Version          EnableTraceParametersVersion
	EnableProperty   ULONG
	ControlFlags     ULONG
	SourceId         GUID
	EnableFilterDesc PEVENT_FILTER_DESCRIPTOR
	FilterDescCount  ULONG
}

package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/cwindows_windows_struct_type.tmpl
var cwindowsWindowsStructTypeTmpl string

//go:embed templates/cwindowsnative_windows_struct_type.tmpl
var cwindowsnativeWindowsStructTypeTmpl string

type WindowsStructType struct {
	Name         string
	Import       string
	Struct       string
	CustomStruct string
	Functions    []string
}

var WindowsStructTypes = []WindowsStructType{
	{
		Name: "ENABLE_TRACE_PARAMETERS",
		Struct: `
	Version          ULONG
	EnableProperty   ULONG
	ControlFlags     ULONG
	SourceId         GUID
	EnableFilterDesc PEVENT_FILTER_DESCRIPTOR
	FilterDescCount  ULONG
`,
		CustomStruct: `
	Version          EnableTraceParametersVersion
	EnableProperty   ULONG
	ControlFlags     ULONG
	SourceId         GUID
	EnableFilterDesc PEVENT_FILTER_DESCRIPTOR
	FilterDescCount  ULONG
`,
	},
	{
		Name: "ETW_BUFFER_CALLBACK_INFORMATION",
		Struct: `
	TraceHandle   PROCESSTRACE_HANDLE
	LogfileHeader *TRACE_LOGFILE_HEADER
	BuffersRead   ULONG
`,
	},
	{
		Name: "ETW_BUFFER_CONTEXT",
		Struct: `
	DUMMYUNIONNAME [2]byte
	LoggerId       USHORT
`,
	},
	{
		Name: "ETW_BUFFER_HEADER",
		Struct: `
	Reserved1     [4]ULONG
	TimeStamp     LARGE_INTEGER
	Reserved2     [2]ULONG
	ClientContext ETW_BUFFER_CONTEXT
	Reserved3     ULONG
	FilledBytes   ULONG
	Reserved4     [5]ULONG
`,
	},
	{
		Name: "ETW_OPEN_TRACE_OPTIONS",
		Import: `
import (
	"unsafe"

	"golang.org/x/sys/windows"
)
`,
		Struct: `
	ProcessTraceModes     ETW_PROCESS_TRACE_MODES
	EventCallback         uintptr
	EventCallbackContext  unsafe.Pointer
	BufferCallback        uintptr
	BufferCallbackContext unsafe.Pointer
`,
		Functions: []string{
			`
func (e *ETW_OPEN_TRACE_OPTIONS) SetEventCallback(value PEVENT_RECORD_CALLBACK) {
	e.EventCallback = windows.NewCallback(func(EventRecord PEVENT_RECORD) uintptr {
		value(EventRecord)
		return 0
	})
}
`,
			`
func (e *ETW_OPEN_TRACE_OPTIONS) SetBufferCallback(value PETW_BUFFER_CALLBACK) {
	e.BufferCallback = windows.NewCallback(func(Buffer *ETW_BUFFER_HEADER, BufferSize ULONG, ConsumerInfo *ETW_BUFFER_CALLBACK_INFORMATION, CallbackContext unsafe.Pointer) BOOL {
		return value(Buffer, BufferSize, ConsumerInfo, CallbackContext)
	})
}
`,
		},
	},
	{
		Name: "EVENT_DESCRIPTOR",
		Struct: `
	Id      USHORT
	Version UCHAR
	Channel UCHAR
	Level   UCHAR
	Opcode  UCHAR
	Task    USHORT
	Keyword ULONGLONG
`,
		CustomStruct: `
	Id      ProcessEventId
	Version UCHAR
	Channel UCHAR
	Level   WinEventLevel
	Opcode  WinEventOpCode
	Task    USHORT
	Keyword ULONGLONG
`,
	},
	{
		Name: "EVENT_FILTER_DESCRIPTOR",
		Struct: `
	Ptr  ULONGLONG
	Size ULONG
	Type ULONG
`,
		CustomStruct: `
	Ptr  ULONGLONG
	Size ULONG
	Type EventFilterType
`,
	},
	{
		Name: "EVENT_HEADER",
		Struct: `
	Size            USHORT
	HeaderType      USHORT
	Flags           USHORT
	EventProperty   USHORT
	ThreadId        ULONG
	ProcessId       ULONG
	TimeStamp       LARGE_INTEGER
	ProviderId      GUID
	EventDescriptor EVENT_DESCRIPTOR
	DUMMYUNIONNAME  [8]byte
	ActivityId      GUID
`,
		CustomStruct: `
	Size            USHORT
	HeaderType      USHORT
	Flags           EventHeaderFlag
	EventProperty   EventHeaderProperty
	ThreadId        ULONG
	ProcessId       ULONG
	TimeStamp       LARGE_INTEGER
	ProviderId      GUID
	EventDescriptor EVENT_DESCRIPTOR
	DUMMYUNIONNAME  [8]byte
	ActivityId      GUID
`,
	},
	{
		Name: "EVENT_HEADER_EXTENDED_DATA_ITEM",
		Struct: `
	Reserved1 USHORT
	ExtType   USHORT
	_         [2]byte
	DataSize  USHORT
	DataPtr   ULONGLONG
`,
		CustomStruct: `
	Reserved1 USHORT
	ExtType   EventHeaderExtType
	_         [2]byte
	DataSize  USHORT
	DataPtr   ULONGLONG
`,
	},
	{
		Name: "EVENT_RECORD",
		Struct: `
	EventHeader       EVENT_HEADER
	BufferContext     ETW_BUFFER_CONTEXT
	ExtendedDataCount USHORT
	UserDataLength    USHORT
	ExtendedData      PEVENT_HEADER_EXTENDED_DATA_ITEM
	UserData          PVOID
	UserContext       PVOID
`,
	},
	{
		Name: "EVENT_TRACE_PROPERTIES",
		Struct: `
	Wnode               WNODE_HEADER
	BufferSize          ULONG
	MinimumBuffers      ULONG
	MaximumBuffers      ULONG
	MaximumFileSize     ULONG
	LogFileMode         ULONG
	FlushTimer          ULONG
	EnableFlags         ULONG
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
`,
		CustomStruct: `
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
`,
	},
	{
		Name: "EVENT_TRACE_PROPERTIES_V2",
		Struct: `
	EVENT_TRACE_PROPERTIES
	DUMMYUNIONNAME2 [4]byte
	FilterDescCount ULONG
	FilterDesc      PEVENT_FILTER_DESCRIPTOR
	DUMMYUNIONNAME3 [8]byte
`,
	},
	{
		Name: "FILETIME",
		Struct: `
	LowDateTime  DWORD
	HighDateTime DWORD
`,
	},
	{
		Name: "GUID",
		Struct: `
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
`,
	},
	{
		Name: "PROCESSENTRY32",
		Struct: `
	Size            DWORD
	Usage           DWORD
	ProcessID       DWORD
	DefaultHeapID   ULONG_PTR
	ModuleID        DWORD
	Threads         DWORD
	ParentProcessID DWORD
	PriClassBase    LONG
	Flags           DWORD
	ExeFile         [MAX_PATH]WCHAR
`,
	},
	{
		Name: "SID",
		Struct: `
	Revision            BYTE
	SubAuthorityCount   BYTE
	IdentifierAuthority SID_IDENTIFIER_AUTHORITY
	SubAuthority        []DWORD
`,
	},
	{
		Name: "SID_IDENTIFIER_AUTHORITY",
		Struct: `
	Value [6]BYTE
`,
		CustomStruct: `
	Value SIDIdentifierAuthority
`,
	},
	{
		Name: "SYSTEMTIME",
		Struct: `
	Year         WORD
	Month        WORD
	DayOfWeek    WORD
	Day          WORD
	Hour         WORD
	Minute       WORD
	Second       WORD
	Milliseconds WORD
`,
	},
	{
		Name: "TIME_ZONE_INFORMATION",
		Struct: `
	Bias         LONG
	StandardName [32]WCHAR
	StandardDate SYSTEMTIME
	StandardBias LONG
	DaylightName [32]WCHAR
	DaylightDate SYSTEMTIME
	DaylightBias LONG
`,
	},
	{
		Name: "TRACE_LOGFILE_HEADER",
		Struct: `
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
`,
		CustomStruct: `
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
`,
	},
	{
		Name: "WNODE_HEADER",
		Struct: `
	BufferSize    ULONG
	ProviderId    ULONG
	union1        [8]byte
	union2        [8]byte
	Guid          GUID
	ClientContext ULONG
	Flags         ULONG
`,
		CustomStruct: `
	BufferSize    ULONG
	ProviderId    ULONG
	union1        [8]byte
	union2        [8]byte
	Guid          GUID
	ClientContext ETWClockType
	Flags         WnodeFlag
`,
	},
}

type WindowsStructTypeTemplateParameter struct {
	Name      string
	Import    string
	Struct    string
	Functions []string
}

func generateWindowsStructType() {
	cwindowsTmpl := template.Must(template.New("cwindows_windows_struct_type.tmpl").Parse(cwindowsWindowsStructTypeTmpl))
	cwindowsnativeTmpl := template.Must(template.New("cwindowsnative_windows_struct_type.tmpl").Parse(cwindowsnativeWindowsStructTypeTmpl))

	for _, value := range WindowsStructTypes {
		cwindowsFile, err := os.Create(filepath.Join("cwindows", fmt.Sprintf("windows_type_%s.go", value.Name)))
		if err != nil {
			panic(err)
		}
		defer cwindowsFile.Close()
		cwindowsStruct := value.CustomStruct
		if cwindowsStruct == "" {
			cwindowsStruct = value.Struct
		}
		if err := cwindowsTmpl.Execute(cwindowsFile, WindowsStructTypeTemplateParameter{
			Name:      value.Name,
			Import:    value.Import,
			Struct:    cwindowsStruct,
			Functions: value.Functions,
		}); err != nil {
			panic(err)
		}
		cwindowsnativeFile, err := os.Create(filepath.Join("cwindowsnative", fmt.Sprintf("windows_type_%s.go", value.Name)))
		if err != nil {
			panic(err)
		}
		defer cwindowsnativeFile.Close()
		if err := cwindowsnativeTmpl.Execute(cwindowsnativeFile, value); err != nil {
			panic(err)
		}
	}
}

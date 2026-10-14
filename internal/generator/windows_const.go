package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/cwindows_windows_const.tmpl
var cwindowsWindowsConstTmpl string

//go:embed templates/cwindowsnative_windows_const.tmpl
var cwindowsnativeWindowsConstTmpl string

type WindowsConst struct {
	Name       string
	IsVar      bool
	Type       string
	CustomType string
	Value      string
}

var WindowsConsts = []WindowsConst{
	{
		Name:  "ANYSIZE_ARRAY",
		Value: "1",
	},
	{
		Name:       "ENABLE_TRACE_PARAMETERS_VERSION",
		CustomType: "EnableTraceParametersVersion",
		Value:      "1",
	},
	{
		Name:       "ENABLE_TRACE_PARAMETERS_VERSION_2",
		CustomType: "EnableTraceParametersVersion",
		Value:      "2",
	},
	{
		Name:       "ERROR_NO_MORE_FILES",
		CustomType: "Error",
		Value:      "18",
	},
	{
		Name:       "ERROR_SUCCESS",
		CustomType: "Error",
		Value:      "0",
	},
	{
		Name:  "ETW_PROCESS_TRACE_MODE_NONE",
		Type:  "ETW_PROCESS_TRACE_MODES",
		Value: "0",
	},
	{
		Name:  "ETW_PROCESS_TRACE_MODE_RAW_TIMESTAMP",
		Type:  "ETW_PROCESS_TRACE_MODES",
		Value: "0x00000001",
	},
	{
		Name:       "EVENT_CONTROL_CODE_ENABLE_PROVIDER",
		CustomType: "EventControlCode",
		Value:      "1",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_CONTAINER_ID",
		CustomType: "EventHeaderExtType",
		Value:      "0x0010",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_CONTROL_GUID",
		CustomType: "EventHeaderExtType",
		Value:      "0x000E",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_EVENT_KEY",
		CustomType: "EventHeaderExtType",
		Value:      "0x000A",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_EVENT_SCHEMA_TL",
		CustomType: "EventHeaderExtType",
		Value:      "0x000B",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_INSTANCE_INFO",
		CustomType: "EventHeaderExtType",
		Value:      "0x0004",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_MAX",
		CustomType: "EventHeaderExtType",
		Value:      "0x0013",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_PEBS_INDEX",
		CustomType: "EventHeaderExtType",
		Value:      "0x0007",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_PMC_COUNTERS",
		CustomType: "EventHeaderExtType",
		Value:      "0x0008",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_PROCESS_START_KEY",
		CustomType: "EventHeaderExtType",
		Value:      "0x000D",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_PROV_TRAITS",
		CustomType: "EventHeaderExtType",
		Value:      "0x000C",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_PSM_KEY",
		CustomType: "EventHeaderExtType",
		Value:      "0x0009",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_QPC_DELTA",
		CustomType: "EventHeaderExtType",
		Value:      "0x000F",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_RELATED_ACTIVITYID",
		CustomType: "EventHeaderExtType",
		Value:      "0x0001",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_SID",
		CustomType: "EventHeaderExtType",
		Value:      "0x0002",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_STACK_KEY32",
		CustomType: "EventHeaderExtType",
		Value:      "0x0011",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_STACK_KEY64",
		CustomType: "EventHeaderExtType",
		Value:      "0x0012",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_STACK_TRACE32",
		CustomType: "EventHeaderExtType",
		Value:      "0x0005",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_STACK_TRACE64",
		CustomType: "EventHeaderExtType",
		Value:      "0x0006",
	},
	{
		Name:       "EVENT_HEADER_EXT_TYPE_TS_ID",
		CustomType: "EventHeaderExtType",
		Value:      "0x0003",
	},
	{
		Name:       "EVENT_HEADER_FLAG_32_BIT_HEADER",
		CustomType: "EventHeaderFlag",
		Value:      "0x0020",
	},
	{
		Name:       "EVENT_HEADER_FLAG_64_BIT_HEADER",
		CustomType: "EventHeaderFlag",
		Value:      "0x0040",
	},
	{
		Name:       "EVENT_HEADER_FLAG_CLASSIC_HEADER",
		CustomType: "EventHeaderFlag",
		Value:      "0x0100",
	},
	{
		Name:       "EVENT_HEADER_FLAG_DECODE_GUID",
		CustomType: "EventHeaderFlag",
		Value:      "0x0080",
	},
	{
		Name:       "EVENT_HEADER_FLAG_EXTENDED_INFO",
		CustomType: "EventHeaderFlag",
		Value:      "0x0001",
	},
	{
		Name:       "EVENT_HEADER_FLAG_NO_CPUTIME",
		CustomType: "EventHeaderFlag",
		Value:      "0x0010",
	},
	{
		Name:       "EVENT_HEADER_FLAG_PRIVATE_SESSION",
		CustomType: "EventHeaderFlag",
		Value:      "0x0002",
	},
	{
		Name:       "EVENT_HEADER_FLAG_PROCESSOR_INDEX",
		CustomType: "EventHeaderFlag",
		Value:      "0x0200",
	},
	{
		Name:       "EVENT_HEADER_FLAG_RESERVED1",
		CustomType: "EventHeaderFlag",
		Value:      "0x0400",
	},
	{
		Name:       "EVENT_HEADER_FLAG_STRING_ONLY",
		CustomType: "EventHeaderFlag",
		Value:      "0x0004",
	},
	{
		Name:       "EVENT_HEADER_FLAG_TRACE_MESSAGE",
		CustomType: "EventHeaderFlag",
		Value:      "0x0008",
	},
	{
		Name:       "EVENT_HEADER_PROPERTY_FORWARDED_XML",
		CustomType: "EventHeaderProperty",
		Value:      "0x0002",
	},
	{
		Name:       "EVENT_HEADER_PROPERTY_LEGACY_EVENTLOG",
		CustomType: "EventHeaderProperty",
		Value:      "0x0004",
	},
	{
		Name:       "EVENT_HEADER_PROPERTY_RELOGGABLE",
		CustomType: "EventHeaderProperty",
		Value:      "0x0008",
	},
	{
		Name:       "EVENT_HEADER_PROPERTY_XML",
		CustomType: "EventHeaderProperty",
		Value:      "0x0001",
	},
	{
		Name:       "EVENT_TRACE_ADD_HEADER_MODE",
		CustomType: "LoggingMode",
		Value:      "0x00001000",
	},
	{
		Name:       "EVENT_TRACE_ADDTO_TRIAGE_DUMP",
		CustomType: "LoggingMode",
		Value:      "0x80000000",
	},
	{
		Name:       "EVENT_TRACE_BUFFERING_MODE",
		CustomType: "LoggingMode",
		Value:      "0x00000400",
	},
	{
		Name:       "EVENT_TRACE_CONTROL_CONVERT_TO_REALTIME",
		CustomType: "ControlTraceCode",
		Value:      "5",
	},
	{
		Name:       "EVENT_TRACE_CONTROL_CONVERT_TO_REALTIME",
		CustomType: "ControlTraceCode",
		Value:      "5",
	},
	{
		Name:       "EVENT_TRACE_CONTROL_FLUSH",
		CustomType: "ControlTraceCode",
		Value:      "3",
	},
	{
		Name:       "EVENT_TRACE_CONTROL_INCREMENT_FILE",
		CustomType: "ControlTraceCode",
		Value:      "4",
	},
	{
		Name:       "EVENT_TRACE_CONTROL_QUERY",
		CustomType: "ControlTraceCode",
		Value:      "0",
	},
	{
		Name:       "EVENT_TRACE_CONTROL_STOP",
		CustomType: "ControlTraceCode",
		Value:      "1",
	},
	{
		Name:       "EVENT_TRACE_CONTROL_UPDATE",
		CustomType: "ControlTraceCode",
		Value:      "2",
	},
	{
		Name:       "EVENT_TRACE_DELAY_OPEN_FILE_MODE",
		CustomType: "LoggingMode",
		Value:      "0x00000200",
	},
	{
		Name:       "EVENT_TRACE_FILE_MODE_APPEND",
		CustomType: "LoggingMode",
		Value:      "0x00000004",
	},
	{
		Name:       "EVENT_TRACE_FILE_MODE_CIRCULAR",
		CustomType: "LoggingMode",
		Value:      "0x00000002",
	},
	{
		Name:       "EVENT_TRACE_FILE_MODE_NEWFILE",
		CustomType: "LoggingMode",
		Value:      "0x00000008",
	},
	{
		Name:       "EVENT_TRACE_FILE_MODE_NONE",
		CustomType: "LoggingMode",
		Value:      "0x00000000",
	},
	{
		Name:       "EVENT_TRACE_FILE_MODE_PREALLOCATE",
		CustomType: "LoggingMode",
		Value:      "0x00000020",
	},
	{
		Name:       "EVENT_TRACE_FILE_MODE_SEQUENTIAL",
		CustomType: "LoggingMode",
		Value:      "0x00000001",
	},
	{
		Name:       "EVENT_TRACE_INDEPENDENT_SESSION_MODE",
		CustomType: "LoggingMode",
		Value:      "0x08000000",
	},
	{
		Name:       "EVENT_TRACE_MODE_RESERVED",
		CustomType: "LoggingMode",
		Value:      "0x00100000",
	},
	{
		Name:       "EVENT_TRACE_NO_PER_PROCESSOR_BUFFERING",
		CustomType: "LoggingMode",
		Value:      "0x10000000",
	},
	{
		Name:       "EVENT_TRACE_NONSTOPPABLE_MODE",
		CustomType: "LoggingMode",
		Value:      "0x00000040",
	},
	{
		Name:       "EVENT_TRACE_PERSIST_ON_HYBRID_SHUTDOWN",
		CustomType: "LoggingMode",
		Value:      "0x00800000",
	},
	{
		Name:       "EVENT_TRACE_PRIVATE_IN_PROC",
		CustomType: "LoggingMode",
		Value:      "0x00020000",
	},
	{
		Name:       "EVENT_TRACE_PRIVATE_LOGGER_MODE",
		CustomType: "LoggingMode",
		Value:      "0x00000800",
	},
	{
		Name:       "EVENT_TRACE_REAL_TIME_MODE",
		CustomType: "LoggingMode",
		Value:      "0x00000100",
	},
	{
		Name:       "EVENT_TRACE_RELOG_MODE",
		CustomType: "LoggingMode",
		Value:      "0x00010000",
	},
	{
		Name:       "EVENT_TRACE_SECURE_MODE",
		CustomType: "LoggingMode",
		Value:      "0x00000080",
	},
	{
		Name:       "EVENT_TRACE_STOP_ON_HYBRID_SHUTDOWN",
		CustomType: "LoggingMode",
		Value:      "0x00400000",
	},
	{
		Name:       "EVENT_TRACE_SYSTEM_LOGGER_MODE",
		CustomType: "LoggingMode",
		Value:      "0x02000000",
	},
	{
		Name:       "EVENT_TRACE_USE_GLOBAL_SEQUENCE",
		CustomType: "LoggingMode",
		Value:      "0x00004000",
	},
	{
		Name:       "EVENT_TRACE_USE_KBYTES_FOR_SIZE",
		CustomType: "LoggingMode",
		Value:      "0x00002000",
	},
	{
		Name:       "EVENT_TRACE_USE_LOCAL_SEQUENCE",
		CustomType: "LoggingMode",
		Value:      "0x00008000",
	},
	{
		Name:       "EVENT_TRACE_USE_PAGED_MEMORY",
		CustomType: "LoggingMode",
		Value:      "0x01000000",
	},
	{
		Name:       "FALSE",
		CustomType: "BOOL",
		Value:      "0",
	},
	{
		Name:  "INVALID_HANDLE_VALUE",
		Value: "^HANDLE(0)",
	},
	{
		Name:  "INVALID_PROCESSTRACE_HANDLE",
		Value: "PROCESSTRACE_HANDLE(INVALID_HANDLE_VALUE)",
	},
	{
		Name:  "MAX_PATH",
		Value: "260",
	},
	{
		Name:       "PROCESS_ALL_ACCESS",
		CustomType: "ProcessAccess",
		Value:      "STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xFFFF",
	},
	{
		Name:       "PROCESS_CREATE_PROCESS",
		CustomType: "ProcessAccess",
		Value:      "0x0080",
	},
	{
		Name:       "PROCESS_CREATE_THREAD",
		CustomType: "ProcessAccess",
		Value:      "0x0002",
	},
	{
		Name:       "PROCESS_DUP_HANDLE",
		CustomType: "ProcessAccess",
		Value:      "0x0040",
	},
	{
		Name:       "PROCESS_QUERY_INFORMATION",
		CustomType: "ProcessAccess",
		Value:      "0x0400",
	},
	{
		Name:       "PROCESS_QUERY_LIMITED_INFORMATION",
		CustomType: "ProcessAccess",
		Value:      "0x1000",
	},
	{
		Name:       "PROCESS_SET_INFORMATION",
		CustomType: "ProcessAccess",
		Value:      "0x0200",
	},
	{
		Name:       "PROCESS_SET_LIMITED_INFORMATION",
		CustomType: "ProcessAccess",
		Value:      "0x2000",
	},
	{
		Name:       "PROCESS_SET_QUOTA",
		CustomType: "ProcessAccess",
		Value:      "0x0100",
	},
	{
		Name:       "PROCESS_SET_SESSIONID",
		CustomType: "ProcessAccess",
		Value:      "0x0004",
	},
	{
		Name:       "PROCESS_SUSPEND_RESUME",
		CustomType: "ProcessAccess",
		Value:      "0x0800",
	},
	{
		Name:       "PROCESS_TERMINATE",
		CustomType: "ProcessAccess",
		Value:      "0x0001",
	},
	{
		Name:       "PROCESS_VM_OPERATION",
		CustomType: "ProcessAccess",
		Value:      "0x0008",
	},
	{
		Name:       "PROCESS_VM_READ",
		CustomType: "ProcessAccess",
		Value:      "0x0010",
	},
	{
		Name:       "PROCESS_VM_WRITE",
		CustomType: "ProcessAccess",
		Value:      "0x0020",
	},
	{
		Name:       "SECURITY_CREATOR_SID_AUTHORITY",
		IsVar:      true,
		CustomType: "SIDIdentifierAuthority",
		Value:      "[6]BYTE{0, 0, 0, 0, 0, 3}",
	},
	{
		Name:       "SECURITY_LOCAL_SID_AUTHORITY",
		IsVar:      true,
		CustomType: "SIDIdentifierAuthority",
		Value:      "[6]BYTE{0, 0, 0, 0, 0, 2}",
	},
	{
		Name:       "SECURITY_NON_UNIQUE_AUTHORITY",
		IsVar:      true,
		CustomType: "SIDIdentifierAuthority",
		Value:      "[6]BYTE{0, 0, 0, 0, 0, 4}",
	},
	{
		Name:       "SECURITY_NULL_SID_AUTHORITY",
		IsVar:      true,
		CustomType: "SIDIdentifierAuthority",
		Value:      "[6]BYTE{0, 0, 0, 0, 0, 0}",
	},
	{
		Name:       "SECURITY_RESOURCE_MANAGER_AUTHORITY",
		IsVar:      true,
		CustomType: "SIDIdentifierAuthority",
		Value:      "[6]BYTE{0, 0, 0, 0, 0, 9}",
	},
	{
		Name:       "SECURITY_WORLD_SID_AUTHORITY",
		IsVar:      true,
		CustomType: "SIDIdentifierAuthority",
		Value:      "[6]BYTE{0, 0, 0, 0, 0, 1}",
	},
	{
		Name:  "STANDARD_RIGHTS_REQUIRED",
		Value: "0x000F0000",
	},
	{
		Name:  "SYNCHRONIZE",
		Value: "0x00100000",
	},
	{
		Name:       "TH32CS_INHERIT",
		CustomType: "TH32CS",
		Value:      "0x80000000",
	},
	{
		Name:       "TH32CS_SNAPALL",
		CustomType: "TH32CS",
		Value:      "TH32CS_SNAPHEAPLIST | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD | TH32CS_SNAPMODULE",
	},
	{
		Name:       "TH32CS_SNAPHEAPLIST",
		CustomType: "TH32CS",
		Value:      "0x00000001",
	},
	{
		Name:       "TH32CS_SNAPMODULE",
		CustomType: "TH32CS",
		Value:      "0x00000008",
	},
	{
		Name:       "TH32CS_SNAPMODULE32",
		CustomType: "TH32CS",
		Value:      "0x00000010",
	},
	{
		Name:       "TH32CS_SNAPPROCESS",
		CustomType: "TH32CS",
		Value:      "0x00000002",
	},
	{
		Name:       "TH32CS_SNAPTHREAD",
		CustomType: "TH32CS",
		Value:      "0x00000004",
	},
	{
		Name:       "TRACE_LEVEL_CRITICAL",
		CustomType: "TraceLevel",
		Value:      "1",
	},
	{
		Name:       "TRACE_LEVEL_ERROR",
		CustomType: "TraceLevel",
		Value:      "2",
	},
	{
		Name:       "TRACE_LEVEL_INFORMATION",
		CustomType: "TraceLevel",
		Value:      "4",
	},
	{
		Name:       "TRACE_LEVEL_VERBOSE",
		CustomType: "TraceLevel",
		Value:      "5",
	},
	{
		Name:       "TRACE_LEVEL_WARNING",
		CustomType: "TraceLevel",
		Value:      "3",
	},
	{
		Name:       "TRUE",
		CustomType: "BOOL",
		Value:      "1",
	},
	{
		Name:       "WINEVENT_LEVEL_CRITICAL",
		CustomType: "WinEventLevel",
		Value:      "0x1",
	},
	{
		Name:       "WINEVENT_LEVEL_ERROR",
		CustomType: "WinEventLevel",
		Value:      "0x3",
	},
	{
		Name:       "WINEVENT_LEVEL_INFO",
		CustomType: "WinEventLevel",
		Value:      "0x4",
	},
	{
		Name:       "WINEVENT_LEVEL_LOG_ALWAYS",
		CustomType: "WinEventLevel",
		Value:      "0x0",
	},
	{
		Name:       "WINEVENT_LEVEL_VERBOSE",
		CustomType: "WinEventLevel",
		Value:      "0x5",
	},
	{
		Name:       "WINEVENT_LEVEL_WARNING",
		CustomType: "WinEventLevel",
		Value:      "0x3",
	},
	{
		Name:       "WINEVENT_OPCODE_DC_START",
		CustomType: "WinEventOpCode",
		Value:      "0x3",
	},
	{
		Name:       "WINEVENT_OPCODE_DC_STOP",
		CustomType: "WinEventOpCode",
		Value:      "0x4",
	},
	{
		Name:       "WINEVENT_OPCODE_EXTENSION",
		CustomType: "WinEventOpCode",
		Value:      "0x5",
	},
	{
		Name:       "WINEVENT_OPCODE_INFO",
		CustomType: "WinEventOpCode",
		Value:      "0x0",
	},
	{
		Name:       "WINEVENT_OPCODE_REPLY",
		CustomType: "WinEventOpCode",
		Value:      "0x6",
	},
	{
		Name:       "WINEVENT_OPCODE_RESUME",
		CustomType: "WinEventOpCode",
		Value:      "0x7",
	},
	{
		Name:       "WINEVENT_OPCODE_SEND",
		CustomType: "WinEventOpCode",
		Value:      "0x9",
	},
	{
		Name:       "WINEVENT_OPCODE_START",
		CustomType: "WinEventOpCode",
		Value:      "0x1",
	},
	{
		Name:       "WINEVENT_OPCODE_STOP",
		CustomType: "WinEventOpCode",
		Value:      "0x2",
	},
	{
		Name:       "WINEVENT_OPCODE_SUSPEND",
		CustomType: "WinEventOpCode",
		Value:      "0x8",
	},
	{
		Name:       "WNODE_FLAG_ALL_DATA",
		CustomType: "WnodeFlag",
		Value:      "0x00000001",
	},
	{
		Name:       "WNODE_FLAG_ANSI_INSTANCENAMES",
		CustomType: "WnodeFlag",
		Value:      "0x00004000",
	},
	{
		Name:       "WNODE_FLAG_EVENT_ITEM",
		CustomType: "WnodeFlag",
		Value:      "0x00000008",
	},
	{
		Name:       "WNODE_FLAG_EVENT_REFERENCE",
		CustomType: "WnodeFlag",
		Value:      "0x00002000",
	},
	{
		Name:       "WNODE_FLAG_FIXED_INSTANCE_SIZE",
		CustomType: "WnodeFlag",
		Value:      "0x00000010",
	},
	{
		Name:       "WNODE_FLAG_INSTANCES_SAME",
		CustomType: "WnodeFlag",
		Value:      "0x00000040",
	},
	{
		Name:       "WNODE_FLAG_INTERNAL",
		CustomType: "WnodeFlag",
		Value:      "0x00000100",
	},
	{
		Name:       "WNODE_FLAG_LOG_WNODE",
		CustomType: "WnodeFlag",
		Value:      "0x00040000",
	},
	{
		Name:       "WNODE_FLAG_METHOD_ITEM",
		CustomType: "WnodeFlag",
		Value:      "0x00008000",
	},
	{
		Name:       "WNODE_FLAG_NO_HEADER",
		CustomType: "WnodeFlag",
		Value:      "0x00200000",
	},
	{
		Name:       "WNODE_FLAG_PDO_INSTANCE_NAMES",
		CustomType: "WnodeFlag",
		Value:      "0x00010000",
	},
	{
		Name:       "WNODE_FLAG_PERSIST_EVENT",
		CustomType: "WnodeFlag",
		Value:      "0x00000400",
	},
	{
		Name:       "WNODE_FLAG_SEND_DATA_BLOCK",
		CustomType: "WnodeFlag",
		Value:      "0x00400000",
	},
	{
		Name:       "WNODE_FLAG_SEVERITY_MASK",
		CustomType: "WnodeFlag",
		Value:      "0xff000000",
	},
	{
		Name:       "WNODE_FLAG_SINGLE_INSTANCE",
		CustomType: "WnodeFlag",
		Value:      "0x00000002",
	},
	{
		Name:       "WNODE_FLAG_SINGLE_ITEM",
		CustomType: "WnodeFlag",
		Value:      "0x00000004",
	},
	{
		Name:       "WNODE_FLAG_STATIC_INSTANCE_NAMES",
		CustomType: "WnodeFlag",
		Value:      "0x00000080",
	},
	{
		Name:       "WNODE_FLAG_TRACED_GUID",
		CustomType: "WnodeFlag",
		Value:      "0x00020000",
	},
	{
		Name:       "WNODE_FLAG_TOO_SMALL",
		CustomType: "WnodeFlag",
		Value:      "0x00000020",
	},
	{
		Name:       "WNODE_FLAG_USE_GUID_PTR",
		CustomType: "WnodeFlag",
		Value:      "0x00080000",
	},
	{
		Name:       "WNODE_FLAG_USE_MOF_PTR",
		CustomType: "WnodeFlag",
		Value:      "0x00100000",
	},
	{
		Name:       "WNODE_FLAG_USE_TIMESTAMP",
		CustomType: "WnodeFlag",
		Value:      "0x00000200",
	},
	{
		Name:       "WNODE_FLAG_VERSIONED_PROPERTIES",
		CustomType: "WnodeFlag",
		Value:      "0x00800000",
	},
}

func generateWindowsConst() {
	cwindowsTmpl := template.Must(template.New("cwindows_windows_const.tmpl").Parse(cwindowsWindowsConstTmpl))
	cwindowsnativeTmpl := template.Must(template.New("cwindowsnative_windows_const.tmpl").Parse(cwindowsnativeWindowsConstTmpl))

	for _, value := range WindowsConsts {
		cwindowsFile, err := os.Create(filepath.Join("cwindows", fmt.Sprintf("windows_const_%s.go", value.Name)))
		if err != nil {
			panic(err)
		}
		defer cwindowsFile.Close()
		if err := cwindowsTmpl.Execute(cwindowsFile, value); err != nil {
			panic(err)
		}
		cwindowsnativeFile, err := os.Create(filepath.Join("cwindowsnative", fmt.Sprintf("windows_const_%s.go", value.Name)))
		if err != nil {
			panic(err)
		}
		defer cwindowsnativeFile.Close()
		if err := cwindowsnativeTmpl.Execute(cwindowsnativeFile, value); err != nil {
			panic(err)
		}
	}
}

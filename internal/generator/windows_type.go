package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/cwindows_windows_type.tmpl
var cwindowsWindowsTypeTmpl string

//go:embed templates/cwindowsnative_windows_type.tmpl
var cwindowsnativeWindowsTypeTmpl string

type WindowsType struct {
	Name     string
	Import   string
	BaseType string
}

var WindowsTypes = []WindowsType{
	{
		Name:     "BOOL",
		BaseType: "int32",
	},
	{
		Name:     "BYTE",
		BaseType: "byte",
	},
	{
		Name:     "CONTROLTRACE_ID",
		BaseType: "ULONG64",
	},
	{
		Name:     "DWORD",
		BaseType: "uint32",
	},
	{
		Name:     "DWORD_PTR",
		BaseType: "ULONG_PTR",
	},
	{
		Name:     "ETW_PROCESS_TRACE_MODES",
		BaseType: "int32",
	},
	{
		Name: "HANDLE",
		// 本来はPVOIDが適切
		BaseType: "uintptr",
	},
	{
		Name:     "LARGE_INTEGER",
		BaseType: "[8]byte",
	},
	{
		Name:     "LONG",
		BaseType: "int32",
	},
	{
		Name:     "LPCGUID",
		BaseType: "*GUID",
	},
	{
		Name:     "LPCWSTR",
		BaseType: "*WCHAR",
	},
	{
		Name:     "LPFILETIME",
		BaseType: "*FILETIME",
	},
	{
		Name:     "LPPROCESSENTRY32",
		BaseType: "*PROCESSENTRY32",
	},
	{
		Name:     "LPWSTR",
		BaseType: "*WCHAR",
	},
	{
		Name:     "PCWSTR",
		BaseType: "*WCHAR",
	},
	{
		Name:     "PENABLE_TRACE_PARAMETERS",
		BaseType: "*ENABLE_TRACE_PARAMETERS",
	},
	{
		Name:     "PETW_BUFFER_CALLBACK",
		Import:   `import "unsafe"`,
		BaseType: "func(Buffer *ETW_BUFFER_HEADER, BufferSize ULONG, ConsumerInfo *ETW_BUFFER_CALLBACK_INFORMATION, CallbackContext unsafe.Pointer) BOOL",
	},
	{
		Name:     "PEVENT_FILTER_DESCRIPTOR",
		BaseType: "*EVENT_FILTER_DESCRIPTOR",
	},
	{
		Name:     "PEVENT_HEADER_EXTENDED_DATA_ITEM",
		BaseType: "*EVENT_HEADER_EXTENDED_DATA_ITEM",
	},
	{
		Name:     "PEVENT_RECORD",
		BaseType: "*EVENT_RECORD",
	},
	{
		Name:     "PEVENT_RECORD_CALLBACK",
		BaseType: "func(EventRecord PEVENT_RECORD)",
	},
	{
		Name:     "PEVENT_TRACE_PROPERTIES",
		BaseType: "*EVENT_TRACE_PROPERTIES",
	},
	{
		Name:     "PEVENT_TRACE_PROPERTIES_V2",
		BaseType: "*EVENT_TRACE_PROPERTIES_V2",
	},
	{
		Name:     "PROCESSTRACE_HANDLE",
		BaseType: "ULONG64",
	},
	{
		Name:     "PVOID",
		Import:   `import "unsafe"`,
		BaseType: "unsafe.Pointer",
	},
	{
		Name:     "UCHAR",
		BaseType: "uint8",
	},
	{
		Name:     "ULONG",
		BaseType: "uint32",
	},
	{
		Name:     "ULONG64",
		BaseType: "uint64",
	},
	{
		Name:     "ULONG_PTR",
		BaseType: "uintptr",
	},
	{
		Name:     "ULONGLONG",
		BaseType: "uint64",
	},
	{
		Name:     "USHORT",
		BaseType: "uint16",
	},
	{
		Name:     "WCHAR",
		BaseType: "uint16",
	},
	{
		Name:     "WORD",
		BaseType: "uint16",
	},
}

func generateWindowsType() {
	cwindowsTmpl := template.Must(template.New("cwindows_windows_type.tmpl").Parse(cwindowsWindowsTypeTmpl))
	cwindowsnativeTmpl := template.Must(template.New("cwindowsnative_windows_type.tmpl").Parse(cwindowsnativeWindowsTypeTmpl))

	for _, value := range WindowsTypes {
		cwindowsFile, err := os.Create(filepath.Join("cwindows", fmt.Sprintf("windows_type_%s.go", value.Name)))
		if err != nil {
			panic(err)
		}
		defer cwindowsFile.Close()
		if err := cwindowsTmpl.Execute(cwindowsFile, value); err != nil {
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

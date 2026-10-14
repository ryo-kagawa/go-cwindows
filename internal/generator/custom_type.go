package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/cwindows_custom_type.tmpl
var cwindowsCustomTypeTmpl string

type CustomType struct {
	Name     string
	BaseType string
}

var CustomTypes = []CustomType{
	{
		Name:     "ControlTraceCode",
		BaseType: "int32",
	},
	{
		Name:     "EnableTraceParametersVersion",
		BaseType: "ULONG",
	},
	{
		Name:     "Error",
		BaseType: "uint32",
	},
	{
		Name:     "ETWClockType",
		BaseType: "ULONG",
	},
	{
		Name:     "EventControlCode",
		BaseType: "ULONG",
	},
	{
		Name:     "EventFilterType",
		BaseType: "ULONG",
	},
	{
		Name:     "EventHeaderExtType",
		BaseType: "USHORT",
	},
	{
		Name:     "EventHeaderFlag",
		BaseType: "USHORT",
	},
	{
		Name:     "EventHeaderProperty",
		BaseType: "USHORT",
	},
	{
		Name:     "EventTraceEnableFlag",
		BaseType: "ULONG",
	},
	{
		Name:     "LoggingMode",
		BaseType: "ULONG",
	},
	{
		Name:     "ProcessAccess",
		BaseType: "DWORD",
	},
	{
		Name:     "ProcessEventId",
		BaseType: "USHORT",
	},
	{
		Name:     "SIDIdentifierAuthority",
		BaseType: "[6]BYTE",
	},
	{
		Name:     "TH32CS",
		BaseType: "DWORD",
	},
	{
		Name:     "TraceLevel",
		BaseType: "UCHAR",
	},
	{
		Name:     "WinEventLevel",
		BaseType: "UCHAR",
	},
	{
		Name:     "WinEventOpCode",
		BaseType: "UCHAR",
	},
	{
		Name:     "WnodeFlag",
		BaseType: "ULONG",
	},
}

func generateCustomType() {
	cwindowsTmpl := template.Must(template.New("cwindows_custom_type.tmpl").Parse(cwindowsCustomTypeTmpl))

	for _, value := range CustomTypes {
		cwindowsFile, err := os.Create(filepath.Join("cwindows", fmt.Sprintf("custom_type_%s.go", value.Name)))
		if err != nil {
			panic(err)
		}
		defer cwindowsFile.Close()
		if err := cwindowsTmpl.Execute(cwindowsFile, value); err != nil {
			panic(err)
		}
	}
}

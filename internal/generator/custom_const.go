package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/cwindows_custom_const.tmpl
var cwindowsCustomConstTmpl string

type CustomConst struct {
	Name  string
	IsVar bool
	Type  string
	Value string
}

var CustomConsts = []CustomConst{
	{
		Name:  "CPUCycle",
		Type:  "ETWClockType",
		Value: "3",
	},
	{
		Name:  "MicrosoftWindowsKernelProcessEventIdProcessStart",
		Type:  "ProcessEventId",
		Value: "1",
	},
	{
		Name:  "MicrosoftWindowsKernelProcess",
		IsVar: true,
		Value: `GUID{
	Data1: 0x22fb2cd6,
	Data2: 0x0e7b,
	Data3: 0x422b,
	Data4: [8]byte{0xa0, 0xc7, 0x2f, 0xad, 0x1f, 0xd0, 0xe7, 0x16},
}`,
	},
	{
		Name:  "MicrosoftWindowsKernelProcessWINEVENT_KEYWORD_PROCESS",
		Type:  "ULONGLONG",
		Value: "0x8000000000000010",
	},
	{
		Name:  "QueryPerformanceCounter",
		Type:  "ETWClockType",
		Value: "1",
	},
	{
		Name:  "SystemTimer",
		Type:  "ETWClockType",
		Value: "2",
	},
}

func generateCustomConst() {
	cwindowsTmpl := template.Must(template.New("cwindows_custom_const.tmpl").Parse(cwindowsCustomConstTmpl))

	for _, value := range CustomConsts {
		cwindowsFile, err := os.Create(filepath.Join("cwindows", fmt.Sprintf("custom_const_%s.go", value.Name)))
		if err != nil {
			panic(err)
		}
		defer cwindowsFile.Close()
		if err := cwindowsTmpl.Execute(cwindowsFile, value); err != nil {
			panic(err)
		}
	}
}

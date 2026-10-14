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

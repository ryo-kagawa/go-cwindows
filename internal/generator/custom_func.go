package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/cwindows_custom_func.tmpl
var cwindowsCustomFuncTmpl string

type CustomFunc struct {
	Name     string
	Import   string
	Function string
}

var CustomFuncs = []CustomFunc{
	{
		Name: "UTF16ToWCHAR",
		Function: `
func UTF16ToWCHAR(value []uint16) []WCHAR {
	result := make([]WCHAR, len(value))
	for i, c := range value {
		result[i] = WCHAR(c)
	}
	return result
}
`,
	},
	{
		Name: "WCHARToString",
		Import: `
import (
	"unsafe"

	"golang.org/x/sys/windows"
)
`,
		Function: `
func WCHARToString(value []WCHAR) string {
	if len(value) == 0 {
		return ""
	}

	return windows.UTF16ToString(
		unsafe.Slice(
			(*uint16)(unsafe.Pointer(&value[0])),
			len(value),
		),
	)
}
`,
	},
}

func generateCustomFunc() {
	cwindowsTmpl := template.Must(template.New("cwindows_custom_func.tmpl").Parse(cwindowsCustomFuncTmpl))

	for _, value := range CustomFuncs {
		cwindowsFile, err := os.Create(filepath.Join("cwindows", fmt.Sprintf("custom_func_%s.go", value.Name)))
		if err != nil {
			panic(err)
		}
		defer cwindowsFile.Close()
		if err := cwindowsTmpl.Execute(cwindowsFile, value); err != nil {
			panic(err)
		}
	}
}

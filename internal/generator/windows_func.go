package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/dll_func.tmpl
var dllFuncTmpl string

//go:embed templates/cwindows_windows_func.tmpl
var cwindowsWindowsFuncTmpl string

//go:embed templates/cwindowsnative_windows_func.tmpl
var cwindowsnativeWindowsFuncTmpl string

type DLLFuncParameterTemplateParameter struct {
	Name string
}

type DLLFuncTemplateParameter struct {
	FuncPackageName string
	DLLPackageName  string
	DLLModuleName   string
	DLLName         string
	FuncName        string
	Wide            bool
	Parameters      []DLLFuncParameterTemplateParameter
	Return          bool
	ReturnError     bool
}

type WindowsFuncParameter struct {
	Name       string
	Unsafe     bool
	Type       string
	CustomType string
}

type WindowsFunc struct {
	Name          string
	Unsafe        bool
	TypeParameter string
	Parameters    WindowsFuncParameters
	ReturnType    string
	ReturnError   bool
}

type WindowsFuncParameterTemplateParameter struct {
	Name   string
	Unsafe bool
	Type   string
}
type WindowsFuncTemplateParameter struct {
	Unsafe          bool
	DLLPackageName  string
	DLLModuleName   string
	DLLName         string
	FuncName        string
	FuncPackageName string
	TypeParameter   string
	Parameters      []WindowsFuncParameterTemplateParameter
	ReturnType      string
	ReturnError     bool
}

var WindowsFuncs = map[string]WindowsFunc{
	"CloseHandle": {
		Name: "CloseHandle",
		Parameters: WindowsFuncParameters{
			{
				Name: "hObject",
				Type: "HANDLE",
			},
		},
		ReturnType:  "BOOL",
		ReturnError: true,
	},
	"ControlTrace": {
		Name:          "ControlTraceW",
		Unsafe:        true,
		TypeParameter: "TProperties PEVENT_TRACE_PROPERTIES | PEVENT_TRACE_PROPERTIES_V2",
		Parameters: WindowsFuncParameters{
			{
				Name: "TraceId",
				Type: "CONTROLTRACE_ID",
			},
			{
				Name:   "InstanceName",
				Unsafe: true,
				Type:   "LPCWSTR",
			},
			{
				Name:   "Properties",
				Unsafe: true,
				Type:   "TProperties",
			},
			{
				Name:       "ControlCode",
				Type:       "ULONG",
				CustomType: "ControlTraceCode",
			},
		},
		ReturnType: "ULONG",
	},
	"CreateToolhelp32Snapshot": {
		Name: "CreateToolhelp32Snapshot",
		Parameters: WindowsFuncParameters{
			{
				Name:       "dwFlags",
				Type:       "DWORD",
				CustomType: "TH32CS",
			},
			{
				Name: "th32ProcessID",
				Type: "DWORD",
			},
		},
		ReturnType:  "HANDLE",
		ReturnError: true,
	},
	"EnableTraceEx2": {
		Name:   "EnableTraceEx2",
		Unsafe: true,
		Parameters: []WindowsFuncParameter{
			{
				Name: "TraceId",
				Type: "CONTROLTRACE_ID",
			},
			{
				Name:   "ProviderId",
				Unsafe: true,
				Type:   "LPCGUID",
			},
			{
				Name:       "ControlCode",
				Type:       "ULONG",
				CustomType: "EventControlCode",
			},
			{
				Name:       "Level",
				Type:       "UCHAR",
				CustomType: "TraceLevel",
			},
			{
				Name: "MatchAnyKeyword",
				Type: "ULONGLONG",
			},
			{
				Name: "MatchAllKeyword",
				Type: "ULONGLONG",
			},
			{
				Name: "Timeout",
				Type: "ULONG",
			},
			{
				Name:   "EnableParameters",
				Unsafe: true,
				Type:   "PENABLE_TRACE_PARAMETERS",
			},
		},
		ReturnType: "ULONG",
	},
	"OpenProcess": {
		Name: "OpenProcess",
		Parameters: []WindowsFuncParameter{
			{
				Name:       "dwDesiredAccess",
				Type:       "DWORD",
				CustomType: "ProcessAccess",
			},
			{
				Name: "bInheritHandle",
				Type: "BOOL",
			},
			{
				Name: "dwProcessId",
				Type: "DWORD",
			},
		},
		ReturnType:  "HANDLE",
		ReturnError: true,
	},
	"OpenTraceFromRealTimeLogger": {
		Name:   "OpenTraceFromRealTimeLogger",
		Unsafe: true,
		Parameters: []WindowsFuncParameter{
			{
				Name:   "LoggerName",
				Unsafe: true,
				Type:   "PCWSTR",
			},
			{
				Name:   "Options",
				Unsafe: true,
				Type:   "*ETW_OPEN_TRACE_OPTIONS",
			},
			{
				Name:   "LogFileHeader",
				Unsafe: true,
				Type:   "*TRACE_LOGFILE_HEADER",
			},
		},
		ReturnType:  "PROCESSTRACE_HANDLE",
		ReturnError: true,
	},
	"Process32First": {
		Name:   "Process32FirstW",
		Unsafe: true,
		Parameters: []WindowsFuncParameter{
			{
				Name: "hSnapshot",
				Type: "HANDLE",
			},
			{
				Name:   "lppe",
				Unsafe: true,
				Type:   "LPPROCESSENTRY32",
			},
		},
		ReturnType:  "BOOL",
		ReturnError: true,
	},
	"Process32Next": {
		Name:   "Process32NextW",
		Unsafe: true,
		Parameters: []WindowsFuncParameter{
			{
				Name: "hSnapshot",
				Type: "HANDLE",
			},
			{
				Name:   "lppe",
				Unsafe: true,
				Type:   "LPPROCESSENTRY32",
			},
		},
		ReturnType:  "BOOL",
		ReturnError: true,
	},
	"ProcessTrace": {
		Name:   "ProcessTrace",
		Unsafe: true,
		Parameters: []WindowsFuncParameter{
			{
				Name:   "HandleArray",
				Unsafe: true,
				Type:   "*PROCESSTRACE_HANDLE",
			},
			{
				Name: "HandleCount",
				Type: "ULONG",
			},
			{
				Name:   "StartTime",
				Unsafe: true,
				Type:   "LPFILETIME",
			},
			{
				Name:   "EndTime",
				Unsafe: true,
				Type:   "LPFILETIME",
			},
		},
		ReturnType: "ULONG",
	},
	"SetPriorityClass": {
		Name: "SetPriorityClass",
		Parameters: []WindowsFuncParameter{
			{
				Name: "hProcess",
				Type: "HANDLE",
			},
			{
				Name: "dwPriorityClass",
				Type: "DWORD",
			},
		},
		ReturnType:  "BOOL",
		ReturnError: true,
	},
	"SetProcessAffinityMask": {
		Name: "SetProcessAffinityMask",
		Parameters: []WindowsFuncParameter{
			{
				Name: "hProcess",
				Type: "HANDLE",
			},
			{
				Name: "dwProcessAffinityMask",
				Type: "DWORD_PTR",
			},
		},
		ReturnType:  "BOOL",
		ReturnError: true,
	},
	"StartTrace": {
		Name:          "StartTraceW",
		Unsafe:        true,
		TypeParameter: "TProperties PEVENT_TRACE_PROPERTIES | PEVENT_TRACE_PROPERTIES_V2",
		Parameters: []WindowsFuncParameter{
			{
				Name:   "TraceId",
				Unsafe: true,
				Type:   "*CONTROLTRACE_ID",
			},
			{
				Name:   "InstanceName",
				Unsafe: true,
				Type:   "LPCWSTR",
			},
			{
				Name:   "Properties",
				Unsafe: true,
				Type:   "TProperties",
			},
		},
		ReturnType: "ULONG",
	},
}

type WindowsFuncParameters []WindowsFuncParameter

func (w WindowsFuncParameters) ToDLLFunc() []DLLFuncParameterTemplateParameter {
	result := make([]DLLFuncParameterTemplateParameter, 0, len(w))
	for _, v := range w {
		result = append(result, DLLFuncParameterTemplateParameter{
			Name: v.Name,
		})
	}
	return result
}

func (w WindowsFuncParameters) ToCWindows() []WindowsFuncParameterTemplateParameter {
	result := make([]WindowsFuncParameterTemplateParameter, 0, len(w))
	for _, v := range w {
		t := v.Type
		if v.CustomType != "" {
			t = v.CustomType
		}
		result = append(result, WindowsFuncParameterTemplateParameter{
			Name:   v.Name,
			Unsafe: v.Unsafe,
			Type:   t,
		})
	}
	return result
}

func (w WindowsFuncParameters) ToCWindowsNative() []WindowsFuncParameterTemplateParameter {
	result := make([]WindowsFuncParameterTemplateParameter, 0, len(w))
	for _, v := range w {
		result = append(result, WindowsFuncParameterTemplateParameter{
			Name:   v.Name,
			Unsafe: v.Unsafe,
			Type:   v.Type,
		})
	}
	return result
}

func (windowsFunc WindowsFunc) generateWindowsFunc(dllName string) {
	dllFuncTmpl := template.Must(template.New("dll_func.tmpl").Parse(dllFuncTmpl))
	cwindowsTmpl := template.Must(template.New("cwindows_windows_func.tmpl").Parse(cwindowsWindowsFuncTmpl))
	cwindowsnativeTmpl := template.Must(template.New("cwindowsnative_windows_func.tmpl").Parse(cwindowsnativeWindowsFuncTmpl))

	funcName := strings.TrimSuffix(windowsFunc.Name, "W")

	if err := os.Mkdir(filepath.Join("dll", dllName, funcName), 0644); err != nil {
		panic(err)
	}
	dllFuncFile, err := os.Create(filepath.Join("dll", dllName, funcName, funcName+".go"))
	if err != nil {
		panic(err)
	}
	defer dllFuncFile.Close()
	if err := dllFuncTmpl.Execute(dllFuncFile, DLLFuncTemplateParameter{
		FuncPackageName: strings.ToLower(funcName),
		DLLPackageName:  strings.ToLower(dllName),
		DLLModuleName:   dllName,
		DLLName:         ToUpperFirst(dllName),
		FuncName:        funcName,
		Wide:            strings.HasSuffix(windowsFunc.Name, "W"),
		Parameters:      windowsFunc.Parameters.ToDLLFunc(),
		Return:          windowsFunc.ReturnType != "",
		ReturnError:     windowsFunc.ReturnError,
	}); err != nil {
		panic(err)
	}

	cwindowsFile, err := os.Create(filepath.Join("cwindows", fmt.Sprintf("windows_func_%s.go", funcName)))
	if err != nil {
		panic(err)
	}
	defer cwindowsFile.Close()
	if err := cwindowsTmpl.Execute(cwindowsFile, WindowsFuncTemplateParameter{
		Unsafe:          windowsFunc.Unsafe,
		DLLPackageName:  strings.ToLower(dllName),
		DLLModuleName:   dllName,
		DLLName:         ToUpperFirst(dllName),
		FuncName:        funcName,
		FuncPackageName: strings.ToLower(funcName),
		TypeParameter:   windowsFunc.TypeParameter,
		Parameters:      windowsFunc.Parameters.ToCWindows(),
		ReturnType:      windowsFunc.ReturnType,
		ReturnError:     windowsFunc.ReturnError,
	}); err != nil {
		panic(err)
	}
	cwindowsnativeFile, err := os.Create(filepath.Join("cwindowsnative", fmt.Sprintf("windows_func_%s.go", funcName)))
	if err != nil {
		panic(err)
	}
	defer cwindowsnativeFile.Close()
	if err := cwindowsnativeTmpl.Execute(cwindowsnativeFile, WindowsFuncTemplateParameter{
		Unsafe:          windowsFunc.Unsafe,
		DLLPackageName:  strings.ToLower(dllName),
		DLLModuleName:   dllName,
		DLLName:         ToUpperFirst(dllName),
		FuncName:        funcName,
		FuncPackageName: strings.ToLower(funcName),
		TypeParameter:   windowsFunc.TypeParameter,
		Parameters:      windowsFunc.Parameters.ToCWindowsNative(),
		ReturnType:      windowsFunc.ReturnType,
		ReturnError:     windowsFunc.ReturnError,
	}); err != nil {
		panic(err)
	}
}

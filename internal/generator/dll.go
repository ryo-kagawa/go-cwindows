package main

import (
	_ "embed"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/dll.tmpl
var dllTmpl string

type DLL struct {
	FileName     string
	WinodwsFuncs []WindowsFunc
}

type DLLTemplate struct {
	DLLPackageName string
	DLLFileName    string
	DLLModuleName  string
}

var DLLs = []DLL{
	{
		FileName: "kernel32.dll",
		WinodwsFuncs: []WindowsFunc{
			WindowsFuncs["CloseHandle"],
			WindowsFuncs["CreateToolhelp32Snapshot"],
			WindowsFuncs["OpenProcess"],
			WindowsFuncs["Process32First"],
			WindowsFuncs["Process32Next"],
			WindowsFuncs["SetPriorityClass"],
			WindowsFuncs["SetProcessAffinityMask"],
		},
	},
	{
		FileName: "sechost.dll",
		WinodwsFuncs: []WindowsFunc{
			WindowsFuncs["ControlTrace"],
			WindowsFuncs["EnableTraceEx2"],
			WindowsFuncs["OpenTraceFromRealTimeLogger"],
			WindowsFuncs["ProcessTrace"],
			WindowsFuncs["StartTrace"],
		},
	},
}

func generateDLL() {
	dllTmpl := template.Must(template.New("dll.tmpl").Parse(dllTmpl))

	for _, dll := range DLLs {
		moduleName := strings.TrimSuffix(dll.FileName, ".dll")
		os.Mkdir(filepath.Join("dll", moduleName), 0644)
		dllFile, err := os.Create(filepath.Join("dll", moduleName, moduleName+".go"))
		if err != nil {
			panic(err)
		}
		defer dllFile.Close()
		if err := dllTmpl.Execute(dllFile, DLLTemplate{
			DLLPackageName: strings.ToLower(moduleName),
			DLLModuleName:  "Mod" + ToUpperFirst(moduleName),
			DLLFileName:    dll.FileName,
		}); err != nil {
			panic(err)
		}
		for _, windowsFunc := range dll.WinodwsFuncs {
			windowsFunc.generateWindowsFunc(moduleName)
		}
	}
}

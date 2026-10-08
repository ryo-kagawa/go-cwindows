package sechost

import "golang.org/x/sys/windows"

var ModSechost = windows.NewLazySystemDLL("Sechost.dll")

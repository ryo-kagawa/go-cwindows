package main

import "os"

func main() {
	if err := os.RemoveAll("cwindows"); err != nil {
		panic(err)
	}
	if err := os.RemoveAll("cwindowsnative"); err != nil {
		panic(err)
	}
	if err := os.RemoveAll("dll"); err != nil {
		panic(err)
	}

	if err := os.Mkdir("cwindows", 0644); err != nil {
		panic(err)
	}
	if err := os.Mkdir("cwindowsnative", 0644); err != nil {
		panic(err)
	}
	if err := os.Mkdir("dll", 0644); err != nil {
		panic(err)
	}
	generateWindowsType()
	generateWindowsStructType()
	generateCustomType()
	generateCustomStructType()
	generateWindowsConst()
	generateCustomConst()
	generateDLL()
	generateCustomFunc()
}

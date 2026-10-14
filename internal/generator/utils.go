package main

import "strings"

func ToUpperFirst(value string) string {
	return strings.ToUpper(value[:1]) + value[1:]
}

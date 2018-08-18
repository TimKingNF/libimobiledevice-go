package utils

import (
	"C"
	"unsafe"
)

// GoStrings converts C *char array to Go string slice
func GoStrings(argc int, argv unsafe.Pointer) []string {
	tmpslice := (*[1 << 30]*C.char)(argv)[:argc:argc]
	gostrings := make([]string, argc)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}
	return gostrings
}

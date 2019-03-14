package main

import (
	"os"
	"strings"
)

import "C"

//export Args
func Args() *C.char {
	args := os.Args
	return C.CString(strings.Join(args, " "))
}

//export Environ
func Environ() *C.char {
	envp := os.Environ()
	return C.CString(strings.Join(envp, " "))
}

//export Executable
func Executable() *C.char {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return C.CString(exePath)
}

func main() {}

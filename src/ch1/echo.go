package main

import (
	"os"
	"strings"
)

func main() {
	result, sep := "", ""
	for _, arg := range os.Args {
		result += sep + arg
		sep = " "
	}
	println(result)

	println(strings.Join(os.Args, " "))
}

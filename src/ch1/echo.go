package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"strings"
)

func euler() {
	fmt.Printf("%.3f", cmplx.Exp(1i*math.Pi)+1)
}

const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

func main() {
	result, sep := "", ""
	for _, arg := range os.Args {
		result += sep + arg
		sep = " "

	}
	println(result)
	println(b, kb, mb, gb, tb, pb)

	println(strings.Join(os.Args, " "))
	euler()
}

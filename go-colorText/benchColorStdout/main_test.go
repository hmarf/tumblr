package main

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
	gcolor "github.com/gookit/color"
)

// https://github.com/fatih/color
func Benchmark_fatih_color(t *testing.B) {
	color.Cyan("Prints text in cyan.")
}

// https://github.com/gookit/color
func Benchmark_gookit_color(t *testing.B) {
	gcolor.Cyan.Println("Prints text in cyan.")
}

// https://groups.google.com/g/golang-nuts/c/nluStAtr8NA?pli=1
func Benchmark_normalFmt_color(t *testing.B) {
	fmt.Printf("\x1b[36m%s\x1b[0m", "Prints text in cyan.")
}

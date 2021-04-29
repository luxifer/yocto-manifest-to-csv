package main

import (
	"flag"
	"fmt"
	"os"

	parser "github.com/luxifer/yocto-manifest-to-csv"
)

var inputFile string
var outputFile string

func init() {
	flag.StringVar(&inputFile, "input", "", "path of the yocto file to parse")
	flag.StringVar(&outputFile, "output", "", "path of the CSV file to write, default to stdout")
}

func main() {
	flag.Parse()

	fi, err := os.Open(inputFile)
	if err != nil {
		printErr("Unable to open file: %s", err)
		os.Exit(1)
	}

	fo := os.Stdout

	if outputFile != "" {
		fo, err = os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
		if err != nil {
			printErr("Unable to open file: %s", err)
			os.Exit(1)
		}
		defer fo.Close()
	}

	_ = parser.Parse(fi, fo)
}

func printErr(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
}

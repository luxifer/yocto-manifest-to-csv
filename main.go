package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
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

	defer fi.Close()

	s := bufio.NewScanner(fi)

	var list [][]string
	lastItem := make([]string, 0)
	header := make([]string, 0)
	var headerDone bool

	for s.Scan() {
		line := s.Text()

		if line == "" {
			list = append(list, lastItem)
			lastItem = make([]string, 0)
			headerDone = true
			continue
		}

		splits := strings.SplitN(line, ":", 2)
		lastItem = append(lastItem, strings.TrimLeft(splits[1], " "))

		if !headerDone {
			header = append(header, splits[0])
		}
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

	w := csv.NewWriter(fo)
	_ = w.Write(header)

	for _, v := range list {
		_ = w.Write(v)
	}

	w.Flush()
}

func printErr(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
}

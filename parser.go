package parser

import (
	"bufio"
	"encoding/csv"
	"io"
	"strings"
)

func Parse(in io.Reader, out io.Writer) error {
	s := bufio.NewScanner(in)

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

	w := csv.NewWriter(out)
	err := w.Write(header)
	if err != nil {
		return err
	}

	return w.WriteAll(list)
}

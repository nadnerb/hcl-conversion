package main

import (
	"bytes"
	"strings"
)

func output (results []*Result) (string, error) {

	var all []string

	for _, n := range results {
		var buffer bytes.Buffer
		buffer.WriteString(n.Name)
		buffer.WriteString("=")
		buffer.WriteString(n.Value)
		all = append(all, buffer.String())
	}
	return strings.Join(all, " "), nil
}


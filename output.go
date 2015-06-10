package main

import (
	"bytes"
	"strings"
)

func output (values []*ConfigValue) (string, error) {

	var all []string

	for _, n := range values {
		var buffer bytes.Buffer
		buffer.WriteString(n.Name)
		buffer.WriteString("=")
		buffer.WriteString(n.Value)
		all = append(all, buffer.String())
	}
	return strings.Join(all, " "), nil
}


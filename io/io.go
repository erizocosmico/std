package io

import (
	"bufio"
	"os"
	"strings"
)

// Readln reads a single line from the standard input.
func Readln() (string, error) {
	ln, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(ln)), nil
}

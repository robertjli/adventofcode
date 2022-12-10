package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadFile returns the specified file in its entirety, as a string slice with one line per element.
func ReadFile(path string) []string {
	b, err := os.ReadFile(path)
	FailIf(err, "unable to read file", path)
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}

// NewScanner returns a Scanner to read lines from the specified file
func NewScanner(path string) (*bufio.Scanner, func()) {
	f, err := os.Open(path)
	FailIf(err, "unable to open file", path)

	return bufio.NewScanner(f), func() { _ = f.Close() }
}

// ProcessLines calls the process() function on each line in the specified file.
func ProcessLines(path string, process func(line string)) {
	scanner, closeFunc := NewScanner(path)
	defer closeFunc()

	for scanner.Scan() {
		process(scanner.Text())
	}
}

// ParseInt converts the string s to an int, or fails
func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	FailIf(err)
	return i
}

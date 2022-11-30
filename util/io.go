package util

import (
	"bufio"
	"os"
	"strings"
)

// ReadFile returns the specified file in its entirety, as a string slice with one line per element.
func ReadFile(path string) []string {
	b, err := os.ReadFile(path)
	FailIf(err, "unable to read file", path)
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}

// ProcessLines calls the process() function on each line in the specified file.
func ProcessLines(path string, process func(string)) {
	f, err := os.Open(path)
	FailIf(err, "unable to open file", path)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		process(scanner.Text())
	}

	_ = f.Close()
}

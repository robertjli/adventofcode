package util

import (
	"fmt"
	"strings"
)

// FailIf panics if it encounters an error.
func FailIf(err error, msg ...string) {
	if err != nil {
		if len(msg) > 0 {
			err = fmt.Errorf("%s: %w", strings.Join(msg, " "), err)
		}
		panic(err)
	}
}

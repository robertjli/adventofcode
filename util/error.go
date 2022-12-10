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

// Assert panics if actual != expected
func Assert(actual, expected any) {
	if actual != expected {
		panic(fmt.Sprintf("Assertion failed, expected %v but got %v", expected, actual))
	}
	fmt.Print(" ✅ ")
}

// AssertSlice panics if not all elements in actual == elements in expected
func AssertSlice[T comparable](actual, expected []T) {
	if len(actual) != len(expected) {
		panic(fmt.Sprintf("Assertion failed, lengths differ: expected %d but got %d",
			len(expected), len(actual)))
	}
	for i := range actual {
		if actual[i] != expected[i] {
			panic(fmt.Sprintf("Assertion failed, element %d differs: expected %v but got %v",
				i, expected[i], actual[i]))
		}
	}
	fmt.Print(" ✅ ")
}

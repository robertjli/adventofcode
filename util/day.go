package util

import "fmt"

func DayPath(year, day int) string {
	return fmt.Sprintf("%d/day%02d/", year, day)
}

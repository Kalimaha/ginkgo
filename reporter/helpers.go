package reporter

import (
	"fmt"
)

type rspecReporter struct {
	leaves map[string]Leaf
}

type Leaf struct {
	Description string
	Leaves      map[string]Leaf
	Passed      int
	Level       int
	Duration    string
}

var (
	Green = Color("\033[1;32m%s\033[0m")
	Red   = Color("\033[1;31m%s\033[0m")
)

func icon(passed int) (out string) {
	if passed == 1 {
		return Green("✔️")
	} else if passed == -1 {
		return Red("✗")
	}

	return out
}

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString, fmt.Sprint(args...))
	}
	return sprint
}

func bool2int(passed bool) int {
	if passed == true {
		return 1
	}
	return -1
}

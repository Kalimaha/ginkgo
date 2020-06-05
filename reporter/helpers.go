package reporter

import (
	"errors"
	"fmt"
	"github.com/onsi/ginkgo/types"
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
	Link		string
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

func newLeaf(description string, level int, spec *types.SpecSummary) Leaf {
	return Leaf{
		Description: description,
		Leaves: make(map[string]Leaf),
		Passed: bool2int(spec.Passed()),
		Level: level,
		Duration: spec.RunTime.String(),
		Link: spec.ComponentCodeLocations[2].String(),
	}
}

func FindParent(leaves map[string]Leaf, keys []string, currentParent Leaf) (out Leaf, err error) {
	currentLeaf, ok := leaves[keys[0]]

	if !ok {
		return currentParent, err
	}

	if len(keys) == 1 && currentLeaf.Description == keys[0] {
		return currentParent, err
	}

	if len(keys) > 1 {
		_, tail := keys[0], keys[1:]
		return FindParent(currentLeaf.Leaves, tail, currentLeaf)
	}

	return out, errors.New("boom")
}


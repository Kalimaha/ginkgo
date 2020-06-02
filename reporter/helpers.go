package reporter

import (
	"fmt"
	"github.com/onsi/ginkgo/types"
	"strings"
)

type rspecReporter struct {
	leaves []Leaf
}

type Leaf struct {
	Description string
	Leaves      []Leaf
	Passed      int
}

var (
	Green = Color("\033[1;32m%s\033[0m")
	Red   = Color("\033[1;31m%s\033[0m")
)

func seekAndInsert(leaves *[]Leaf, spec *types.SpecSummary, level int) {
	found := false
	for _, leaf := range *leaves {
		if leaf.Description == formatOut(spec.ComponentTexts[level], level) {
			found = true
			break
		}
	}

	if len(spec.ComponentTexts) > level && !found {
		*leaves = append(*leaves, newLeaf(spec, level))
	}
}

func newLeaf(spec *types.SpecSummary, level int) Leaf {
	return Leaf{Description: formatOut(spec.ComponentTexts[level], level), Passed: bool2int(spec.Passed())}
}

func formatOut(msg string, level int) string {
	blanks := strings.Builder{}
	for i := 0; i < level; i++ {
		blanks.WriteString("  ")
	}
	blanks.WriteString("%s")

	return fmt.Sprintf(blanks.String(), msg)
}

func icon(passed int) (out string) {
	if passed == 1 {
		return Green("✔️ ")
	} else if passed == -1 {
		return Red("✗ ")
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

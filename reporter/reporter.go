package reporter

import (
	"fmt"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

func New() *rspecReporter {
	return &rspecReporter{
		leaves: map[string]Leaf{},
	}
}

func (r *rspecReporter) SpecSuiteWillBegin(config.GinkgoConfigType, *types.SuiteSummary) {}
func (r *rspecReporter) BeforeSuiteDidRun(*types.SetupSummary)                           {}
func (r *rspecReporter) AfterSuiteDidRun(*types.SetupSummary)                            {}
func (r *rspecReporter) SpecWillRun(*types.SpecSummary)                                  {}

func (r *rspecReporter) SpecDidComplete(spec *types.SpecSummary) {
	var keys []string
	keysMap := make(map[int]string)

	for i := 1; i < len(spec.ComponentTexts); i++ {
		keys = append(keys, spec.ComponentTexts[i])
		keysMap[i] = spec.ComponentTexts[i]
	}

	lastLeaf, ok := r.leaves[spec.ComponentTexts[1]]
	if !ok {
		r.leaves[spec.ComponentTexts[1]] = newLeaf(spec.ComponentTexts[1], 1, spec)
	}

	for i := 2; i <= len(keys); i++ {
		lastLeaf, ok = lastLeaf.Leaves[spec.ComponentTexts[i]]
		if !ok {
			var parentKeys []string
			for j := 0; j < i; j++ {
				parentKeys = append(parentKeys, keys[j])
			}
			parentLeaf, _ := FindParent(r.leaves, parentKeys, Leaf{})
			parentLeaf.Leaves[spec.ComponentTexts[i]] = newLeaf(spec.ComponentTexts[i], i, spec)
		}
	}
}

func (r *rspecReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	fmt.Println(summary.SuiteDescription)
	PrintLeaves(r.leaves)
	fmt.Println()
}

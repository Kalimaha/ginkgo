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
	keysMap := make(map[int]string)
	for i := 1; i < len(spec.ComponentTexts); i++ {
		keysMap[i] = spec.ComponentTexts[i]
	}

	for i := 1; i < len(spec.ComponentTexts); i++ {
		if i == 1 {
			key1 := keysMap[1]
			_, ok := r.leaves[key1]
			if !ok {
				r.leaves[key1] = newLeaf(key1, 1, spec)
			}
		}

		if i == 2 {
			key1 := keysMap[1]
			key2 := keysMap[2]

			_, ok := r.leaves[key1].Leaves[key2]
			if !ok {
				r.leaves[key1].Leaves[key2] = newLeaf(key2, 2, spec)
			}
		}

		if i == 3 {
			key1 := keysMap[1]
			key2 := keysMap[2]
			key3 := keysMap[3]

			_, ok := r.leaves[key1].Leaves[key2].Leaves[key3]
			if !ok {
				r.leaves[key1].Leaves[key2].Leaves[key3] = newLeaf(key3, 3, spec)
			}
		}

		if i == 4 {
			key1 := keysMap[1]
			key2 := keysMap[2]
			key3 := keysMap[3]
			key4 := keysMap[4]

			_, ok := r.leaves[key1].Leaves[key2].Leaves[key3].Leaves[key4]
			if !ok {
				r.leaves[key1].Leaves[key2].Leaves[key3].Leaves[key4] = newLeaf(key4, 4, spec)
			}
		}
	}
}

func newLeaf(description string, level int, spec *types.SpecSummary) Leaf {
	return Leaf{
		Description: description,
		Leaves: make(map[string]Leaf),
		Passed: bool2int(spec.Passed()),
		Level: level,
		Duration: spec.RunTime.String(),
	}
}

func (r *rspecReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	fmt.Println(summary.SuiteDescription)
	PrintLeaves(r.leaves)
	fmt.Println()
}



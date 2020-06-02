package reporter

import (
	"fmt"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

func New() *rspecReporter {
	return &rspecReporter{
		leaves: []Leaf{},
	}
}

func (r *rspecReporter) SpecSuiteWillBegin(config.GinkgoConfigType, *types.SuiteSummary) {}
func (r *rspecReporter) BeforeSuiteDidRun(*types.SetupSummary)                           {}
func (r *rspecReporter) AfterSuiteDidRun(*types.SetupSummary)                            {}
func (r *rspecReporter) SpecWillRun(*types.SpecSummary)                                  {}

func (r *rspecReporter) SpecDidComplete(spec *types.SpecSummary) {
	seekAndInsert(&r.leaves, spec, 1)

	for idx1, v1 := range r.leaves {
		if v1.Description == formatOut(spec.ComponentTexts[1], 1) {
			seekAndInsert(&r.leaves[idx1].Leaves, spec, 2)
		}
	}

	for idx1, v1 := range r.leaves {
		if v1.Description == formatOut(spec.ComponentTexts[1], 1) {
			for idx2, v2 := range r.leaves[idx1].Leaves {
				if v2.Description == formatOut(spec.ComponentTexts[2], 2) {
					seekAndInsert(&r.leaves[idx1].Leaves[idx2].Leaves, spec, 3)
				}
			}
		}
	}

	for idx1, v1 := range r.leaves {
		if v1.Description == formatOut(spec.ComponentTexts[1], 1) {
			for idx2, v2 := range r.leaves[idx1].Leaves {
				if v2.Description == formatOut(spec.ComponentTexts[2], 2) {
					for idx3, v3 := range r.leaves[idx1].Leaves[idx2].Leaves {
						if v3.Description == formatOut(spec.ComponentTexts[3], 3) {
							seekAndInsert(&r.leaves[idx1].Leaves[idx2].Leaves[idx3].Leaves, spec, 4)
						}
					}
				}
			}
		}
	}
}

func (r *rspecReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	fmt.Println()
	fmt.Println(summary.SuiteDescription)
	PrintLeaves(r.leaves)
	fmt.Println()
}

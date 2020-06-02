# RSpec-like custom reporter for Ginkgo

Custom reporter for [Ginkgo](http://onsi.github.io/ginkgo/) testing framework.

## Usage

```go
package example_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/Kalimaha/ginkgo/reporter"
)

func TestExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "Example Test", []Reporter{reporter.New()})
}
```
# RSpec-like custom reporter for Ginkgo

Custom reporter for [Ginkgo](http://onsi.github.io/ginkgo/) testing framework.

## Usage

Just add `github.com/Kalimaha/ginkgo/reporter` to the imports:

```go
package calculator

import (
	"github.com/Kalimaha/ginkgo/reporter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestReducers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "Calculator", []Reporter{reporter.New()})
}
```

## Example

The `calculator_test.go` included in this repo will produce the following output:

```bash
Calculator
  Divide
    divides two numbers, A and B ✔️ 
    but when B is 0
      returns an error ✔️ 
  Sum
    sums two numbers ✔️ 

PASS
```
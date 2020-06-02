package calculator

import (
	"../reporter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestReducers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "Calculator", []Reporter{reporter.New()})
}

var _ = Describe("Sum", func() {
	var a = 2
	var b = 3

	It("sums two numbers", func() {
		Expect(Sum(a, b)).To(Equal(5))
	})
})

var _ = Describe("Divide", func() {
	var a = float32(9)
	var b = float32(3)
	var err, result = Divide(a, b)

	It("divides two numbers, A and B", func() {
		Expect(result).To(Equal(float32(3)))
	})

	Context("but when B is 0", func() {
		BeforeEach(func() {
			b = 0
			err, _ = Divide(a, b)
		})

		It("returns an error", func() {
			Expect(err.Error()).To(Equal("division by 0"))
		})
	})
})

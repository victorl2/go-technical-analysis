package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MEAN", func() {
	It("should return the correct value", func() {
		in := []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0}
		Expect(MEAN(in)).To(Equal(4.5))
	})
})

var _ = Describe("SMA", func() {
	It("should return the correct values", func() {
		in := []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0}
		out := []float64{2.0, 3.0, 4.0, 5.0, 6.0, 7.0}
		Expect(SMA(in, 5)).To(Equal(out))
	})

	It("should return a nil array the values slice is shorter than the period", func() {
		in := []float64{0.0}
		var out = []float64{0.0}
		Expect(SMA(in, 5)).To(Equal(out))
	})
})

var _ = Describe("WMA", func() {
	It("should return the correct values", func() {
		in := []float64{10.0, 11.0, 12.0}
		out := []float64{11.333333333333332}
		Expect(WMA(in, 3)).To(Equal(out))
	})
})

var _ = Describe("EMA", func() {
	It("should return the correct values", func() {
		in := []float64{26.0, 54.0, 8.0, 77.0, 61.0, 39.0, 44.0, 91.0, 98.0, 17.0}
		out := []float64{45.2, 43.13333333333333, 43.422222222222224, 59.28148148148148, 72.18765432098765, 53.7917695473251}
		Expect(EMA(in, 5)).To(Equal(out))
	})
})

var _ = Describe("DEMA", func() {
	It("should return the correct values", func() {
		in := []float64{26.0, 54.0, 8.0, 77.0, 61.0, 39.0, 44.0, 91.0, 98.0, 17.0}
		out := []float64{91.73037037037035, 54.55632373113854}
		Expect(DEMA(in, 5)).To(Equal(out))
	})
})

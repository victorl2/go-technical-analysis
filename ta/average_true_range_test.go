package ta

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ATR", func() {
	It("should return the correct values", func() {
		in := []Candlestick{
			{High: 1.1, Low: 0.9, Close: 1.0},
			{High: 2.1, Low: 1.9, Close: 2.0},
			{High: 3.1, Low: 2.9, Close: 3.0},
			{High: 4.1, Low: 3.9, Close: 4.0},
			{High: 5.2, Low: 4.9, Close: 5.0},
		}
		out := []float64{1.0999999999999999, 1.1333333333333333}
		result := ATR(in, 3)
		Expect(result).To(Equal(out))
	})
})

var _ = Describe("trueRanges", func() {
	It("should return the correct values", func() {
		in := []Candlestick{
			{High: 1.1, Low: 0.9, Close: 1.0},
			{High: 2.1, Low: 1.9, Close: 2.0},
			{High: 3.1, Low: 2.9, Close: 3.0},
			{High: 4.1, Low: 3.9, Close: 4.0},
			{High: 5.2, Low: 4.9, Close: 5.0},
		}
		out := []float64{1.1, 1.1, 1.0999999999999996, 1.2000000000000002}
		result := trueRanges(in)
		Expect(result).To(Equal(out))
	})
})

var _ = Describe("trueRange", func() {
	It("should return the correct values", func() {
		curHigh := 1.0
		curLow := 0.5
		prevClose := 0.7
		out := 0.5
		result := trueRange(curHigh, curLow, prevClose)
		Expect(result).To(Equal(out))
	})
})

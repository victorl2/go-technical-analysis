package ta

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ATR", func() {
	It("should return the correct values", func() {
		in := []OHLCV{
			&Candlestick{open: 0, high: 1.1, low: 0.9, close: 1.0},
			&Candlestick{open: 0, high: 2.1, low: 1.9, close: 2.0},
			&Candlestick{open: 0, high: 3.1, low: 2.9, close: 3.0},
			&Candlestick{open: 0, high: 4.1, low: 3.9, close: 4.0},
			&Candlestick{open: 0, high: 5.2, low: 4.9, close: 5.0},
		}
		out := []float64{1.0999999999999999, 1.1333333333333333}
		result := ATR(in, 3)
		Expect(result).To(Equal(out))
	})
})

var _ = Describe("trueRanges", func() {
	It("should return the correct values", func() {
		in := []OHLCV{
			&Candlestick{open: 0, high: 1.1, low: 0.9, close: 1.0},
			&Candlestick{open: 0, high: 2.1, low: 1.9, close: 2.0},
			&Candlestick{open: 0, high: 3.1, low: 2.9, close: 3.0},
			&Candlestick{open: 0, high: 4.1, low: 3.9, close: 4.0},
			&Candlestick{open: 0, high: 5.2, low: 4.9, close: 5.0},
		}
		out := []float64{1.1, 1.1, 1.0999999999999996, 1.2000000000000002}
		result := trueRanges(in)
		Expect(result).To(Equal(out))
	})
})

var _ = Describe("trueRange", func() {
	It("should return the correct values", func() {
		curhigh := 1.0
		curlow := 0.5
		prevclose := 0.7
		out := 0.5
		result := trueRange(curhigh, curlow, prevclose)
		Expect(result).To(Equal(out))
	})
})

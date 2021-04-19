package ta

import (
	"math"
)

type Macd struct {
	MacdLine   []float64
	SignalLine []float64
	Histogram  []float64
}

func MACDDefault(prices []float64) *Macd {
	return MACD(prices, 26, 12, 9)
}

//MACD calculates the moving average convergence divergence
//returns the macd line, signal line and the histogram
func MACD(prices []float64, slowPeriod, fastPeriod, signalPeriod int) *Macd {
	if len(prices) < slowPeriod || slowPeriod < fastPeriod {
		return nil
	}
	macd := Macd{}

	var slowEMA = EMA(prices, slowPeriod)
	var fastEMA = EMA(prices, fastPeriod)

	macd.MacdLine = subtract(fastEMA, slowEMA)
	macd.SignalLine = EMA(macd.MacdLine, signalPeriod)
	macd.Histogram = subtract(macd.MacdLine, macd.SignalLine)
	macd.MacdLine, macd.SignalLine = normalizeSize(macd.MacdLine, macd.SignalLine)
	return &macd
}

func subtract(valuesA []float64, valuesB []float64) []float64 {
	var first, second = normalizeSize(valuesA, valuesB)
	var result []float64
	for i, valueFirst := range first {
		result = append(result, valueFirst-second[i])
	}
	return result
}

func normalizeSize(valuesA []float64, valuesB []float64) ([]float64, []float64) {
	offsetA := int(math.Max(0, float64(len(valuesA)-len(valuesB))))
	offsetB := int(math.Max(0, float64(len(valuesB)-len(valuesA))))
	return valuesA[offsetA:], valuesB[offsetB:]
}

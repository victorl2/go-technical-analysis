package main

import (
	"math"
)

// ATR calculates the Average True Range of an slice of Candlesticks for a given period
func ATR(values []Candlestick, period int) []float64 {
	var atrs []float64
	startIdx := period
	endIdx := len(values) - 1
	tr := trueRanges(values)
	sma := SMA(tr[:startIdx], period)
	atrs = append(atrs, sma[len(sma)-1])
	for i := startIdx; i < endIdx; i++ {
		lastAtr := atrs[len(atrs)-1]
		atr := ((lastAtr * float64(period-1)) + tr[i]) / float64(period)
		atrs = append(atrs, atr)
	}
	return atrs
}

func trueRanges(values []Candlestick) []float64 {
	var result []float64
	for idx, candle := range values[1:] {
		tr := trueRange(candle.High, candle.Low, values[idx].Close)
		result = append(result, tr)
	}
	return result
}

func trueRange(curHigh float64, curLow float64, prevClose float64) float64 {
	tempMax := math.Max(curHigh-curLow, math.Abs(curHigh-prevClose))
	return math.Max(tempMax, math.Abs(curLow-prevClose))
}

package ta

import "math"

func toCandles(values []float64) []OHLCV {
	var candles []OHLCV
	for _, value := range values {
		candles = append(candles, &Candlestick{close: value})
	}
	return candles
}

func roundedValues(values []float64, roundOn float64, places int) []float64 {
	var roundedVals []float64
	for _, el := range values {
		roundedVals = append(roundedVals, round(el, roundOn, places))
	}
	return roundedVals
}

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

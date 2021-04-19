package main

//OBV calculates the On Balance Volume indicator
func OBV(values []Candlestick) []float64 {
	var result []float64
	var acc = 0.0
	for i, candle := range values[1:] {
		if values[i].Close > candle.Close {
			acc = acc - candle.Volume
			result = append(result, acc)
		} else if values[i].Close < candle.Close {
			acc = acc + candle.Volume
			result = append(result, acc)
		} else {
			result = append(result, acc)
		}
	}
	return result
}

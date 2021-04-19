package main

// Stochastic calculates the Stochastic Oscillator for the candlestick for the given period
func Stochastic(candles []Candlestick, kPeriod, dPeriod int) ([]float64, []float64) {
	var kLine []float64
	for i, candle := range candles[kPeriod-1:] {
		var highest, lowest = limits(candles[i : int(kPeriod-1)+i])
		kLine = append(kLine, (candle.Close-lowest)/(highest-lowest)*100)
	}
	return kLine, SMA(kLine, dPeriod)
}

//Limits returns the highest and lowest values present in candlestick slice
func limits(candles []Candlestick) (float64, float64) {
	if len(candles) == 0 {
		return -1, -1
	}

	var highest, lowest = candles[0].High, candles[0].Low

	for _, candle := range candles {
		if candle.High > highest {
			highest = candle.High
		}

		if candle.Low < lowest {
			lowest = candle.Low
		}
	}
	return highest, lowest
}

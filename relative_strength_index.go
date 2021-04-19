package main

import "math"

// RSI calculates the Relative Strength for the candlesticks in the provided period
func RSI(candles []Candlestick, period int) []float64 {
	var up []float64
	var down []float64
	for i, candle := range candles[1:] {
		previous := candles[i]
		if candle.Close >= previous.Close {
			up = append(up, candle.Close-previous.Close)
			down = append(down, 0)
		} else {
			down = append(down, previous.Close-candle.Close)
			up = append(up, 0)
		}
	}
	var rs []float64
	var rsi []float64
	prevAvgUp := MEAN(up[0 : period-1])
	prevAvgDown := MEAN(down[0 : period-1])

	for i := period - 1; i < len(up); i++ {
		currAvgUp := (prevAvgUp*float64(period-1) + up[i]) / float64(period)
		currAvgDown := (prevAvgDown*float64(period-1) + down[i]) / float64(period)
		currRS := currAvgUp / math.Abs(currAvgDown)
		rs = append(rs, currRS)
		currRSI := 100 - 100/(1+currRS)
		rsi = append(rsi, currRSI)
		prevAvgUp = currAvgUp
		prevAvgDown = currAvgDown
	}
	return rsi
}

package ta

import "math"

// RSI calculates the Relative Strength for the closePricesticks in the provided period
func RSI(prices []float64, period int) []float64 {
	var up []float64
	var down []float64
	for i, closePrice := range prices[1:] {
		previous := prices[i]
		if closePrice >= previous {
			up = append(up, closePrice-previous)
			down = append(down, 0)
		} else {
			down = append(down, previous-closePrice)
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

func LaguerreRSIDefault(vals []float64) []float64 {
	return LaguerreRSI(vals, 0.5)
}

// LaguerreRSI calculates the Laguerre Relative Strength for the closePricesticks in the provided period
// from paper: http://mesasoftware.com/papers/TimeWarp.pdf
func LaguerreRSI(prices []float64, gamma float64) []float64 {
	l0s := make([]float64, len(prices))
	l1s := make([]float64, len(prices))
	l2s := make([]float64, len(prices))
	l3s := make([]float64, len(prices))
	rsi := make([]float64, len(prices))

	for i := 1; len(prices)-1 < i; i++ {
		l0s[i] = (1-gamma)*prices[i] + gamma*l0s[i-1]
		l1s[i] = -gamma*l0s[i] + l0s[i-1] + gamma*l1s[i-1]
		l2s[i] = -gamma*l1s[i] + l1s[i-1] + gamma*l2s[i-1]
		l3s[i] = -gamma*l2s[i] + l2s[i-1] + gamma*l3s[i-1]

		var cu float64
		var cd float64
		if l0s[i] >= l1s[i] {
			cu = l0s[i] - l1s[i]
		} else {
			cd = l1s[i] - l0s[i]
		}
		if l1s[i] >= l2s[i] {
			cu += +l1s[i] - l2s[i]
		} else {
			cd += l2s[i] - l1s[i]
		}
		if l2s[i] >= l3s[i] {
			cu += l2s[i] - l3s[i]
		} else {
			cd += l3s[i] - l2s[i]
		}

		if cu+cd != 0 {
			rsi[i] = cu / (cu + cd)
		}
	}
	return rsi
}

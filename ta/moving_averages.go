package ta

// MEAN calculates the mean value for values provided
func MEAN(values []float64) float64 {
	var total float64 = 0
	for _, element := range values {
		total += element
	}
	return total / float64(len(values))
}

// SMA calculates the Simple Moving Average for the provided period
func SMA(values []float64, period int) []float64 {
	if period == 0 || len(values) < period {
		return values
	}
	var sma []float64
	for index := range values {
		indexPlusOne := index + 1
		if indexPlusOne >= period {
			avg := MEAN(values[indexPlusOne-period : indexPlusOne])
			sma = append(sma, avg)
		}
	}
	return sma
}

// WMA calculates the Weighted Moving Average for the provided period
func WMA(values []float64, period int) []float64 {
	var result []float64
	for index := range values {
		indexPlusOne := index + 1
		if indexPlusOne >= period {
			var res []float64
			sl := values[indexPlusOne-period : indexPlusOne]

			// Get the sum of the number of entries
			var sum float64 = 0
			for i := range sl {
				sum += float64(i + 1)
			}

			for i, element := range sl {
				res = append(res, element*(float64(i+1)/sum))
			}

			var total float64 = 0
			for _, element := range res {
				total += element
			}
			result = append(result, total)
		}
	}
	return result
}

// EMA calculates the Exponential Moving Average for the provided period
func EMA(values []float64, period int) []float64 {
	sma := SMA(values, period)
	var result []float64
	var multiplier = (2.0 / (float64(period) + 1.0))

	result = append(result, sma[0])
	for i := (len(values) - len(sma)) + 1; i < len(values); i++ {
		lastVal := result[len(result)-1]
		ema := (values[i]-lastVal)*multiplier + lastVal
		result = append(result, ema)
	}
	return result
}

// DEMA calculates the Double Exponential Moving Average for the provided period
func DEMA(values []float64, period int) []float64 {
	var result []float64
	ema := EMA(values, period)
	emaAgain := EMA(ema, period)

	var emaDouble []float64
	for _, element := range ema {
		emaDouble = append(emaDouble, (2.0 * element))
	}
	offset := len(emaDouble) - len(emaAgain)
	for index, element := range emaAgain {
		result = append(result, (emaDouble[index+offset] - element))
	}
	return result
}

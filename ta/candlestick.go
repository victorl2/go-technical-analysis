package ta

//Candlestick are used to determine price movement and calculate indicators
type Candlestick struct {
	Open, High, Low, Close, Volume float64
}

func NewCandlestick(open, high, low, close, volume float64) Candlestick {
	return Candlestick{open, high, low, close, volume}
}

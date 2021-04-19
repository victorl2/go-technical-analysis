package ta

//Candlestick are used to determine price movement and calculate indicators
type Candlestick struct {
	Timestamp                      string
	Open, High, Low, Close, Volume float64
}

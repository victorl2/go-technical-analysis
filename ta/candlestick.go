package ta

type Value float64

func (p Value) ClosePrice() float64 {
	return float64(p)
}

//ClosePrice represents the close price of a candle
type ClosePrice interface {
	Close() float64
}

//OHLC - Represents a datapoint in candle format that contain Open,High, Low and Close prices
type OHLC interface {
	ClosePrice
	Open() float64
	High() float64
	Low() float64
}

//OHLCV - Represents a datapoint in candle format that contain Open,High, Low, Close prices and Volume data
type OHLCV interface {
	OHLC
	Volume() float64
}

//Candlestick are used to determine price movement and calculate indicators
type Candlestick struct {
	open, high, low, close, volume float64
}

func NewCandlestick(open, high, low, close, volume float64) OHLCV {
	return &Candlestick{open, high, low, close, volume}
}

func (candle *Candlestick) Open() float64 {
	return candle.open
}

func (candle *Candlestick) High() float64 {
	return candle.high
}

func (candle *Candlestick) Low() float64 {
	return candle.low
}

func (candle *Candlestick) Close() float64 {
	return candle.close
}

func (candle *Candlestick) Volume() float64 {
	return candle.volume
}

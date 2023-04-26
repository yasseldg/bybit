package constants

type Symbol string

const (
	Symbol_BTCUSDT   = Symbol("BTCUSDT")
	Symbol_ETHUSDT   = Symbol("ETHUSDT")
	Symbol_XRPUSDT   = Symbol("XRPUSDT")
	Symbol_ADAUSDT   = Symbol("ADAUSDT")
	Symbol_LTCUSDT   = Symbol("LTCUSDT")
	Symbol_SOLUSDT   = Symbol("SOLUSDT")
	Symbol_MATICUSDT = Symbol("MATICUSDT")
	Symbol_DOGEUSDT  = Symbol("DOGEUSDT")
	Symbol_BNBUSDT   = Symbol("BNBUSDT")
	Symbol_AVAXUSDT  = Symbol("AVAXUSDT")

	Symbol_BTCUSD  = Symbol("BTCUSD")
	Symbol_ETHUSD  = Symbol("ETHUSD")
	Symbol_XRPUSD  = Symbol("XRPUSD")
	Symbol_ADAUSD  = Symbol("ADAUSD")
	Symbol_LTCUSD  = Symbol("LTCUSD")
	Symbol_AVAXUSD = Symbol("AVAXUSD")
)

type Interval string

const (
	Interval_1m  = Interval("1")
	Interval_5m  = Interval("5")
	Interval_15m = Interval("15")
	Interval_30m = Interval("30")
	Interval_1h  = Interval("60")
	Interval_2h  = Interval("120")
	Interval_4h  = Interval("240")
	Interval_1d  = Interval("D")
)

const (
	/*
	  http headers
	*/
	ContentType      = "Content-Type"
	AccessKey        = "ACCESS-KEY"
	AccessSign       = "ACCESS-SIGN"
	AccessTimestamp  = "ACCESS-TIMESTAMP"
	AccessPassphrase = "ACCESS-PASSPHRASE"
	ApplicationJson  = "application/json"

	EN_US  = "en_US"
	ZH_CN  = "zh_CN"
	LOCALE = "locale="

	/*
	  http methods
	*/
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
)

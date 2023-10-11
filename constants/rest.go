package constants

const (
	Env_DEFAULT = "Bybit"
)

type Recourse string

const (
	Recourse_Market = Recourse("market")
)

type EndPoint string

const (
	// Market
	EndPoint_InstrumentsInfo = EndPoint("instruments-info")
)

type Category string

// linear, inverse, option, spot
const (
	Category_Linear  = Category("linear")
	Category_Inverse = Category("inverse")
	Category_Option  = Category("option")
	Category_Spot    = Category("spot")
)

type InstrumentStatus string

const (
	InstrumentStatus_PreLaunch  = InstrumentStatus("PreLaunch")
	InstrumentStatus_Trading    = InstrumentStatus("Trading")
	InstrumentStatus_Settling   = InstrumentStatus("Settling")
	InstrumentStatus_Delivering = InstrumentStatus("Delivering")
	InstrumentStatus_Closed     = InstrumentStatus("Closed")
)

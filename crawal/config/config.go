package config

const (
	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"

	ParseCarDetail = "ParseCarDetail"
	ParseCarList   = "ParseCarList"
	ParseCarModel  = "ParseCarModel"

	NilParser = "NilParser"

	// ElasticSearch
	//ElasticIndex = "car_profile"
	ElasticIndex = "dating_profile"
	ItemSaverRpc = "ItemSaverService.Save"

	// Rate limiting
	Qps = 2
)

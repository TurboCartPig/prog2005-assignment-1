package exchange

const (
	// Root URL for the restcountries.eu api
	RestCountriesRoot = "https://restcountries.eu/rest/v2"
	// Root URL for the exchangeratesapi.io api
	ExchangeRatesAPIRoot = "https://api.exchangeratesapi.io"
)

type Countries = []Country

type Country struct {
	Name           string            `json:"name"`
	TopLevelDomain []string          `json:"topLevelDomain"`
	Alpha2Code     string            `json:"alpha2Code"`
	Alpha3Code     string            `json:"alpha3Code"`
	CallingCodes   []string          `json:"callingCodes"`
	Capital        string            `json:"capital"`
	AltSpellings   []string          `json:"altSpellings"`
	Region         string            `json:"region"`
	Subregion      string            `json:"subregion"`
	Popultaion     int               `json:"population"`
	LatLng         []float32         `json:"latlng"`
	Demonym        string            `json:"demonym"`
	Area           float32           `json:"area"`
	Gini           float32           `json:"gini"`
	Timezones      []string          `json:"timezones"`
	Borders        []string          `json:"borders"`
	NativeName     string            `json:"nativeName"`
	NumericCode    string            `json:"numericCode"`
	Currencies     []Currency        `json:"currencies"`
	Languages      []Language        `json:"languages"`
	Translations   map[string]string `json:"translations"`
	Flag           string            `json:"flag"`
	RegionalBlocks []ReginalBlock    `json:"regionalBlocks"`
	Cioc           string            `json:"cioc"`
}

type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Language struct {
	ISO639_1   string `json:"iso639_1"`
	ISO639_2   string `json:"iso639_2"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

type ReginalBlock struct {
	Acronym       string   `json:"acronym"`
	Name          string   `json:"name"`
	OtherAcronyms []string `json:"otherAcronyms"`
	OtherNames    []string `json:"otherNames"`
}

// Describes the rates for exchange between a base currency and a list of other currencies in some timeperiod.
type Rates struct {
	// Base currency
	Base string `json:"base"`
	// Starting date for lookup
	StartAt string `json:"start_at"`
	// End date for lookup
	EndAt string `json:"end_at"`
	// Rate for date for currency
	Rates map[string]map[string]float32 `json:"rates"`
}

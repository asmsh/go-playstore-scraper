package urls

// all the related Play store url query parameters
const (
	AppIDQueryParam      = "id"
	LangQueryParam       = "hl"
	CountryQueryParam    = "gl"
	ResultsNumQueryParam = "num"
)

type QueryParams struct {
	UrlType int // The type of the url that contains these query params,
	// 0: for AppUrl
	// 1: for AppsStoreUrl
	AppID      string
	Lang       string
	Country    string
	ResultsNum int
}

// StoreUrlQueryParams is an array containing all the related query params for the store url
var StoreUrlQueryParams = []string{
	LangQueryParam, CountryQueryParam, ResultsNumQueryParam,
}

// AppUrlQueryParams is an array containing all the related query params for the app page url
var AppUrlQueryParams = []string{
	AppIDQueryParam, LangQueryParam, CountryQueryParam,
}

package urls

import "github.com/asmsh/go-playstore-scraper/engine/urls/locales"

// query parameters common to all url types
const (
	langQueryParam    = "hl"
	countryQueryParam = "gl"
)

// query parameters specific to app page urls
const (
	appIDQueryParam = "id"
)

// query parameters common to collection urls
const (
	resultsNumQueryParam = "num"
)

// internal type
type urlType string

const (
	appPageUrlType        urlType = "APP_PAGE_URL"
	appsCollectionUrlType         = "APPS_COLLECTION_URL"
)

// internal type
type queryParams struct {
	UrlType    urlType // The type of the url that contains these query params
	Country    locales.Country
	Lang       locales.Language
	AppID      string
	ResultsNum int
}

// CollectionUrlQueryParams is an array containing all the related query params for a url that represents a collection in the store
var collectionUrlQueryParams = []string{
	langQueryParam, countryQueryParam, resultsNumQueryParam,
}

// AppUrlQueryParams is an array containing all the related query params for a url that represents an app page
var appUrlQueryParams = []string{
	langQueryParam, countryQueryParam, appIDQueryParam,
}

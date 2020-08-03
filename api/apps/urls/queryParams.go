package urls

import (
	"github.com/asmsh/go-playstore-scraper/locales"
)

// query parameters common to all url types
const (
	langQueryKey    = "hl"
	countryQueryKey = "gl"
)

// query parameters specific to app page urls
const (
	appIdQueryKey = "id"
)

// internal type
type urlType uint8

const (
	appPageUrlType urlType = iota
)

// internal type
type queryParams struct {
	urlType urlType // The type of the url that contains these query params
	country locales.Country
	lang    locales.Language
	appId   string
}

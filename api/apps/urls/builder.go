package urls

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asmsh/go-playstore-scraper/locales"
)

const (
	// appPageBaseUrl is the base part of the url that displays the details of a single app
	appPageBaseUrl = "https://play.google.com/store/apps/details?"
)

var (
	// ignoring the returned error, because an error will not happen
	stdAppUrlBase, _ = url.Parse(appPageBaseUrl)
)

// AppUrl holds the information needed to build the url of a specific app
// in a specific language and country.
type AppUrl struct {
	appId    string
	country  locales.Country
	language locales.Language
}

// Getters for AppUrl
func (appURL AppUrl) AppId() string {
	return appURL.appId
}
func (appURL AppUrl) Country() locales.Country {
	return appURL.country
}
func (appURL AppUrl) Language() locales.Language {
	return appURL.language
}

// NewAppUrl creates a new instance of AppUrl of an app with the provided parameters,
// if the provided parameters is valid, and return it, or return an error otherwise.
func NewAppUrl(appId string, country locales.Country, language locales.Language) (*AppUrl, error) {
	validId, e := validateAppID(appId)
	if e != nil {
		return nil, e
	}
	if !locales.IsValidCountry(country) {
		return nil, fmt.Errorf("invalid or unsupported country code '%s'", country)
	}
	if !locales.IsValidLanguage(language) {
		return nil, fmt.Errorf("invalid or unsupported language code '%s'", language)
	}

	return &AppUrl{appId: validId, country: country, language: language}, nil
}

// NewAppUrlFrom creates a new AppUrl instance from the given url, if it's a valid
// app url with valid query, and return it, or return an error otherwise.
func NewAppUrlFrom(urlStr string) (*AppUrl, error) {
	if len(urlStr) == 0 {
		return nil, fmt.Errorf("the url shouldn't be empty")
	}

	appUrl, e := url.Parse(urlStr)
	if e != nil {
		return nil, e
	}

	if appUrl.Host != stdAppUrlBase.Host || !strings.HasPrefix(appUrl.Path, stdAppUrlBase.Path) {
		return nil, fmt.Errorf("the provided url isn't a valid app url, found '%s'", urlStr)
	}

	qParams, e := readQueryParams(appUrl.Query(), appPageUrlType)
	if e != nil {
		return nil, fmt.Errorf("the url isn't a valid app page url: '%s'", e.Error())
	}

	return NewAppUrl(qParams.appId, qParams.country, qParams.lang)
}

// String returns the app url as a string
func (appURL AppUrl) String() string {
	urlBuild := strings.Builder{}
	urlBuild.WriteString(appPageBaseUrl)

	urlBuild.WriteString(appIdQueryKey)
	urlBuild.WriteString("=")
	urlBuild.WriteString(appURL.appId)

	if len(appURL.country) != 0 {
		urlBuild.WriteString("&")
		urlBuild.WriteString(countryQueryKey)
		urlBuild.WriteString("=")
		urlBuild.WriteString(string(appURL.country))
	}

	if len(appURL.language) != 0 {
		urlBuild.WriteString("&")
		urlBuild.WriteString(langQueryKey)
		urlBuild.WriteString("=")
		urlBuild.WriteString(string(appURL.language))
	}

	return urlBuild.String()
}

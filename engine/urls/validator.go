package urls

import (
	"fmt"
	"github.com/asmsh/go-playstore-scraper/engine/urls/categories"
	"github.com/asmsh/go-playstore-scraper/engine/urls/locales"
	"strconv"
	"strings"
	"unicode"
)

// these consts to handle dropping the 'https' protocol or dropping the protocol field at all from the url
const (
	appPageBaseUrl        = "https://play.google.com/store/apps/details?"
	appPageBaseUrlHTTP    = "http://play.google.com/store/apps/details?"
	appPageBaseUrlNoProto = "play.google.com/store/apps/details?"

	appsStoreBaseUrl        = "https://play.google.com/store/apps"
	appsStoreBaseUrlHTTP    = "http://play.google.com/store/apps"
	appsStoreBaseUrlNoProto = "play.google.com/store/apps"
)

//this file will contain functions that validate the correctness of the arguments of the UrlBuilder and other things across the engine

// ValidateAppID checks the entered appID if it's in a valid appID form,
// and removing any found preceding whitespace, then return a valid and ready to use appID,
// and if the entered appID's from isn't correct, it return an error.
func ValidateAppID(id string) (string, error) {
	if len(id) == 0 {
		return "", fmt.Errorf("the app id shouldn't be empty")
	}

	fields := strings.Fields(id)
	if len(fields) != 1 {
		return "", fmt.Errorf("the app id shouldn't have spaces, found '%s'", id)
	}

	possibleID := fields[0]
	parts := strings.Split(possibleID, ".")

	// The app identifier must not begin with an uppercase letter or number
	// Each part must start with a letter and should contain only letters and numbers.
	for idx, part := range parts {
		if len(part) == 0 {
			return "", fmt.Errorf("the app id is invalid, each app id part mustn't be empty, found '%s'", id)
		}

		if idx == 0 && (unicode.IsDigit(rune(part[0])) || unicode.IsUpper(rune(part[0]))) {
			return "", fmt.Errorf("the app id shouldn't begin with an uppercase letter or number, found '%s'", id)
		}

		for i, c := range part {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				if i == 0 && unicode.IsDigit(c) {
					return "", fmt.Errorf("the app id is invalid, each app id part must start with a letter, found '%s'", id)
				}
			} else {
				return "", fmt.Errorf("the app id should contain only letters and numbers, found '%s'", id)
			}
		}
	}

	return possibleID, nil
}

// ValidateAppPageURL checks if the given url is a valid app page url with valid query parameters,
// and return an instance of AppUrl, or a non-nil error otherwise
func ValidateAppPageURL(url string) (*AppUrl, error) {
	if len(url) == 0 {
		return nil, fmt.Errorf("the url shouldn't be empty")
	}

	if !strings.HasPrefix(url, appPageBaseUrl) && !strings.HasPrefix(url, appPageBaseUrlHTTP) &&
		!strings.HasPrefix(url, appPageBaseUrlNoProto) {
		return nil, fmt.Errorf("the url isn't a valid app page url, found '%s'", url)
	}

	fields := strings.Fields(url)
	if len(fields) != 1 {
		return nil, fmt.Errorf("the url shouldn't have spaces, found '%s'", url)
	}

	possibleURL := fields[0]

	qParams, e := validateQueryParams(possibleURL, 0)
	if e != nil {
		return nil, e
	}

	return NewAppUrl(qParams.AppID, qParams.Country, qParams.Lang)
}

// ValidateAppsStoreURL checks if the given url is a valid store url with valid query parameters,
// and return error otherwise
func ValidateAppsStoreURL(url string) (*QueryParams, error) {
	if len(url) == 0 {
		return nil, fmt.Errorf("the url of the store can't be empty\n")
	}

	if !strings.HasPrefix(url, appsStoreBaseUrl) && !strings.HasPrefix(url, appsStoreBaseUrlHTTP) &&
		!strings.HasPrefix(url, appsStoreBaseUrlNoProto) {
		return nil, fmt.Errorf("the given url isn't a valid play store url\n")
	}

	qParams, e := validateQueryParams(url, 1)
	if e != nil {
		return nil, e
	}

	return qParams, nil
}

// this is a helper function that checks for the queries in the url and sees if they are valid
// mode = 0 : checks for App URL
// mode = 1 : checks for Store URL
func validateQueryParams(url string, mode int) (*QueryParams, error) {
	var qParams = new(QueryParams)

	fi := strings.Index(url, "?")
	if fi == -1 {
		return nil, fmt.Errorf("couldn't found the query string")
	}
	if len(url) <= fi+1 {
		return nil, fmt.Errorf("missing the params of the query string")
	}

	queryString := url[fi+1:]
	queryParts := strings.Split(queryString, "&")

	for _, qp := range queryParts {
		queryContent := strings.Split(qp, "=")
		key, val := queryContent[0], queryContent[1]

		if len(queryContent) != 2 {
			return nil, fmt.Errorf("the query: '%s' in the given url is invalid", qp)
		}
		if len(key) == 0 {
			return nil, fmt.Errorf("the key of the query param shouldn't be empty")
		}
		if len(val) == 0 {
			return nil, fmt.Errorf("the value of the query param shouldn't be empty")
		}

		if mode == 0 {
			qParams.UrlType = 0
			if !aContainS(AppUrlQueryParams, key) {
				return nil, fmt.Errorf("the query: '%s' is unsupported for the given mode: '%d'", key, mode)
			}
		} else if mode == 1 {
			qParams.UrlType = 1
			if !aContainS(StoreUrlQueryParams, key) {
				return nil, fmt.Errorf("the query: '%s' is unsupported for the given mode: '%d'", key, mode)
			}
		} else {
			return nil, fmt.Errorf("the given mode: '%d' is unsupported", mode)
		}

		switch key {
		case AppIDQueryParam:
			if mode != 0 {
				continue
			}
			if vID, e := ValidateAppID(val); e != nil {
				return nil, fmt.Errorf("the app id in the url is invalid, found id: '%s', error: %s", val, e.Error())
			} else {
				qParams.AppID = vID
			}
		case LangQueryParam:
			if !IsValidLanguage(locales.Language(val)) {
				return nil, fmt.Errorf("the languane code in the url is invalid or unsupported, found: %s", val)
			} else {
				qParams.Lang = val
			}
		case CountryQueryParam:
			if !IsValidCountry(locales.Country(val)) {
				return nil, fmt.Errorf("the country code in the url is invalid or unsupported, found: '%s'", val)
			} else {
				qParams.Country = val
			}
		case ResultsNumQueryParam:
			if mode != 1 {
				continue
			}
			if rs, e := strconv.Atoi(val); e != nil {
				return nil, fmt.Errorf("the number of results in the url isn't a valid number, found: '%s'", val)
			} else {
				qParams.ResultsNum = rs
			}
		}
	}

	return qParams, nil
}

func IsCategory(cat categories.Category) bool {
	return IsGameCategory(cat) || IsAppCategory(cat)
}

func IsGameCategory(cat categories.Category) bool {
	for _, v := range categories.GamesCategories {
		if v == cat {
			return true
		}
	}
	return false
}

func IsAppCategory(cat categories.Category) bool {
	for _, v := range categories.AppsCategories {
		if v == cat {
			return true
		}
	}
	return false
}

func IsCollection(col categories.Collection) bool {
	for _, v := range categories.Collections {
		if v == col {
			return true
		}
	}
	return false
}

func IsValidCountry(c locales.Country) bool {
	for _, v := range locales.Countries {
		if v == c {
			return true
		}
	}
	return false
}

func IsValidLanguage(l locales.Language) bool {
	for _, v := range locales.Languages {
		if v == l {
			return true
		}
	}
	return false
}

func IsValidResultNum(n int) bool {
	return n >= 0 && n <= 120
}

func ValidateResultNum(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("the result number shouldn't be less than zero")
	} else if n > 120 {
		return 120, nil
	} else {
		return n, nil
	}
}

// aContainS checks if array 'a' contains string 's'
func aContainS(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

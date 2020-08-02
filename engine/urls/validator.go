package urls

import (
	"fmt"
	"github.com/asmsh/go-playstore-scraper/engine/urls/categories"
	"github.com/asmsh/go-playstore-scraper/engine/urls/locales"
	"net/url"
	"strconv"
	"strings"
	"unicode"
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
func ValidateAppPageURL(urlStr string) (*AppUrl, error) {
	if len(urlStr) == 0 {
		return nil, fmt.Errorf("the app page url shouldn't be empty")
	}

	// ignoring the returned error, as this will not happen
	stdURL, _ := url.Parse(AppPageBaseUrl)

	appUrl, e := url.Parse(urlStr)
	if e != nil {
		return nil, fmt.Errorf("the url isn't a valid app page url, couldn't parse with error: '%s'", e.Error())
	}

	if appUrl.Host != stdURL.Host || !strings.HasPrefix(appUrl.Path, stdURL.Path) {
		return nil, fmt.Errorf("the url isn't a valid app page url, found '%s'", urlStr)
	}

	qParams, e := validateQueryParams(appUrl.Query(), appPageUrlType)
	if e != nil {
		return nil, fmt.Errorf("the url isn't a valid app page url: '%s'", e.Error())
	}

	return NewAppUrl(qParams.AppID, qParams.Country, qParams.Lang)
}

// ValidateAppsStoreURL checks if the given url is a valid url that represent some collection for some app category,
// and has valid query parameters, then return a representation of that url as AppsStoreUrl and nil error,
// or return non-nil error otherwise
func ValidateAppsStoreURL(urlStr string) (*AppsStoreUrl, error) {
	if len(urlStr) == 0 {
		return nil, fmt.Errorf("the apps store url shouldn't be empty")
	}

	// ignoring the returned error, as this will not happen
	stdURL, _ := url.Parse(AppsStoreBaseUrl)

	colUrl, e := url.Parse(urlStr)
	if e != nil {
		return nil, fmt.Errorf("the given url isn't a valid apps store url, couldn't parse with error: %s", e.Error())
	}

	if colUrl.Host != stdURL.Host || !strings.HasPrefix(colUrl.Path, stdURL.Path) {
		return nil, fmt.Errorf("the url isn't a valid apps store url, found '%s'", urlStr)
	}

	paths := strings.Split(colUrl.Path, "/")

	var cat categories.Category
	if i := findSInA(paths, categoryQualifier); i == -1 {
		cat = categories.NoCategory
	} else if i == len(paths)-1 {
		return nil, fmt.Errorf("the url isn't a valid apps store url, missing required fields: category name")
	} else {
		cat = categories.Category(paths[i+1])
	}

	var col categories.Collection
	if i := findSInA(paths, collectionQualifier); i == -1 {
		return nil, fmt.Errorf("the url isn't a valid apps store url, missing required fields: collection")
	} else if i == len(paths)-1 {
		return nil, fmt.Errorf("the url isn't a valid apps store url, missing required fields: collection name")
	} else {
		col = categories.Collection(paths[i+1])
	}

	qParams, e := validateQueryParams(colUrl.Query(), appsCollectionUrlType)
	if e != nil {
		return nil, fmt.Errorf("the url isn't a valid apps store url: %s", e.Error())
	}

	return NewAppsStoreUrl(cat, col, qParams.Country, qParams.Lang, qParams.ResultsNum)
}

// this is a helper function that checks for the queries in the url and sees if they are valid
func validateQueryParams(queries url.Values, ut urlType) (*queryParams, error) {
	var qParams = new(queryParams)

	if !isSupportedUrlType(ut) {
		return nil, fmt.Errorf("unsupported urlType passed: %s", ut)
	}

	qParams.UrlType = ut

	for key, vals := range queries {
		switch {
		case key == countryQueryParam:
			if len(vals) != 1 {
				return nil, fmt.Errorf("unsupported query parameters found for key: %s", key)
			}

			if !IsValidCountry(locales.Country(vals[0])) {
				return nil, fmt.Errorf("invalid or unsupported country code '%s'", vals[0])
			} else {
				qParams.Country = locales.Country(vals[0])
			}
		case key == langQueryParam:
			if len(vals) != 1 {
				return nil, fmt.Errorf("unsupported query parameters found for key: %s", key)
			}

			if !IsValidLanguage(locales.Language(vals[0])) {
				return nil, fmt.Errorf("invalid or unsupported language code '%s'", vals[0])
			} else {
				qParams.Lang = locales.Language(vals[0])
			}
		case key == appIDQueryParam && ut == appPageUrlType:
			if len(vals) != 1 {
				return nil, fmt.Errorf("unsupported query parameters found for key: %s", key)
			}

			if id, e := ValidateAppID(vals[0]); e != nil {
				return nil, e
			} else {
				qParams.AppID = id
			}
		case key == resultsNumQueryParam && ut == appsCollectionUrlType:
			if len(vals) != 1 {
				return nil, fmt.Errorf("unsupported query parameters found for key: %s", key)
			}

			n, e := strconv.Atoi(vals[0])
			if e != nil {
				return nil, fmt.Errorf("invalid number of results: %s", e.Error())
			}

			if rn, e := ValidateResultNum(n); e != nil {
				return nil, fmt.Errorf("invalid or unsupported number of results: %d", n)
			} else {
				qParams.ResultsNum = rn
			}
		}
	}

	// the app page url has to have at least one query parameter(app id)
	if ut == appPageUrlType && len(qParams.AppID) == 0 {
		return nil, fmt.Errorf("missing required query parameters: id")
	}

	return qParams, nil
}

//
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

// findSInA return the first index of string 's' in array 'a' or -1 if it's not found
func findSInA(a []string, s string) int {
	for i, v := range a {
		if v == s {
			return i
		}
	}
	return -1
}

func isSupportedUrlType(ut urlType) bool {
	return ut == appPageUrlType || ut == appsCollectionUrlType
}

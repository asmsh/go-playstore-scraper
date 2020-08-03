package urls

import (
	"fmt"
	"net/url"
	"strings"
	"unicode"

	"github.com/asmsh/go-playstore-scraper/locales"
)

// this file will contain functions that validate the correctness of the arguments of
// the UrlBuilder and other things across the internal.

// validateAppID checks the entered appId if it's in a valid appId form,
// and removing any found preceding whitespace, then return a valid and ready to use appId,
// and if the entered appId's from isn't correct, it return an error.
func validateAppID(id string) (string, error) {
	if len(id) == 0 {
		return "", fmt.Errorf("the app id shouldn't be empty")
	}

	fields := strings.Fields(id)
	if len(fields) != 1 {
		return "", fmt.Errorf("the app id shouldn't have spaces, found '%s'", id)
	}

	updatedId := fields[0]
	parts := strings.Split(updatedId, ".")

	// The app identifier must not begin with an uppercase letter or number
	// Each part must start with a letter and should contain only letters and numbers.
	for idx, part := range parts {
		if len(part) == 0 {
			return "", fmt.Errorf("each app id part mustn't be empty, found '%s'", updatedId)
		}

		if idx == 0 && (unicode.IsDigit(rune(part[0])) || unicode.IsUpper(rune(part[0]))) {
			return "", fmt.Errorf("the app id shouldn't begin with an uppercase letter or number, found '%s'", updatedId)
		}

		for i, c := range part {
			isLetter := unicode.IsLetter(c)
			isDigit := unicode.IsDigit(c)
			if isLetter || isDigit {
				if i == 0 && isDigit {
					return "", fmt.Errorf("each app id part must start with a letter, found '%s'", updatedId)
				}
			} else {
				return "", fmt.Errorf("the app id should contain only letters and numbers, found '%s'", updatedId)
			}
		}
	}

	return updatedId, nil
}

// this is a helper function that checks for the queries in the url and sees if they are valid
func readQueryParams(queryVals url.Values, ut urlType) (*queryParams, error) {
	qParams := new(queryParams)
	qParams.urlType = ut

	for key, vals := range queryVals {
		switch key {
		case countryQueryKey:
			if len(vals) != 1 {
				return nil, fmt.Errorf("unsupported query value for key '%s'", key)
			}

			if !locales.IsValidCountry(locales.Country(vals[0])) {
				return nil, fmt.Errorf("invalid or unsupported country code '%s'", vals[0])
			} else {
				qParams.country = locales.Country(vals[0])
			}
		case langQueryKey:
			if len(vals) != 1 {
				return nil, fmt.Errorf("unsupported query value for key '%s'", key)
			}

			if !locales.IsValidLanguage(locales.Language(vals[0])) {
				return nil, fmt.Errorf("invalid or unsupported language code '%s'", vals[0])
			} else {
				qParams.lang = locales.Language(vals[0])
			}
		case appIdQueryKey:
			if ut == appPageUrlType {
				if len(vals) != 1 {
					return nil, fmt.Errorf("unsupported query value for key '%s'", key)
				}

				if id, e := validateAppID(vals[0]); e != nil {
					return nil, e
				} else {
					qParams.appId = id
				}
			}
		}
	}

	// the app page url has to have at least one query parameter(app id)
	if ut == appPageUrlType && len(qParams.appId) == 0 {
		return nil, fmt.Errorf("missing required query parameters: id")
	}

	return qParams, nil
}

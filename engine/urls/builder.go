package urls

import (
	"fmt"
	"github.com/asmsh/go-playstore-scraper/engine/urls/categories"
	"github.com/asmsh/go-playstore-scraper/engine/urls/locales"
	"strconv"
)

const (
	// AppsStoreBaseUrl is the Play Store apps home page url
	AppsStoreBaseUrl = "https://play.google.com/store/apps"
	// AppPageBaseUrl is the base part of the url that displays the details of a single app
	AppPageBaseUrl = "https://play.google.com/store/apps/details?"
)

// internal constants for building and verifying the urls
const (
	categoryQualifier   = "category"
	categoryPath        = "/" + categoryQualifier + "/"
	collectionQualifier = "collection"
	collectionPath      = "/" + collectionQualifier + "/"
)

// AppUrl holds the information needed to build an app page url
type AppUrl struct {
	appID    string
	country  locales.Country
	language locales.Language
}

// Getters for AppUrl
func (appURL AppUrl) AppID() string {
	return appURL.appID
}
func (appURL AppUrl) Country() locales.Country {
	return appURL.country
}
func (appURL AppUrl) Language() locales.Language {
	return appURL.language
}

// NewAppUrl returns a new instance of AppUrl
// the order of the optional parameters matters,
// so, if only one optional parameter is passed it will be used as the country code,
// if two optional parameters are passed they will be used according to the following order: country, language,
// and, if more than two optional parameters are passed, the func will return an error.
// the country is expected to be of type Country or type string.
// the language is expected to be of type Language or type string.
func NewAppUrl(appID string, restParams ...interface{}) (*AppUrl, error) {
	const errPrefix = "couldn't build the app url:"

	validID, e := ValidateAppID(appID)
	if e != nil {
		return nil, fmt.Errorf("%s %s", errPrefix, e.Error())
	}

	switch len(restParams) {
	case 0:
		return &AppUrl{appID: validID, country: locales.CountryNone, language: locales.LanguageNone}, nil
	case 1:
		param1 := restParams[0]
		var country locales.Country

		switch param1.(type) {
		case string:
			country = locales.Country(param1.(string))
		case locales.Country:
			country = param1.(locales.Country)
		default:
			return nil, fmt.Errorf("%s the country should be of type 'Country' or 'string', found '%T'", errPrefix, param1)
		}

		if !IsValidCountry(country) {
			return nil, fmt.Errorf("%s invalid or unsupported country code '%s'", errPrefix, country)
		}

		return &AppUrl{appID: validID, country: country, language: locales.LanguageNone}, nil
	case 2:
		param1 := restParams[0]
		param2 := restParams[1]

		var country locales.Country
		var lang locales.Language

		switch param1.(type) {
		case string:
			country = locales.Country(param1.(string))
		case locales.Country:
			country = param1.(locales.Country)
		default:
			return nil, fmt.Errorf("%s the country should be of type 'Country' or 'string', found '%T'", errPrefix, param1)
		}

		switch param2.(type) {
		case string:
			lang = locales.Language(param2.(string))
		case locales.Language:
			lang = param2.(locales.Language)
		default:
			return nil, fmt.Errorf("%s the language should be of type 'Language' or 'string', found '%T'", errPrefix, param2)
		}

		if !IsValidCountry(country) {
			return nil, fmt.Errorf("%s invalid or unsupported country code '%s'", errPrefix, country)
		}
		if !IsValidLanguage(lang) {
			return nil, fmt.Errorf("%s invalid or unsupported language code '%s'", errPrefix, lang)
		}

		return &AppUrl{appID: validID, country: country, language: lang}, nil
	default:
		return nil, fmt.Errorf("%s invalid number of arguments passed, expected: 1-3, found: %d", errPrefix, len(restParams))
	}
}

// String returns the app url as a string
func (appURL AppUrl) String() string {
	url := AppPageBaseUrl

	url += appIDQueryParam + "=" + appURL.appID

	if len(appURL.country) != 0 {
		url += "&" + countryQueryParam + "=" + string(appURL.country)
	}

	if len(appURL.language) != 0 {
		url += "&" + langQueryParam + "=" + string(appURL.language)
	}

	return url
}

// AppsStoreUrl holds the required info that represent the url of some collection for some app category,
type AppsStoreUrl struct {
	category   categories.Category
	collection categories.Collection
	country    locales.Country  // the country of the result page
	language   locales.Language // the language of the result page
	resultsNum int              // the number of result per page, max. 120
}

// Getters for AppsStoreUrl
func (storeUrl AppsStoreUrl) Category() categories.Category {
	return storeUrl.category
}
func (storeUrl AppsStoreUrl) Collection() categories.Collection {
	return storeUrl.collection
}
func (storeUrl AppsStoreUrl) Country() locales.Country {
	return storeUrl.country
}
func (storeUrl AppsStoreUrl) Language() locales.Language {
	return storeUrl.language
}
func (storeUrl AppsStoreUrl) ResultsNum() int {
	return storeUrl.resultsNum
}

// NewAppsStoreUrl returns a new instance of AppsStoreUrl
//
// the order of the optional parameters matters,
// so, if only one optional parameter is passed it will be used as the country code,
// if two optional parameters are passed they will be used according to the following order: country, language,
// if three optional parameters are passed they will be used according to the following order: country, language, resultNum,
// and, if more than three optional parameters are passed, the func will return an error.
//
// the type of the optional parameters passed matters as well,
// we expect the country to be of type Country or string, language of type Language or string, and resultNum of type int.
//
// resultNum value should be between 0-120,
// a value of '0' means that the resultNum query of the url will be left out,
// and the website will decide the number of returned results.
func NewAppsStoreUrl(cat categories.Category, col categories.Collection, restParams ...interface{}) (*AppsStoreUrl, error) {
	const errPrefix = "couldn't build the store url:"

	if !(IsGameCategory(cat) || IsAppCategory(cat)) {
		return nil, fmt.Errorf("%s invalid or unsupported category '%s'", errPrefix, cat)
	}
	if !IsCollection(col) {
		return nil, fmt.Errorf("%s invalid or unsupported collection '%s'", errPrefix, col)
	}

	switch len(restParams) {
	case 0:
		return &AppsStoreUrl{category: cat, collection: col, country: locales.CountryNone, language: locales.LanguageNone, resultsNum: 0}, nil
	case 1:
		param1 := restParams[0]

		var country locales.Country

		switch param1.(type) {
		case string:
			country = locales.Country(param1.(string))
		case locales.Country:
			country = param1.(locales.Country)
		default:
			return nil, fmt.Errorf("%s the country should be of type 'Country' or 'string', found '%T'", errPrefix, param1)
		}

		if !IsValidCountry(country) {
			return nil, fmt.Errorf("%s invalid or unsupported country code '%s'", errPrefix, country)
		}

		return &AppsStoreUrl{category: cat, collection: col, country: country, language: locales.LanguageNone, resultsNum: 0}, nil
	case 2:
		param1 := restParams[0]
		param2 := restParams[1]

		var country locales.Country
		var lang locales.Language

		switch param1.(type) {
		case string:
			country = locales.Country(param1.(string))
		case locales.Country:
			country = param1.(locales.Country)
		default:
			return nil, fmt.Errorf("%s the country should be of type 'Country' or 'string', found '%T'", errPrefix, param1)
		}

		switch param2.(type) {
		case string:
			lang = locales.Language(param2.(string))
		case locales.Language:
			lang = param2.(locales.Language)
		default:
			return nil, fmt.Errorf("%s the language should be of type 'Language' or 'string', found '%T'", errPrefix, param2)
		}

		if !IsValidCountry(country) {
			return nil, fmt.Errorf("%s invalid or unsupported country code '%s'", errPrefix, country)
		}
		if !IsValidLanguage(lang) {
			return nil, fmt.Errorf("%s invalid or unsupported language code '%s'", errPrefix, lang)
		}

		return &AppsStoreUrl{category: cat, collection: col, country: country, language: lang, resultsNum: 0}, nil
	case 3:
		param1 := restParams[0]
		param2 := restParams[1]
		param3 := restParams[2]

		var country locales.Country
		var lang locales.Language

		switch param1.(type) {
		case string:
			country = locales.Country(param1.(string))
		case locales.Country:
			country = param1.(locales.Country)
		default:
			return nil, fmt.Errorf("%s the country should be of type 'Country' or 'string', found '%T'", errPrefix, param1)
		}

		switch param2.(type) {
		case string:
			lang = locales.Language(param2.(string))
		case locales.Language:
			lang = param2.(locales.Language)
		default:
			return nil, fmt.Errorf("%s the language should be of type 'Language' or 'string', found '%T'", errPrefix, param2)
		}

		resNum, ok := param3.(int)
		if !ok {
			return nil, fmt.Errorf("%s the result number should be of type 'int', found '%T'", errPrefix, param3)
		}

		if !IsValidCountry(country) {
			return nil, fmt.Errorf("%s invalid or unsupported country code '%s'", errPrefix, country)
		}
		if !IsValidLanguage(lang) {
			return nil, fmt.Errorf("%s invalid or unsupported language code '%s'", errPrefix, lang)
		}
		resNum, e := ValidateResultNum(resNum)
		if e != nil {
			return nil, fmt.Errorf("%s invalid or unsupported result number: '%s'", errPrefix, e.Error())
		}
		/*if !IsValidResultNum(resNum) {
			return nil, fmt.Errorf("%s invalid or unsupported result number '%d'", errPrefix, resNum)
		}*/

		return &AppsStoreUrl{category: cat, collection: col, country: country, language: lang, resultsNum: resNum}, nil
	default:
		return nil, fmt.Errorf("%s invalid number of arguments passed, expected: 2-5, found: %d", errPrefix)
	}
}

// String returns the apps store url as a string
func (storeUrl AppsStoreUrl) String() string {
	url := AppsStoreBaseUrl

	if storeUrl.category != categories.NoCategory {
		url += categoryPath
		url += string(storeUrl.category)
	}

	url += collectionPath
	url += string(storeUrl.collection)

	if storeUrl.country != "" || storeUrl.language != "" || storeUrl.resultsNum != 0 {
		url += "?"
	} else {
		return url
	}

	if len(storeUrl.country) != 0 {
		url += countryQueryParam + "=" + string(storeUrl.country)

		if storeUrl.language != "" || storeUrl.resultsNum != 0 {
			url += "&"
		} else {
			return url
		}
	}

	if len(storeUrl.language) != 0 {
		url += langQueryParam + "=" + string(storeUrl.language)

		if storeUrl.resultsNum != 0 {
			url += "&"
		} else {
			return url
		}
	}

	// we know for sure it't not a zero
	if storeUrl.resultsNum != 0 {
		url += resultsNumQueryParam + "=" + strconv.Itoa(storeUrl.resultsNum)
	}

	return url
}

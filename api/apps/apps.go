package apps

import (
	"github.com/asmsh/go-playstore-scraper/api/apps/fields"
	"github.com/asmsh/go-playstore-scraper/api/apps/urls"
	"github.com/asmsh/go-playstore-scraper/locales"
)

// GetAppById returns an AppInfo struct representing the info of the app with id
// equal to the provided appId. It returns an nil error if all went well, otherwise
// it returns a non-nil error.
// If there are specific fields provided, then only these fields will be populated
// in the returned AppInfo struct.
func GetAppById(appId string, fields ...fields.AppField) (*AppInfo, error) {
	appUrl, e := urls.NewAppUrl(appId, locales.DefCountry, locales.DefLanguage)
	if e != nil {
		return nil, e
	}

	return parseAppPage(appUrl, fields)
}

// GetAppByIdAdv returns an AppInfo struct representing the info of the app with
// id equal to the provided appId, as would be seen when localized to the provided
// country and to the provided language. It returns an nil error if all went well,
// otherwise it returns a non-nil error.
// If there are specific fields provided, then only these fields will be populated
// in the returned AppInfo struct.
func GetAppByIdAdv(appId string, country locales.Country, language locales.Language, fields ...fields.AppField) (*AppInfo, error) {
	appUrl, e := urls.NewAppUrl(appId, country, language)
	if e != nil {
		return nil, e
	}

	return parseAppPage(appUrl, fields)
}

// GetAppByUrl returns an AppInfo struct representing the info of the app at the
// provided url. It returns an nil error if all went well, otherwise it returns
// a non-nil error.
// If there are specific fields provided, then only these fields will be populated
// in the returned AppInfo struct.
func GetAppByUrl(url string, fields ...fields.AppField) (*AppInfo, error) {
	appUrl, e := urls.NewAppUrlFrom(url)
	if e != nil {
		return nil, e
	}

	return parseAppPage(appUrl, fields)
}

// GetAppByUrlAdv returns an AppInfo struct representing the info of the app at the
// provided url, as would be seen when localized to the provided country and to the
// provided language. It returns an nil error if all went well, otherwise it returns
// a non-nil error.
// If there are specific fields provided, then only these fields will be populated
// in the returned AppInfo struct.
func GetAppByUrlAdv(url string, country locales.Country, language locales.Language, fields ...fields.AppField) (*AppInfo, error) {
	appUrl, e := urls.NewAppUrlFrom(url)
	if e != nil {
		return nil, e
	}
	appUrl, e = urls.NewAppUrl(appUrl.AppId(), country, language)
	if e != nil {
		return nil, e
	}

	return parseAppPage(appUrl, fields)
}

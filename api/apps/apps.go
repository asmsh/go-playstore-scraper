package apps

import (
	"github.com/asmsh/go-playstore-scraper/api/apps/fields"
	"github.com/asmsh/go-playstore-scraper/api/apps/urls"
	"github.com/asmsh/go-playstore-scraper/locales"
)

func GetAppById(appId string, fields ...fields.AppField) (*AppInfo, error) {
	appUrl, e := urls.NewAppUrl(appId, locales.DefCountry, locales.DefLanguage)
	if e != nil {
		return nil, e
	}

	return parseApp(appUrl, fields...)
}

func GetAppByIdAdv(appId string, country locales.Country, language locales.Language, fields ...fields.AppField) (*AppInfo, error) {
	appUrl, e := urls.NewAppUrl(appId, country, language)
	if e != nil {
		return nil, e
	}

	return parseApp(appUrl, fields...)
}

func GetAppByUrl(url string, fields ...fields.AppField) (*AppInfo, error) {
	appUrl, e := urls.NewAppUrlFrom(url)
	if e != nil {
		return nil, e
	}

	return parseApp(appUrl, fields...)
}

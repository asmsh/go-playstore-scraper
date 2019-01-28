package apps

import (
	"github.com/asmsh/go-playstore-scraper/engine/core/appPage"
	"github.com/asmsh/go-playstore-scraper/engine/types/appField"
	"github.com/asmsh/go-playstore-scraper/engine/types/fullApp"
	"github.com/asmsh/go-playstore-scraper/engine/urls"
	"github.com/asmsh/go-playstore-scraper/engine/urls/locales"
)

func GetAppByID(appID string, fields ...appField.AppField) (*fullApp.FullApp, error) {
	url, e := urls.NewAppUrl(appID)
	if e != nil {
		return nil, e
	}

	return appPage.ParseApp(url, fields...)
}

func GetAppByIDAdv(appID string, country locales.Country, language locales.Language, fields ...appField.AppField) (*fullApp.FullApp, error) {
	url, e := urls.NewAppUrl(appID, country, language)
	if e != nil {
		return nil, e
	}

	return appPage.ParseApp(url, fields...)
}

func GetAppByUrl(url string, fields ...appField.AppField) (*fullApp.FullApp, error) {
	return appPage.ParseAppByUrl(url, fields...)
}

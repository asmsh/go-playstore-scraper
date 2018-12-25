package api

import (
	"github.com/asmsh/go-playstore-scraper/engine/core/appPage"
	"github.com/asmsh/go-playstore-scraper/engine/restype/fullApp"
	"github.com/asmsh/go-playstore-scraper/engine/urls"
	"github.com/asmsh/go-playstore-scraper/engine/urls/locales"
)

func GetAppByID(appID string) (*fullApp.App, error) {
	url, e := urls.NewAppUrl(appID)
	if e != nil {
		return nil, e
	}

	return appPage.ParseApp(url)
}

func GetAppByIDAdv(appID string, country locales.Country, language locales.Language) (*fullApp.App, error) {
	url, e := urls.NewAppUrl(appID, country, language)
	if e != nil {
		return nil, e
	}

	return appPage.ParseApp(url)
}

func GetAppByUrl(url string) (*fullApp.App, error) {
	return appPage.ParseAppByUrl(url)
}

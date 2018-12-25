package appCard

import "fmt"

const appPageBaseUrl = "https://play.google.com/store/apps/details?id="

type AppCard struct {
	AppName string `json:"appName,omitempty"`
	AppUrl  string `json:"appURL,omitempty"`
	IconUrl string `json:"iconURL,omitempty"`
	DevName string `json:"devName,omitempty"`
	AppID   string `json:"appID,omitempty"`
	Rating  string `json:"rating,omitempty"`
	Price   string `json:"price,omitempty"`
	// this is a short description that's written in the app cards.
	ShortDesc string `json:"shortDesc,omitempty"`
}

func (app *AppCard) SetAppName(appName string) {
	app.AppName = appName
}

func (app *AppCard) SetIconUrl(iconUrl string) {
	app.IconUrl = iconUrl
}

func (app *AppCard) SetDevName(devName string) {
	app.DevName = devName
}

func (app *AppCard) SetAppID(appID string) {
	app.AppID = appID
}

// sets the app page url to the default one(which one play store choose)
func (app *AppCard) SetDefaultAppUrl() error {
	if app.AppID == "" {
		return fmt.Errorf("the app id isn't set to a value yet")
	}
	app.AppUrl = appPageBaseUrl + app.AppID
	return nil
}

func (app *AppCard) SetAppUrl(url string) {
	app.AppUrl = url
}

func (app *AppCard) SetPrice(price string) {
	app.Price = price
}

func (app *AppCard) SetRating(rating string) {
	app.Rating = rating
}

func (app *AppCard) SetShortDesc(shortDesc string) {
	app.ShortDesc = shortDesc
}

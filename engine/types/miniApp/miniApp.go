package miniApp

import "fmt"

const appPageBaseUrl = "https://play.google.com/store/apps/details?id="

type MiniApp struct {
	AppName   string `json:",omitempty"`
	AppUrl    string `json:",omitempty"`
	IconUrl   string `json:",omitempty"`
	DevName   string `json:",omitempty"`
	AppID     string `json:",omitempty"`
	Rating    string `json:",omitempty"`
	Price     string `json:",omitempty"`
	ShortDesc string `json:",omitempty"`
}

func (app *MiniApp) SetAppName(appName string) {
	app.AppName = appName
}

func (app *MiniApp) SetIconUrl(iconUrl string) {
	app.IconUrl = iconUrl
}

func (app *MiniApp) SetDevName(devName string) {
	app.DevName = devName
}

func (app *MiniApp) SetAppID(appID string) {
	app.AppID = appID
}

// sets the app page url to the default one(which one play store choose)
func (app *MiniApp) SetDefaultAppUrl() error {
	if app.AppID == "" {
		return fmt.Errorf("the app id isn't set to a value yet")
	}
	app.AppUrl = appPageBaseUrl + app.AppID
	return nil
}

func (app *MiniApp) SetAppUrl(url string) {
	app.AppUrl = url
}

func (app *MiniApp) SetPrice(price string) {
	app.Price = price
}

func (app *MiniApp) SetRating(rating string) {
	app.Rating = rating
}

func (app *MiniApp) SetShortDesc(shortDesc string) {
	app.ShortDesc = shortDesc
}

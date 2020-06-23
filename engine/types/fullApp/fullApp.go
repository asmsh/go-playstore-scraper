package fullApp

import (
	"encoding/json"
	"fmt"
	"strings"
)

// the base part of the url for an app page
const appPageBaseUrl = "https://play.google.com/store/apps/details?id="

// the attributes that are written in uppercase letters will be exported to the ToJSON
type FullApp struct {
	// this is the id of this app, i.e. the apk name
	AppID  string `json:",omitempty"`
	AppUrl string `json:",omitempty"`
	// image url at the first index is the default(low res) image, and image at the second index is a higher res one.
	IconUrls   []string `json:",omitempty"`
	AppName    string   `json:",omitempty"`
	DevName    string   `json:",omitempty"`
	DevPageUrl string   `json:",omitempty"`
	Category   string   `json:",omitempty"`

	// TODO, handle the sec category(additional family category) extraction.
	//this is a second optional category that some apps and games implement
	//SecCategory string `json:"secCategory,omitempty"`

	// TODO, rename this property to a more related name
	InAppOffering string `json:",omitempty"`

	// TODO,
	//ContainsAds          bool `json:"containsAds"`
	//OffersInAppPurchases bool `json:"offersInAppPurchases"`

	Price string `json:",omitempty"`
	// an array that contains the urls of the app's video trailer(if it provides a video trailer).
	// the trailer's start image's url is at its first position, and the trailer's video's url is at its second position.
	VideoTrailerUrls []string `json:",omitempty"`
	// an array that contains the urls of the screenshots of the app, the array is populated pair by pair,
	// the first element in every pair is the lower resolution screenshot, while the second element is the high resolution one.
	// by that populating pattern, every url at even index(index starts at 0) will be a lower resolution one,
	// and every url at odd index will be a high resolution one.
	ScreenShotsUrls []string `json:",omitempty"`
	Description     string   `json:",omitempty"`
	// the rating of this app got, the aggregate rating
	Rating string `json:",omitempty"`
	// the number of users that rated this app
	RatingCount string `json:",omitempty"`
	// the rating histogram bars from the rating box,
	// the 'key' is the star number, and the 'value' is the number of ratings for each star
	RatingHistogram map[string]string `json:",omitempty"`
	// this holds the what's new section of the app page
	WhatsNew string `json:",omitempty"`

	// TODO
	// the date of the last update for this app
	LastUpdated string `json:",omitempty"`
	Size        string `json:",omitempty"`
	// the number of downloads for this app
	NumOfDownloads string `json:",omitempty"`
	// the version of the current update (app)
	AppVer string `json:",omitempty"`
	// the minimum Android version required for this app to run (install)
	MinAndroidVer string `json:",omitempty"`
	// the age rating for the content of this app, the minimum age required for this app
	AgeRating string `json:",omitempty"`
}

func (app *FullApp) SetAppID(appID string) {
	app.AppID = appID
}

func (app *FullApp) SetAppUrl(appURL string) {
	app.AppUrl = appURL
}

// sets the app page url to the default one(the one play store choose)
func (app *FullApp) SetDefaultAppUrl() error {
	if app.AppID == "" {
		return fmt.Errorf("the app id isn't set to a value yet")
	}
	app.AppUrl = appPageBaseUrl + app.AppID
	return nil
}

func (app *FullApp) SetIconUrls(iconURLs []string) {
	app.IconUrls = iconURLs
}

func (app *FullApp) SetAppName(appName string) {
	app.AppName = appName
}

func (app *FullApp) SetDevPageUrl(url string) {
	app.DevPageUrl = url
}

func (app *FullApp) SetDevName(devName string) {
	app.DevName = devName
}

func (app *FullApp) SetCategory(cat string) {
	app.Category = cat
}

/*
// sets the second category that this app belongs to, if it present
func (app *FullApp) SetSecCategory(secCat string) {
	app.SecCategory = secCat
}
*/

func (app *FullApp) SetInAppOffering(offeringStr string) {
	app.InAppOffering = offeringStr
}

/*
func (app *FullApp) SetContainsAds() {
	app.ContainsAds = true
}

func (app *FullApp) SetOffersInAppPurchases() {
	app.OffersInAppPurchases = true
}
*/

func (app *FullApp) SetAppPrice(price string) {
	app.Price = price
}

func (app *FullApp) SetVideoTrailerUrls(urls []string) {
	app.VideoTrailerUrls = urls
}

func (app *FullApp) SetScreenShotsUrls(urls []string) {
	app.ScreenShotsUrls = urls
}

func (app *FullApp) SetDescription(desc string) {
	app.Description = desc
}

func (app *FullApp) SetRating(rating string) {
	app.Rating = rating
}

func (app *FullApp) SetRatingCount(ratingCount string) {
	app.RatingCount = ratingCount
}

func (app *FullApp) SetRatingHistogram(histogram map[string]string) {
	app.RatingHistogram = histogram
}

func (app *FullApp) SetWhatsNew(wsNew string) {
	app.WhatsNew = wsNew
}

func (app *FullApp) SetLastUpdated(updated string) {
	app.LastUpdated = updated
}

func (app *FullApp) SetAppSize(appSize string) {
	app.Size = appSize
}

func (app *FullApp) SetNumOfDownloads(numOfDls string) {
	app.NumOfDownloads = numOfDls
}

func (app *FullApp) SetAppVer(appVer string) {
	app.AppVer = appVer
}

func (app *FullApp) SetMinAndroidVer(minVer string) {
	app.MinAndroidVer = minVer
}

func (app *FullApp) SetMinAgeRating(minAge string) {
	app.AgeRating = minAge
}

// ToJSON returns the app info as a JSON string
func (app *FullApp) ToJSON() (string, error) {
	ba, e := json.Marshal(app)
	if e != nil {
		return "", e
	}

	return string(ba), nil
}

func (app *FullApp) String() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("AppID: %q\n", app.AppID))
	sb.WriteString(fmt.Sprintf("AppUrl: %q\n", app.AppUrl))
	sb.WriteString(fmt.Sprintf("IconUrls: %q\n", app.IconUrls))
	sb.WriteString(fmt.Sprintf("AppName: %q\n", app.AppName))
	sb.WriteString(fmt.Sprintf("DevName: %q\n", app.DevName))
	sb.WriteString(fmt.Sprintf("DevPageUrl: %q\n", app.DevPageUrl))
	sb.WriteString(fmt.Sprintf("Category: %q\n", app.Category))
	sb.WriteString(fmt.Sprintf("InAppOffering: %q\n", app.InAppOffering))
	sb.WriteString(fmt.Sprintf("Price: %q\n", app.Price))
	sb.WriteString(fmt.Sprintf("VideoTrailerUrls: %q\n", app.VideoTrailerUrls))
	sb.WriteString(fmt.Sprintf("ScreenShotsUrls: %q\n", app.ScreenShotsUrls))
	sb.WriteString(fmt.Sprintf("Description: %q\n", app.Description))
	sb.WriteString(fmt.Sprintf("Rating: %q\n", app.Rating))
	sb.WriteString(fmt.Sprintf("RatingCount: %q\n", app.RatingCount))
	sb.WriteString(fmt.Sprintf("RatingHistogram: %q\n", app.RatingHistogram))
	sb.WriteString(fmt.Sprintf("WhatsNew: %q\n", app.WhatsNew))

	// TODO, update this function after implementing the missing fields

	return sb.String()
}

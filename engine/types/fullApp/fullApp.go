package fullApp

import "fmt"

// the base part of the url for an app page
const appPageBaseUrl = "https://play.google.com/store/apps/details?id="

// the attributes that are written in uppercase letters will be exported to the JSON
type FullApp struct {
	// this is the id of this app, i.e. the apk name
	AppID string `json:",omitempty"`
	// the link of this app on the store
	AppUrl string `json:",omitempty"`
	// the urls for the icon image of this app.
	// image at the first index is the default(low res) image, and image at the second index is a higher res one.
	IconUrls []string `json:",omitempty"`
	// the name of the app
	AppName string `json:",omitempty"`
	// the name of the developer of this app
	DevName string `json:",omitempty"`
	// the url of the developer's page
	DevPageUrl string `json:",omitempty"`
	// the category of this app
	Category string `json:",omitempty"`

	// TODO, handle the sec category(additional family category) extraction.
	//this is a second optional category that some apps and games implement
	//SecCategory string `json:"secCategory,omitempty"`

	// TODO, rename this property to a more related name
	InAppOffering string `json:",omitempty"`

	// TODO,
	//ContainsAds          bool `json:"containsAds"`
	//OffersInAppPurchases bool `json:"offersInAppPurchases"`

	// the price of this app
	Price string `json:",omitempty"`
	// an array that contains the urls of the app's video trailer(if it provides a video trailer).
	// the trailer's start image's url is at its first position, and the trailer's url is at its second position.
	VideoTrailerUrls []string `json:",omitempty"`
	// an array that contains the urls of the screenshots of the app, the array is populated pair by pair,
	// the first element in every pair is the lower resolution screenshot, while the second element is the high resolution one.
	// by that populating pattern, every odd index(index starts at 0) will be a lower resolution one,
	// and every even index will be a high resolution one.
	ScreenShotsUrls []string `json:",omitempty"`
	// the description of this app
	Description string `json:",omitempty"`
	// the rating of this app got, the aggregate rating
	Rating string `json:",omitempty"`
	// the number of users that rated this app
	RatingCount string `json:",omitempty"`
	// the rating histogram bars from the rating box,
	// the 'key' is the star number, and the 'value' is the number of ratings for each star
	RatingHistogram map[string]string `json:",omitempty"`
	// this holds the what's new section of the app page
	WhatsNew string `json:",omitempty"`
	// the date of the last update for this app
	LastUpdated string `json:",omitempty"`
	// the size of the app (this requires some edited browser agent)
	Size string `json:",omitempty"`
	// the number of downloads for this app
	NumOfDownloads string `json:",omitempty"`
	// the version of the current update (app)
	AppVer string `json:",omitempty"`
	// the minimum Android version required for this app to run (install)
	MinAndroidVer string `json:",omitempty"`
	// the age rating for the content of this app, the minimum age required for this app
	AgeRating string `json:",omitempty"`
}

// sets the name of the apk of this app,
func (app *FullApp) SetAppID(appID string) {
	app.AppID = appID
}

// sets the app page url to the given url
func (app *FullApp) SetAppUrl(appURL string) {
	app.AppUrl = appURL
}

// sets the app page url to the default one(which one play store choose)
func (app *FullApp) SetDefaultAppUrl() error {
	if app.AppID == "" {
		return fmt.Errorf("the app id isn't set to a value yet")
	}
	app.AppUrl = appPageBaseUrl + app.AppID
	return nil
}

// sets the url of this app's icon image source
func (app *FullApp) SetIconUrls(iconURLs []string) {
	app.IconUrls = iconURLs
}

// sets the name of the app
func (app *FullApp) SetAppName(appName string) {
	app.AppName = appName
}

func (app *FullApp) SetDevPageUrl(url string) {
	app.DevPageUrl = url
}

// sets the name of the developer of this app
func (app *FullApp) SetDevName(devName string) {
	app.DevName = devName
}

// sets the category that this app belongs to
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
// sets the price of this app
func (app *FullApp) SetAppPrice(price string) {
	app.Price = price
}

func (app *FullApp) SetVideoTrailerUrls(urls []string) {
	app.VideoTrailerUrls = urls
}

// sets the urls of the screenshots for an app
func (app *FullApp) SetScreenShotsUrls(urls []string) {
	app.ScreenShotsUrls = urls
}

// sets the app's description
func (app *FullApp) SetDescription(desc string) {
	app.Description = desc
}

// sets the rating that this app got from the user on the store
func (app *FullApp) SetRating(rating string) {
	app.Rating = rating
}

// sets the number of users that rated this app on the store
func (app *FullApp) SetRatingCount(ratingCount string) {
	app.RatingCount = ratingCount
}

// sets the rating histogram for this app
func (app *FullApp) SetRatingHistogram(histogram map[string]string) {
	app.RatingHistogram = histogram
}

// sets the 'whats's new' section
func (app *FullApp) SetWhatsNew(wsNew string) {
	app.WhatsNew = wsNew
}

// sets the date of the last update
func (app *FullApp) SetLastUpdated(updated string) {
	app.LastUpdated = updated
}

// sets the size of the app
func (app *FullApp) SetAppSize(appSize string) {
	app.Size = appSize
}

// sets the number of downloads for this app
func (app *FullApp) SetNumOfDownloads(numOfDls string) {
	app.NumOfDownloads = numOfDls
}

// sets the app version
func (app *FullApp) SetAppVer(appVer string) {
	app.AppVer = appVer
}

// sets the minimum Android version required by the app to run (install)
func (app *FullApp) SetMinAndroidVer(minVer string) {
	app.MinAndroidVer = minVer
}

// sets the minimum age (should be) required by this app
func (app *FullApp) SetMinAgeRating(minAge string) {
	app.AgeRating = minAge
}

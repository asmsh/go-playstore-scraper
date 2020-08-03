// Package fields represents the info fields of the AppInfo struct
package fields

// AppField represents an app info field that we want to extract,
// like app name, dev name, icon url, ...
type AppField int

// The order(and hence the values) of these constants may change in future,
// so always use the values through this package.
const (
	// AllFields will extract all the fields of the app.
	AllFields AppField = iota
	// AppId will extract only the app id of the app, and populate the
	// AppInfo.AppId field with the data.
	AppId
	// AppUrl will extract only the url of the app, and populate the
	// AppInfo.AppUrl field with the data.
	AppUrl
	// IconUrls will extract only the urls of the icon image of the app,
	// and populate the AppInfo.IconUrls field with the data.
	IconUrls
	// AppName will extract only the name of the app, and populate the
	// AppInfo.AppName field with the data.
	AppName
	// DevInfo will extract only the name and page url of the developer
	// of the app, and populate the AppInfo.DevName and AppInfo.DevPageUrl
	// fields with the data.
	DevInfo
	// Category will extract only the name of the app category of the app,
	// and populate the AppInfo.Category field with the data.
	Category
	// FamilyCategory will extract only the name of the family category
	// of the app, and populate the AppInfo.FamilyCategory field with
	// the data. The data corresponding to this field may not exist in
	// the app page.
	FamilyCategory
	// InAppExperience will extract only the info related to the in-app
	// experience, like ads, and purchases, and populate the
	// AppInfo.InAppExperience field with the data. The data corresponding
	// to this field may not exist in the app page.
	InAppExperience
	// Price will extract only the price text of the app, and populate
	// the AppInfo.Price field with the data.
	Price
	// VideoTrailerUrls will extract only the urls of the video trailer
	// of the app, and populate the AppInfo.VideoTrailerUrls field with
	// the data. The data corresponding to this field may not exist in
	// the app page.
	VideoTrailerUrls
	// ScreenShotsUrls will extract only the urls of the screenshots of
	// the app, and populate the AppInfo.ScreenShotsUrls field with the
	// data. The data corresponding to this field may not exist in the
	// app page.
	ScreenShotsUrls
	// Description will extract only the description text of the app,
	// and populate the AppInfo.Description field with the data.
	Description
	// Rating will extract only the aggregate rating text of the app,
	// and populate the AppInfo.Rating field with the data.
	Rating
	// RatingCount will extract only the rating count text of the app,
	// and populate the AppInfo.RatingCount field with the data.
	RatingCount
	// RatingHistogram will extract only the rating histogram of the app,
	// and populate the AppInfo.RatingHistogram field with the data.
	RatingHistogram
	// WhatsNew will extract only the text of the what's new section in
	// the app page, and populate the AppInfo.WhatsNew field with the data.
	// The data corresponding to this field may not exist in the app page.
	WhatsNew

	// FIXME, the following 6 fields aren't extracted (issue #1, #2, #3, #4, #5, #6)
	LastUpdated
	Size
	NumOfDownloads
	Version
	MinAndroidVer
	AgeRating
)

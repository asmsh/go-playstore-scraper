package apps

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/asmsh/go-playstore-scraper/locales"
)

// the base part of the url for an app page
const appPageBaseUrl = "https://play.google.com/store/apps/details?id="

// AppInfo holds the extracted info for some app. It holds the extracted
// text, as-is, for each field, from each one's corresponding section in
// the app page of that app.
type AppInfo struct {
	// Country is the country of the app page which these fields are
	// extracted from, or locales.DefCountry if no specific country
	// was chosen.
	Country locales.Country
	// Language is the language of the app page which these fields are
	// extracted from, or locales.DefLanguage if no specific language
	// was chosen.
	Language locales.Language

	// AppUrl is the url of this app info instance, which is the url
	// that's used to extract these values from.
	AppUrl string `json:",omitempty"`
	// AppId is the id of this app, i.e. the apk name.
	AppId string `json:",omitempty"`
	// IconUrls holds the different urls of the icon image. It has two
	// elements, the first is the url of the default(low resolution) icon
	// image, and the second is the url of a higher resolution icon image.
	//
	// It will be nil if the required app fields doesn't has this field
	// within it.
	IconUrls []string `json:",omitempty"`
	// AppName is the name of the app as it appear on its app page.
	//
	// It will be empty if the required app fields doesn't has this field
	// within it.
	AppName string `json:",omitempty"`
	// DevName is the name of the developer of this app.
	//
	// It will be empty if the required app fields doesn't has this field
	// within it.
	DevName string `json:",omitempty"`
	// DevPageUrl is the url of the page of the developer of this app.
	//
	// It will be empty if the required app fields doesn't has this field
	// within it.
	DevPageUrl string `json:",omitempty"`
	// Category is the app category which this app belongs to.
	//
	// It will be empty if the required app fields doesn't has this field
	// within it.
	Category string `json:",omitempty"`
	// FamilyCategory is the family category which this app belongs to.
	//
	// It will be empty if this app doesn't belong to one, or the required
	// app fields doesn't has this field within it.
	FamilyCategory string `json:",omitempty"`
	// InAppExperience holds the info about the ads inside this app, and
	// whether this app offer in app purchases or not.
	//
	// It will be empty if the app doesn't have any of the above info, or
	// the required app fields doesn't has this field within it.
	InAppExperience string `json:",omitempty"`
	// Price holds the price related string as seen in this app's page.
	//
	// It will be empty if the required app fields doesn't has this field
	// within it.
	Price string `json:",omitempty"`
	// VideoTrailerUrls holds the urls related to the video trailer of this
	// app. It has two elements, the first is the url of the start image of
	// the video, and the second is the url of the video itself.
	//
	// It will be nil if this app doesn't provide a video trailer, or the
	// required app fields doesn't has this field within it.
	VideoTrailerUrls []string `json:",omitempty"`
	// ScreenShotsUrls holds the urls related to the screenshots of this
	// app. It has even number of elements, for every pair of successive
	// and distinct elements, the first element in the pair will be the
	// url of the default(low resolution) image of some arbitrary screenshot,
	// and the second will be the url of a higher resolution image of the
	// same screenshot.
	//
	// It will be nil if this app doesn't provide any screenshots, or the
	// required app fields doesn't has this field within it.
	ScreenShotsUrls []string `json:",omitempty"`
	// Description holds the description text of this app, as seen in this
	// app's page.
	//
	// It will be empty if this app doesn't provide a description, or the
	// required app fields doesn't has this field within it.
	Description string `json:",omitempty"`
	// Rating holds the aggregate rating string of this app, as seen in
	// this app's page.
	// Note: its value may need some preprocessing before converting it to
	// a number, for using it as a number use method 'AppInfo.RatingNum'.
	//
	// It will be empty if the required app fields doesn't has this field
	// within it.
	Rating string `json:",omitempty"`
	// RatingCount holds the total number of ratings for this app, as seen
	// in this app's page.
	// Note: its value may need some preprocessing before converting it to
	// a number, for using it as a number use method 'AppInfo.RatingCountNum'.
	//
	// It will be empty if the required app fields doesn't has this field
	// within it.
	RatingCount string `json:",omitempty"`
	// RatingHistogram holds the rating histogram bars from the rating box
	// of this app, as seen in this app's page. It's a map with keys equal
	// to the star number as a string("1"-"5"), and the 'value' is the number
	// of ratings for each star.
	//
	// It will be nil if the required app fields doesn't has this field
	// within it.
	//
	// FIXME, the rating histogram's values doesn't represent the number of
	//  ratings for each star (issue #7)
	RatingHistogram map[string]string `json:",omitempty"`
	// WhatsNew holds the what's new text of this app, as seen in this
	// app's page.
	//
	// It will be empty if this app doesn't provide any text in the what's
	// new section in its page, or the required app fields doesn't has this
	// field within it.
	WhatsNew string `json:",omitempty"`

	// FIXME, the following 6 fields aren't extracted (issue #1, #2, #3, #4, #5, #6)
	LastUpdated    string `json:",omitempty"`
	Size           string `json:",omitempty"`
	NumOfDownloads string `json:",omitempty"`
	Version        string `json:",omitempty"`
	MinAndroidVer  string `json:",omitempty"`
	AgeRating      string `json:",omitempty"`
}

// DefaultAppUrl returns the default url of this app, which is the url
// of the app without specifying the country or the language.
func (app AppInfo) DefaultAppUrl() string {
	if app.AppId == "" {
		return ""
	}
	return appPageBaseUrl + app.AppId
}

// ToJSON formats this app info as a JSON string. It returns a non-nil
// error if anything went wrong while formatting the app.
func (app AppInfo) ToJSON() (string, error) {
	ba, e := json.Marshal(app)
	if e != nil {
		return "", e
	}

	return string(ba), nil
}

func (app AppInfo) String() string {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("AppId: %q\n", app.AppId))
	sb.WriteString(fmt.Sprintf("AppUrl: %q\n", app.AppUrl))
	sb.WriteString(fmt.Sprintf("IconUrls: %q\n", app.IconUrls))
	sb.WriteString(fmt.Sprintf("AppName: %q\n", app.AppName))
	sb.WriteString(fmt.Sprintf("DevName: %q\n", app.DevName))
	sb.WriteString(fmt.Sprintf("DevPageUrl: %q\n", app.DevPageUrl))
	sb.WriteString(fmt.Sprintf("Category: %q\n", app.Category))
	sb.WriteString(fmt.Sprintf("FamilyCategory: %q\n", app.FamilyCategory))
	sb.WriteString(fmt.Sprintf("InAppExperience: %q\n", app.InAppExperience))
	sb.WriteString(fmt.Sprintf("Price: %q\n", app.Price))
	sb.WriteString(fmt.Sprintf("VideoTrailerUrls: %q\n", app.VideoTrailerUrls))
	sb.WriteString(fmt.Sprintf("ScreenShotsUrls: %q\n", app.ScreenShotsUrls))
	sb.WriteString(fmt.Sprintf("Description: %q\n", app.Description))
	sb.WriteString(fmt.Sprintf("Rating: %q\n", app.Rating))
	sb.WriteString(fmt.Sprintf("RatingCount: %q\n", app.RatingCount))
	sb.WriteString(fmt.Sprintf("RatingHistogram: %q\n", app.RatingHistogram))
	sb.WriteString(fmt.Sprintf("WhatsNew: %q\n", app.WhatsNew))
	sb.WriteString(fmt.Sprintf("LastUpdated: %q\n", app.LastUpdated))
	sb.WriteString(fmt.Sprintf("Size: %q\n", app.Size))
	sb.WriteString(fmt.Sprintf("NumOfDownloads: %q\n", app.NumOfDownloads))
	sb.WriteString(fmt.Sprintf("Version: %q\n", app.Version))
	sb.WriteString(fmt.Sprintf("MinAndroidVer: %q\n", app.MinAndroidVer))
	sb.WriteString(fmt.Sprintf("AgeRating: %q\n", app.AgeRating))

	return sb.String()
}

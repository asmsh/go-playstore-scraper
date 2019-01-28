package appPage

import (
	"fmt"
	"github.com/asmsh/go-playstore-scraper/engine/types/appField"
	"github.com/asmsh/go-playstore-scraper/engine/types/fullApp"
	"github.com/asmsh/go-playstore-scraper/engine/urls"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

// TODO, feat: add an option to choose the size of the screenshots returned.
// TODO, handle errors
// TODO, move the candidate test out of the switch on the data, and measure the speed

// a global var to track the element num for the required fields
var totalTagsCounter uint64 = 0
var openedTags int64 = 0

func ParseAppFromFileDebugOnly(filePath string, fields ...appField.AppField) (*fullApp.FullApp, uint64, int64, error) {
	file, e := os.Open(filePath)
	if e != nil {
		return nil, 0, 0, e
	}

	fields, e = appField.ValidateAppFields(fields)
	if e != nil {
		return nil, 0, 0, fmt.Errorf("error validating the required app fields")
	}

	tz := html.NewTokenizer(file)
	totalTagsCounter = 0
	openedTags = 0

	app, e := parseAppContent(tz, fields)

	return app, totalTagsCounter, openedTags, e
}

func ParseApp(appUrl *urls.AppUrl, fields ...appField.AppField) (*fullApp.FullApp, error) {
	return parseAppPage(appUrl, fields)
}

// ParseAppByUrl get the full information of an app by providing its store page URL, or otherwise an error.
func ParseAppByUrl(url string, fields ...appField.AppField) (*fullApp.FullApp, error) {
	appUrl, e := urls.ValidateAppPageURL(url)
	if e != nil {
		return nil, fmt.Errorf("failed to retreive the app info with error: %s", e.Error())
	}

	if app, e := parseAppPage(appUrl, fields); e != nil {
		return nil, fmt.Errorf("failed to retreive the app info with error: %s", e.Error())
	} else {
		return app, nil
	}
}

func requestAppPage(url string) (*http.Response, error) {
	resp, e := http.Get(url)
	if e != nil {
		return nil, fmt.Errorf("error with requesting the url: \n" + e.Error())
	}

	//TODO, feat: handle the possible responses
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("error with requesting the url: \n"+"response status is: %s", resp.Status)
	}

	return resp, nil
}

func parseAppPage(appUrl *urls.AppUrl, fields []appField.AppField) (*fullApp.FullApp, error) {
	var e error
	var tmpID, tmpURL string

	fields, e = appField.ValidateAppFields(fields)
	if e != nil {
		return nil, fmt.Errorf("failed to retreive the app info with error: %s", e.Error())
	}

	url := appUrl.String()

	resp, e := requestAppPage(url)
	if e != nil {
		return nil, fmt.Errorf("error with parsing the app page: " + e.Error())
	}
	defer resp.Body.Close()

	tz := html.NewTokenizer(resp.Body)
	totalTagsCounter = 0
	openedTags = 0

	app := new(fullApp.FullApp)
	if fields[0] == appField.AppID {
		tmpID = appUrl.AppID()
		if len(fields) > 1 {
			fields = fields[1:]
		} else {
			goto ret
		}
	}
	if fields[0] == appField.AppUrl {
		tmpURL = url
		if len(fields) > 1 {
			fields = fields[1:]
		} else {
			goto ret
		}
	}
	app, e = parseAppContent(tz, fields)
	if e != nil {
		return nil, fmt.Errorf("failed to retreive the app info with error: %s", e.Error())
	}

ret:
	app.SetAppID(tmpID)
	app.SetAppUrl(tmpURL)
	return app, nil

}

func parseAppContent(acTz *html.Tokenizer, fields []appField.AppField) (*fullApp.FullApp, error) {
	var app = new(fullApp.FullApp)

	for idx, currField := range fields {
		var prevField, nextField appField.AppField

		if openedTags <= 0 {
			// this means that we have exit from the app info content element
		}

		// previous field
		if idx-1 >= 0 {
			prevField = fields[idx-1]
		}
		// next field
		if idx+1 < len(fields) {
			nextField = fields[idx+1]
		}

		switch {
		case currField == appField.IconUrls:
			iconUrls, e := extractIconURL(acTz)
			if e != nil {
				return nil, fmt.Errorf("error getting the icon url: " + e.Error())
			}
			app.SetIconUrls(iconUrls)
		case currField == appField.AppName:
			appName, e := extractAppName(acTz)
			if e != nil {
				return nil, fmt.Errorf("error getting the app name: " + e.Error())
			}
			app.SetAppName(appName)
		case currField == appField.DevName || currField == appField.DevPageUrl:
			// DevName field will always be exactly before DevPageUrl field if they will ever appear together.
			if prevField != appField.DevName {
				devUrl, devName, e := extractDevInfo(acTz)
				if e != nil {
					return nil, fmt.Errorf("error getting the developer info: " + e.Error())
				}

				if currField == appField.DevName {
					app.SetDevPageUrl(devUrl)

					if nextField == appField.DevPageUrl {
						app.SetDevName(devName)
					}
				}
				if currField == appField.DevPageUrl {
					app.SetDevName(devName)
				}
			}
		case currField == appField.Category:
			_, catName, e := extractCategoryInfo(acTz)
			// TODO, feat: handle the retrieved 'catUrl'
			if e != nil {
				return nil, fmt.Errorf("error getting the category info: " + e.Error())
			}
			app.SetCategory(catName)
		case currField == appField.InAppOffering:
			offeringString, e := extractInAppOffering(acTz)
			// TODO, feat: check for the returned strings(with respect to the url language),
			//  and populate their corresponding boolean fields.
			if e != nil {
				return nil, fmt.Errorf("error getting the in app offering info: " + e.Error())
			}
			app.SetInAppOffering(offeringString)
		case currField == appField.Price:
			price, e := extractPrice(acTz)
			if e != nil {
				return nil, fmt.Errorf("error getting the app price: " + e.Error())
			}
			app.SetAppPrice(price)
		case currField == appField.VideoTrailerUrls || currField == appField.ScreenShotsUrls:
			if prevField != appField.VideoTrailerUrls {
				videoURLs, imagesURLs, e := extractMediaUrls(acTz)
				if e != nil {
					return nil, fmt.Errorf("error getting the app media urls: " + e.Error())
				}
				if currField == appField.VideoTrailerUrls {
					app.SetVideoTrailerUrls(videoURLs)

					if nextField == appField.ScreenShotsUrls {
						app.SetScreenShotsUrls(imagesURLs)
					}
				}
				if currField == appField.ScreenShotsUrls {
					app.SetScreenShotsUrls(imagesURLs)
				}
			}
		case currField == appField.Description:
			appDesc, e := extractDescription(acTz)
			if e != nil {
				return nil, fmt.Errorf("error getting the app's description: " + e.Error())
			}
			app.SetDescription(appDesc)
		case currField == appField.Rating || currField == appField.RatingCount:
			if prevField != appField.Rating {
				rating, ratingCount, e := extractRatingInfo(acTz)
				if e != nil {
					return nil, fmt.Errorf("error getting the app's rating info: " + e.Error())
				}
				if currField == appField.Rating {
					app.SetRating(rating)

					if nextField == appField.RatingCount {
						app.SetRatingCount(ratingCount)
					}
				}
				if currField == appField.RatingCount {
					app.SetRatingCount(ratingCount)
				}
			}
		case currField == appField.RatingHistogram:
			histogram, e := extractRatingHistogram(acTz)
			if e != nil {
				return nil, fmt.Errorf("error getting the app's histogram: " + e.Error())
			}
			app.SetRatingHistogram(histogram)
		case currField == appField.WhatsNew:
			whatsNew, e := extractWhatsNew(acTz)
			if e != nil {
				return nil, fmt.Errorf("error getting the app's whatsnew: " + e.Error())
			}
			app.SetWhatsNew(whatsNew)
		}
	}

	return app, nil
}

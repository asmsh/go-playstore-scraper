package appPage

import (
	"fmt"
	"github.com/asmsh/go-playstore-scraper/engine/restype/fullApp"
	"github.com/asmsh/go-playstore-scraper/engine/urls"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
)

// TODO, add an option to choose the size of the screenshots returned.
// TODO, handle errors
// TODO, move the candidate test out of the switch on the data, and measure the speed
// TODO, revise the chain of func calls for each parsing scenario

// a global var to track the element num for the required fields
var totalTagsCounter uint64 = 0
var openedTags int64 = 0

func ParseAppFromFileDebugOnly(filePath string) (*fullApp.App, uint64, int64, error) {
	file, e := os.Open(filePath)
	if e != nil {
		return nil, 0, 0, e
	}

	app, e := parsePage(file)

	return app, totalTagsCounter, openedTags, e
}

func ParseApp(appUrl *urls.AppUrl) (*fullApp.App, error) {
	url := appUrl.String()

	if app, e := requestPage(url); e != nil {
		return nil, fmt.Errorf("failed to retreive the app info with error: %s", e.Error())
	} else {
		app.SetAppID(appUrl.AppID())
		app.SetAppUrl(url)
		return app, nil
	}
}

// ParseAppByID get the full information of an app by providing its appID, or otherwise an error,
// this will assume the language is 'EN_US' and the country is 'US',
// for custom language and/or country page use the 'builder'
func ParseAppByID(id string) (*fullApp.App, error) {
	appUrl, e := urls.NewAppUrl(id)
	if e != nil {
		return nil, fmt.Errorf("failed to retreive the app info with error: %s", e.Error())
	}

	url := appUrl.String()

	if app, e := requestPage(url); e != nil {
		return nil, fmt.Errorf("failed to retreive the app info with error: %s", e.Error())
	} else {
		app.SetAppID(appUrl.AppID())
		app.SetAppUrl(url)
		return app, nil
	}
}

// ParseAppByUrl get the full information of an app by providing its store page URL, or otherwise an error
func ParseAppByUrl(url string) (*fullApp.App, error) {
	qParams, e := urls.ValidateAppPageURL(url)
	if e != nil {
		return nil, fmt.Errorf("failed to retreive the app info with error: %s", e.Error())
	}

	if app, e := requestPage(url); e != nil {
		return nil, fmt.Errorf("failed to retreive the app info with error: %s", e.Error())
	} else {
		app.SetAppID(qParams.AppID)
		app.SetAppUrl(url)
		return app, nil
	}
}

func requestPage(url string) (*fullApp.App, error) {
	resp, e := http.Get(url)
	if e != nil {
		return nil, fmt.Errorf("error with requesting the url: \n" + e.Error())
	}

	//TODO, feat: handle the possible responses
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("error with requesting the url: \n"+"response status is: %s", resp.Status)
	}

	return parsePage(resp.Body)
}

func parsePage(respBody io.ReadCloser) (*fullApp.App, error) {
	//closing the response body at the end of this function
	defer respBody.Close()

	tz := html.NewTokenizer(respBody)
	totalTagsCounter = 0
	openedTags = 0
	var inMainBody bool

mainLoop:
	for {
		tt := tz.Next()
		totalTagsCounter++
		switch tt {
		case html.StartTagToken:
			t := tz.Token()
			switch t.Data {
			case "div":
				// skip non candidate tags to improve the speed
				if len(t.Attr) == 0 {
					continue mainLoop
				}

				for _, attr := range t.Attr {
					if attr.Key == "class" && attr.Val == "UTg3hd" {
						// this means that we are now in the body of the page's main content(app info & suggested apps)
						inMainBody = true
						break mainLoop
					}
				}
			}
		case html.ErrorToken:
			break mainLoop
		}
	}
	if !inMainBody {
		return nil, fmt.Errorf("error: the expected structre isn't valid")
	} else {
		return parseAppContent(tz)
	}
}

func parseAppContent(acTz *html.Tokenizer) (*fullApp.App, error) {
	var app = new(fullApp.App)

	iconUrls, e := extractIconURL(acTz)
	if e != nil {
		return nil, fmt.Errorf("error getting the icon url: " + e.Error())
	}
	if openedTags <= 0 {
		// this means that we have exit from the app info content element
	}

	appName, e := extractAppName(acTz)
	if e != nil {
		return nil, fmt.Errorf("error getting the app name: " + e.Error())
	}
	if openedTags <= 0 {
		// this means that we have exit from the app info content element
	}

	devUrl, devName, e := extractDevInfo(acTz)
	if e != nil {
		return nil, fmt.Errorf("error getting the developer info: " + e.Error())
	}
	if openedTags <= 0 {
		// this means that we have exit from the app info content element
	}

	_, catName, e := extractCategoryInfo(acTz)
	// TODO, feat, handle the retrieved 'catUrl'
	if e != nil {
		return nil, fmt.Errorf("error getting the category info: " + e.Error())
	}
	if openedTags <= 0 {
		// this means that we have exit from the app info content element
	}

	offeringString, e := extractInAppOffering(acTz)
	// TODO, feat, check for the returned strings(with respect to the query language) and populate their corresponding boolean fields.
	if e != nil {
		return nil, fmt.Errorf("error getting the in app offering info: " + e.Error())
	}
	if openedTags <= 0 {
		// this means that we have exit from the app info content element
	}

	price, e := extractPrice(acTz)
	if e != nil {
		return nil, fmt.Errorf("error getting the app price: " + e.Error())
	}
	if openedTags <= 0 {
		// this means that we have exit from the app info content element
	}

	videoURLs, imagesURLs, e := extractMediaUrls(acTz)
	if e != nil {
		return nil, fmt.Errorf("error getting the app media urls: " + e.Error())
	}
	if openedTags <= 0 {
		// this means that we have exit from the app info content element
	}

	appDesc, e := extractDescription(acTz)
	if e != nil {
		return nil, fmt.Errorf("error getting the app's description: " + e.Error())
	}
	if openedTags <= 0 {
		// this means that we have exit from the app info content element
	}

	rating, ratingCount, e := extractRatingInfo(acTz)
	if e != nil {
		return nil, fmt.Errorf("error getting the app's rating info: " + e.Error())
	}
	if openedTags <= 0 {
		// this means that we have exit from the app info content element
	}

	histogram, e := extractRatingHistogram(acTz)
	if e != nil {
		return nil, fmt.Errorf("error getting the app's histogram: " + e.Error())
	}
	if openedTags <= 0 {
		// this means that we have exit from the app info content element
	}

	whatsNew, e := extractWhatsNew(acTz)
	if e != nil {
		return nil, fmt.Errorf("error getting the app's whatsnew: " + e.Error())
	}
	if openedTags <= 0 {
		// this means that we have exit from the app info content element
	}

	app.SetIconUrls(iconUrls)
	app.SetAppName(appName)
	app.SetDevPageUrl(devUrl)
	app.SetDevName(devName)
	app.SetCategory(catName)
	app.SetInAppOffering(offeringString)
	app.SetAppPrice(price)
	app.SetVideoTrailerUrls(videoURLs)
	app.SetScreenShotsUrls(imagesURLs)
	app.SetDescription(appDesc)
	app.SetRating(rating)
	app.SetRatingCount(ratingCount)
	app.SetRatingHistogram(histogram)
	app.SetWhatsNew(whatsNew)

	return app, nil
}

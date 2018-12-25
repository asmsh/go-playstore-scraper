package appPage

import (
	"fmt"
	"github.com/asmsh/go-playstore-scraper/engine/restype/fullApp"
	"github.com/asmsh/go-playstore-scraper/engine/urls"

	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
)

// TODO, add an option to choose the size of the screenshots returned.
// TODO, handle errors
// TODO, move the candidate test out of the switch on the data, and measure the speed

const devPageUrlPrefix string = "play.google.com/store/apps/dev"
const catPageUrlPrefix string = "https://play.google.com/store/apps/category/"

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

func extractIconURL(acTz *html.Tokenizer) ([]string, error) {
	var iconUrlLoRes, iconUrlHiRes string
	var iconUrls = []string{"", ""}
	var containerTagNum int64

containerLoop:
	for {
		if openedTags < containerTagNum {
			// this means that we have exited the container without finding the target element,
			// so we will break the loop as this is an indication on a expected wrong structure from the parser.
			break containerLoop
		}

		tt := acTz.Next()
		totalTagsCounter++
		switch tt {
		case html.StartTagToken:
			t := acTz.Token()
			switch t.Data {
			case "div":
				openedTags++
				// skip non candidate tags to improve the speed
				if len(t.Attr) == 0 {
					continue containerLoop
				}

				for _, attr := range t.Attr {
					if attr.Key == "class" && attr.Val == "dQrBL" {
						containerTagNum = openedTags
						break containerLoop
					}
				}
			default:
				// count even any tag that we might close in parsing.
				handleDefaultStartTagToken(t)
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break containerLoop
		}
	}
imgUrlLoop:
	for {
		if openedTags < containerTagNum {
			break imgUrlLoop
		}

		tt := acTz.Next()
		totalTagsCounter++
		switch tt {
		case html.StartTagToken:
			t := acTz.Token()
			switch t.Data {
			// <img> tags doesn't count against opened tags.
			case "img":
				/*// skip non candidate tags to improve the speed
				if len(t.Attr) < 5 {
					continue imgUrlLoop
				}*/

				var tmpLoRes string
				var tmpHiRes string
				var signs int8
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "class" && attr.Val == "T75of ujDFqe":
						signs++
					case attr.Key == "alt" && attr.Val == "Cover art":
						signs++
					case attr.Key == "itemprop" && attr.Val == "image":
						signs++
					case attr.Key == "src" && attr.Val != "":
						signs++
						tmpLoRes = attr.Val
					case attr.Key == "srcset" && attr.Val != "":
						signs++
						tmpHiRes = attr.Val
					}
				}
				if signs == 5 {
					tmpHiRes, e := formatHiResImgUrl(tmpHiRes)
					if e != nil {
						return nil, e
					}
					iconUrlLoRes = tmpLoRes
					iconUrlHiRes = tmpHiRes
					break imgUrlLoop
				}
			default:
				// count even any tag that we might close in parsing.
				handleDefaultStartTagToken(t)
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break imgUrlLoop
		}
	}
	if iconUrlLoRes == "" || iconUrlHiRes == "" {
		return nil, fmt.Errorf("couldn't extract the app icon urls")
	}

	iconUrls[0] = iconUrlLoRes
	iconUrls[1] = iconUrlHiRes
	return iconUrls, nil
}

func extractAppName(acTz *html.Tokenizer) (string, error) {
	var appName string
	var nameFound bool
	var containerTagNum int64
containerLoop:
	for {
		if openedTags < containerTagNum {
			break containerLoop
		}

		tt := acTz.Next()
		totalTagsCounter++
		switch tt {
		case html.StartTagToken:
			t := acTz.Token()
			switch t.Data {
			case "h1":
				openedTags++
				// skip non candidate tags to improve the speed
				if len(t.Attr) < 2 {
					continue containerLoop
				}

				var signs int8
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "class" && attr.Val == "AHFaub":
						signs++
					case attr.Key == "itemprop" && attr.Val == "name":
						signs++
					}
				}
				if signs == 2 {
					containerTagNum = openedTags
					break containerLoop
				}
			default:
				// count even any tag that we might close in parsing.
				handleDefaultStartTagToken(t)
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break containerLoop
		}
	}
appNameLoop:
	for {
		if openedTags < containerTagNum {
			break appNameLoop
		}

		tt := acTz.Next()
		totalTagsCounter++
		switch tt {
		case html.StartTagToken:
			t := acTz.Token()
			switch t.Data {
			case "span":
				openedTags++
				// to handle breaking out of the containerLoop on errors without finding the target tag.
				if containerTagNum != 0 && openedTags > containerTagNum {
					nameFound = true
				}
			default:
				// count even any tag that we might close in parsing.
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			if nameFound {
				t := acTz.Token()
				appName = t.Data
				nameFound = false
				break appNameLoop
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break appNameLoop
		}
	}
	if appName == "" {
		return "", fmt.Errorf("couldn't extract the app name")
	}
	return appName, nil
}

// (devURL, devName, error)
func extractDevInfo(acTz *html.Tokenizer) (string, string, error) {
	var devName, devURL string
	var urlFound bool
	var containerTagNum int64
mainLoop:
	for {
		tt := acTz.Next()
		totalTagsCounter++
		switch tt {
		case html.StartTagToken:
			t := acTz.Token()
			switch t.Data {
			case "a":
				// count only the tags that we actually open.
				openedTags++
				// skip non candidate tags to improve the speed
				if len(t.Attr) < 2 {
					continue mainLoop
				}

				var signs int8
				var tmp string
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "class" && attr.Val == "hrTbp R8zArc":
						signs++
					case attr.Key == "href" && strings.Contains(attr.Val, devPageUrlPrefix):
						signs++
						tmp = attr.Val
					}
				}
				if signs == 2 {
					devURL = tmp
					urlFound = true
					containerTagNum = openedTags
				}
			default:
				// count even any tag that we might close in parsing.
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			if urlFound && containerTagNum <= openedTags {
				t := acTz.Token()
				devName = t.Data
				urlFound = false
				break mainLoop
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break mainLoop
		}
	}
	if devURL == "" || devName == "" {
		return "", "", fmt.Errorf("couldn't find the dev page url or the dev name")
	}

	return devURL, devName, nil
}

// (catUrl, catName, err)
func extractCategoryInfo(acTz *html.Tokenizer) (string, string, error) {
	var catUrl, catName string
	var urlFound bool
	var containerTagNum int64
mainLoop:
	for {
		tt := acTz.Next()
		totalTagsCounter++
		switch tt {
		case html.StartTagToken:
			t := acTz.Token()
			switch t.Data {
			case "a":
				// count only the tags that we actually open.
				openedTags++
				// skip non candidate tags to improve the speed
				if len(t.Attr) < 3 {
					continue mainLoop
				}

				var signs int8
				var tmp string
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "class" && attr.Val == "hrTbp R8zArc":
						signs++
					case attr.Key == "itemprop" && attr.Val == "genre":
						signs++
					case attr.Key == "href" && strings.Contains(attr.Val, catPageUrlPrefix):
						signs++
						tmp = attr.Val
					}
				}
				if signs == 3 {
					catUrl = tmp
					urlFound = true
					containerTagNum = openedTags
				}
			default:
				// count even any tag that we might close in parsing.
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			// we can do this here(check for equality), as the text should be in the same level.
			if urlFound && containerTagNum == openedTags {
				t := acTz.Token()
				catName = t.Data
				urlFound = false
				break mainLoop
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break mainLoop
		}
	}
	if catUrl == "" || catName == "" {
		return "", "", fmt.Errorf("couldn't find the dev page url or the dev name")
	}

	return catUrl, catName, nil
}

// (containsAds, inAppPurchase, error)
func extractInAppOffering(acTz *html.Tokenizer) (string, error) {
	var offeringString = ""
	var containerFound bool
	var containerTagNum int64
mainLoop:
	for {
		tt := acTz.Next()
		totalTagsCounter++
		switch tt {
		case html.StartTagToken:
			t := acTz.Token()
			switch t.Data {
			case "div":
				openedTags++
				// skip non candidate tags to improve the speed
				if len(t.Attr) == 0 {
					continue mainLoop
				}

				for _, attr := range t.Attr {
					switch {
					case attr.Key == "class" && attr.Val == "wE7q7b":
						// reaching this case means we passed the position without finding any offerings.
						break mainLoop
					case attr.Key == "class" && attr.Val == "rxic6":
						// this is the container we need.
						containerFound = true
						containerTagNum = openedTags
					}
				}
			default:
				// count even any tag that we might close in parsing
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			// we can do this here(check for equality), as all the offering strings should be in the same level.
			if containerFound && containerTagNum == openedTags {
				t := acTz.Token()
				content := strings.TrimSpace(t.Data)
				if len(offeringString) == 0 {
					offeringString = content
				} else {
					offeringString += ", " + content
				}
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break mainLoop
		}
	}
	if containerFound && offeringString == "" {
		return "", fmt.Errorf("the offering container is found, however we couldn't extract the data")
	} else {
		return offeringString, nil
	}
}

func extractPrice(acTz *html.Tokenizer) (string, error) {
	var price string
	var containerFound bool
	var containerTagNum int64
mainLoop:
	for {
		if openedTags < containerTagNum {
			break mainLoop
		}

		tt := acTz.Next()
		totalTagsCounter++
		switch tt {
		case html.StartTagToken:
			t := acTz.Token()
			switch t.Data {
			case "span":
				openedTags++
				// skip non candidate tags to improve the speed
				if len(t.Attr) < 2 {
					continue mainLoop
				}

				var signs int8
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "itemprop" && attr.Val == "offers":
						fallthrough
					case attr.Key == "itemtype" && attr.Val == "https://schema.org/Offer":
						signs++
					}
				}
				if signs == 2 {
					containerFound = true
					containerTagNum = openedTags
				}
			case "meta":
				if containerFound && openedTags == containerTagNum {
					// skip non candidate tags to improve the speed
					/*if len(t.Attr) < 2 {
						continue mainLoop
					}*/

					var tmp string
					var signs int8
					for _, attr := range t.Attr {
						switch {
						case attr.Key == "itemprop" && attr.Val == "price":
							signs++
						case attr.Key == "content" && attr.Val != "":
							signs++
							tmp = attr.Val
						}
					}
					if signs == 2 {
						price = tmp
						break mainLoop
					}
				}
			default:
				handleDefaultStartTagToken(t)
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break mainLoop
		}
	}
	if price == "" {
		return "", fmt.Errorf("couldn't find the app price")
	}

	return price, nil
}

func extractMediaUrls(acTz *html.Tokenizer) ([]string, []string, error) {
	var startImgUrl, videoUrl string
	var videoURLs = []string{"", ""}
	var imagesURLs = make([]string, 0, 30)
	var mediaContainerFound bool
	var mediaContainerTagNum int64
	var startImgContainerFound bool
	var startImgContainerTagNum int64
	var videoContainerFound bool
	var videoContainerTagNum int64

videoLoop:
	for {
		if openedTags < mediaContainerTagNum {
			break videoLoop
		}

		tt := acTz.Next()
		totalTagsCounter++
		switch tt {
		case html.StartTagToken:
			t := acTz.Token()
			switch t.Data {
			case "div":
				openedTags++
				// skip non candidate tags to improve the speed
				if len(t.Attr) < 1 {
					continue videoLoop
				}

			attrLoop:
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "class" && attr.Val == "KDxLi":
						// the container of all the media is found, containing the trailer video with its start screen,
						// and all the containers for the screenshots.
						mediaContainerFound = true
						mediaContainerTagNum = openedTags
						// break out of the attr loop, cause all these cases are mutually exclusive.
						break attrLoop
					case mediaContainerFound && attr.Key == "class" && attr.Val == "MSLVtf NIc6yf":
						startImgContainerFound = true
						startImgContainerTagNum = openedTags
						break attrLoop
					case startImgContainerFound && attr.Key == "class" && attr.Val == "TdqJUe":
						videoContainerFound = true
						videoContainerTagNum = openedTags
						break attrLoop
					}
				}
			case "img":
				if startImgContainerFound && openedTags == startImgContainerTagNum {
					// skip non candidate tags to improve the speed
					/*if len(t.Attr) < 2 {
						continue videoLoop
					}*/

					var tmp string
					var signs int8
					for _, attr := range t.Attr {
						switch {
						case attr.Key == "class" && attr.Val == "T75of lxGQyd":
							signs++
						case attr.Key == "src" && attr.Val != "":
							signs++
							tmp = attr.Val
						}
					}
					if signs == 2 {
						startImgUrl = tmp
					}
				}
			case "button":
				openedTags++
				if videoContainerFound && openedTags == videoContainerTagNum+1 {
					// skip non candidate tags to improve the speed
					/*if len(t.Attr) < 2 {
						continue videoLoop
					}*/

					var tmp string
					var signs int8
					for _, attr := range t.Attr {
						switch {
						/*case attr.Key == "jscontroller" && attr.Val == "HnDLGf":
							signs++*/
						case attr.Key == "class" && attr.Val == "lgooh  ":
							signs++
						case attr.Key == "data-trailer-url" && attr.Val != "":
							signs++
							tmp = attr.Val
						}
					}
					if signs == 2 {
						videoUrl = tmp
						// we have found the video url already,
						// so, break from the video loop to the screenshots loop.
						break videoLoop
					}
				} else {
					// check if the video container hasn't been found and we have skipped to the screenshots directly.
					// skip non candidate tags to improve the speed
					/*if len(t.Attr) < 1 {
						continue videoLoop
					}*/

					for _, attr := range t.Attr {
						switch {
						case attr.Key == "class" && attr.Val == "NIc6yf":
							// the button found is the container for a screenshot,
							// so, break from the video loop to the screenshots loop.
							break videoLoop
						}
					}
				}
			default:
				// count even any tag that we might close in parsing
				handleDefaultStartTagToken(t)
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break videoLoop
		}
	}

screenshotsLoop:
	for {
		// break the loop if we exit the target container tag
		if openedTags < mediaContainerTagNum || !mediaContainerFound {
			break screenshotsLoop
		}

		tt := acTz.Next()
		totalTagsCounter++
		switch tt {
		case html.StartTagToken:
			t := acTz.Token()
			switch t.Data {
			case "img":
				// skip non candidate tags to improve the speed
				/*if len(t.Attr) < 5 {
					continue screenshotsLoop
				}*/

				var tmpSrc string
				var tmpSrcset string
				var signs int8
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "src" && attr.Val != "":
						signs++
						tmpSrc = attr.Val
					case attr.Key == "srcset" && attr.Val != "":
						signs++
						tmpSrcset = attr.Val
					case attr.Key == "alt" && attr.Val == "Screenshot Image":
						signs++
						/*case attr.Key == "class" && attr.Val == "T75of lxGQyd":
							signs++
						case attr.Key == "itemprop" && attr.Val == "image":
							signs++*/
					}
				}
				if signs == 3 {
					highRes, e := formatHiResImgUrl(tmpSrcset)
					if e != nil {
						highRes = tmpSrcset
					}
					imagesURLs = append(imagesURLs, tmpSrc, highRes)
				}
			default:
				// count even any tag that we might close in parsing
				handleDefaultStartTagToken(t)
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break screenshotsLoop
		}
	}

	if mediaContainerFound {
		if startImgContainerFound && (startImgUrl == "" || videoUrl == "") {
			return nil, nil, fmt.Errorf("the video container is found, " +
				"however, we couldn't extract the required urls")

		} else if !startImgContainerFound {
			return nil, imagesURLs, nil
		} else {
			videoURLs[0] = startImgUrl
			videoURLs[1] = videoUrl
			return videoURLs, imagesURLs, nil
		}
	} else {
		return nil, nil, fmt.Errorf("the media container is not found")
	}
}

func extractDescription(acTz *html.Tokenizer) (string, error) {
	/*var descContainerFound bool
	var descContainerTagNum int64
	var desc = ""*/
	// the second method related vars
	var descContainerFound2 bool
	var descContainerTagNum2 int64
	var desc2 = ""

descLoop:
	for {
		/*if openedTags < descContainerTagNum {
			break descLoop
		}*/
		if openedTags < descContainerTagNum2 {
			break descLoop
		}

		tt := acTz.Next()
		totalTagsCounter++

		switch tt {
		case html.StartTagToken:
			t := acTz.Token()

			switch t.Data {
			case "div":
				openedTags++

				if len(t.Attr) < 1 {
					continue descLoop
				}

				for _, attr := range t.Attr {
					switch {
					/*case attr.Key == "class" && attr.Val == "W4P4ne ":
						// related to the first method.
						descContainerFound = true
						descContainerTagNum = openedTags
						break*/
					case attr.Key == "jsname" && attr.Val == "sngebd":
						// related to the second method.
						descContainerFound2 = true
						descContainerTagNum2 = openedTags
						break
					}
				}
			/*case "meta":
				// this is used as a first method to retrieve the description, which get it from the meta tag.
				if !descContainerFound || openedTags != descContainerTagNum || len(t.Attr) < 2 {
					continue descLoop
				}

				var signs uint8
				var tmp string
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "itemprop" && attr.Val == "description":
						signs++
					case attr.Key == "content" && attr.Val != "":
						signs++
						tmp = attr.Val
					}
				}
				if signs == 2 {
					desc = tmp
					break descLoop
				}*/
			default:
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			// this is used as a second method to retrieve the description, which gets it from the text tags.
			if descContainerFound2 && openedTags == descContainerTagNum2 {
				t := acTz.Token()
				desc2 += t.Data
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break descLoop
		}
	}

	/*if !descContainerFound {
		return "", fmt.Errorf("the description tag is not found")
	} else if desc == "" {
		return "", fmt.Errorf("the description tag is found however we couldn't extract the description text")
	} else {
		return desc, nil
	}*/
	if !descContainerFound2 {
		return "", fmt.Errorf("the description tag is not found")
	} else if desc2 == "" {
		return "", fmt.Errorf("the description tag is found however we couldn't extract the description text")
	} else {
		return desc2, nil
	}
}

func extractRatingInfo(acTz *html.Tokenizer) (string, string, error) {
	var ratingContainerFound bool
	var ratingContainerTagNum int64
	var ratingCountOuterContainerFound bool
	var ratingCountInnerContainerFound bool
	var ratingCountOuterContainerTagNum int64
	var ratingCountInnerContainerTagNum int64
	var rating string
	var ratingCount string

ratingLoop:
	for {
		if openedTags < ratingContainerTagNum {
			break ratingLoop
		}

		tt := acTz.Next()
		totalTagsCounter++

		switch tt {
		case html.StartTagToken:
			t := acTz.Token()

			switch t.Data {
			case "div":
				openedTags++

				if len(t.Attr) != 2 {
					continue ratingLoop
				}

				var signs uint8
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "class" && attr.Val == "BHMmbe":
						signs++
					case attr.Key == "aria-label" && attr.Val != "":
						signs++
					}
				}
				if signs == 2 {
					ratingContainerFound = true
					ratingContainerTagNum = openedTags
				}
			default:
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			if ratingContainerFound && openedTags == ratingContainerTagNum {
				rating = acTz.Token().Data
				break ratingLoop
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break ratingLoop
		}
	}

ratingCountLoop:
	for {
		if openedTags < ratingCountOuterContainerTagNum {
			break ratingCountLoop
		}

		tt := acTz.Next()
		totalTagsCounter++

		switch tt {
		case html.StartTagToken:
			t := acTz.Token()

			switch t.Data {
			case "span":
				openedTags++

				if len(t.Attr) < 1 {
					continue ratingCountLoop
				}

				var signs uint8
			attrLoop:
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "class" && attr.Val == "EymY4b":
						ratingCountOuterContainerFound = true
						ratingCountOuterContainerTagNum = openedTags
						break attrLoop
					case ratingCountOuterContainerFound && attr.Key == "class" && attr.Val == "":
						signs++
					case ratingCountOuterContainerFound && attr.Key == "aria-label" && attr.Val != "":
						signs++
					}
				}
				if signs == 2 && openedTags > ratingCountOuterContainerTagNum {
					ratingCountInnerContainerFound = true
					ratingCountInnerContainerTagNum = openedTags
				}
			default:
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			if ratingCountInnerContainerFound && openedTags == ratingCountInnerContainerTagNum {
				ratingCount = acTz.Token().Data
				break ratingCountLoop
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break ratingCountLoop
		}
	}

	if ratingContainerFound && rating == "" {
		return "", "", fmt.Errorf("the rating container is found, however we couldn't extract the rating")
	} else if !ratingContainerFound {
		return "", "", fmt.Errorf("the rating container is not found")
	}
	if ratingCountOuterContainerFound && rating == "" {
		return "", "", fmt.Errorf("the rating count container is found, however we couldn't extract the rating count")
	} else if !ratingCountOuterContainerFound {
		return "", "", fmt.Errorf("the rating count container is not found")
	}

	return rating, ratingCount, nil
}

func extractRatingHistogram(acTz *html.Tokenizer) (map[string]string, error) {
	var histogramContainerFound bool
	var histogramContainerTagNum int64
	var barNumContainerFound bool
	var histogram = make(map[string]string, 5)
	currentBar := ""

mainLoop:
	for {
		if openedTags < histogramContainerTagNum {
			break mainLoop
		}

		tt := acTz.Next()
		totalTagsCounter++

		switch tt {
		case html.StartTagToken:
			t := acTz.Token()

			switch t.Data {
			case "div":
				openedTags++

				if len(t.Attr) != 1 {
					continue mainLoop
				}

				for _, attr := range t.Attr {
					switch {
					case attr.Key == "class" && attr.Val == "VEF2C":
						histogramContainerFound = true
						histogramContainerTagNum = openedTags
					}
				}
			case "span":
				openedTags++

				if len(t.Attr) < 1 || !histogramContainerFound {
					continue mainLoop
				}

				var signs uint8
				var tmpValue string
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "class" && attr.Val == "Gn2mNd":
						barNumContainerFound = true
						continue mainLoop
					/*case attr.Key == "style" && attr.Val != "":
						signs++*/
					case attr.Key == "class" && strings.Contains(attr.Val, "L2o20d"):
						signs++
					case attr.Key == "title" && attr.Val != "":
						signs++
						tmpValue = attr.Val
					}
					if signs == 2 && currentBar != "" {
						histogram[currentBar] = tmpValue
						currentBar = ""
						if len(histogram) == 5 {
							break mainLoop
						}
					}
				}
			default:
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			if barNumContainerFound && currentBar == "" {
				currentBar = acTz.Token().Data
				barNumContainerFound = false
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break mainLoop
		}
	}

	if histogramContainerFound && len(histogram) == 5 {
		return histogram, nil
	} else if len(histogram) != 5 {
		return nil, fmt.Errorf("the histogram container is not found, however we couldn't extract the histogram data")
	} else {
		return nil, fmt.Errorf("the histogram container is not found")
	}
}

// TODO, extractWhatsNew
func extractWhatsNew(acTz *html.Tokenizer) (string, error) {
	var wnOuterContainerFound bool // whatsNewOuterContainer
	var wnOuterContainerTagNum int64
	var wnInnerContainerFound bool
	var wnInnerContainerTagNum int64
	var whatsNewDataFound bool
	var whatsNew string

mainLoop:
	for {
		if openedTags < wnOuterContainerTagNum {
			break mainLoop
		}

		tt := acTz.Next()
		totalTagsCounter++

		switch tt {
		case html.StartTagToken:
			t := acTz.Token()

			switch t.Data {
			case "c-wiz":
				openedTags++

				if len(t.Attr) < 4 {
					continue mainLoop
				}

				var signs uint8
				var tmp string
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "jsrenderer" && attr.Val != "":
						tmp = attr.Val
						fallthrough
					case attr.Key == "jsshadow" && attr.Val == "":
						fallthrough
					case attr.Key == "jsdata" && attr.Val != "":
						fallthrough
					case attr.Key == "jsmodel" && attr.Val == "hc6Ubd":
						signs++
					}
				}
				if signs == 4 {
					if tmp == "FzdkFd" {
						wnOuterContainerFound = true
						wnOuterContainerTagNum = openedTags
					} else if tmp == "Wnurre" {
						wnOuterContainerFound = false
						break mainLoop
					}
				}
			case "div":
				openedTags++

				if len(t.Attr) < 3 || !wnOuterContainerFound {
					continue mainLoop
				}

				var signs uint8
				for _, attr := range t.Attr {
					switch {
					case attr.Key == "class" && attr.Val == "DWPxHb":
						fallthrough
					case attr.Key == "jsname" && attr.Val == "bN97Pc":
						fallthrough
					case attr.Key == "itemprop" && attr.Val == "description":
						signs++
					}
				}
				if signs == 3 {
					wnInnerContainerFound = true
					wnInnerContainerTagNum = openedTags
				}
			case "content":
				openedTags++

				if wnInnerContainerFound && openedTags == wnInnerContainerTagNum+1 && len(t.Attr) == 0 {
					whatsNewDataFound = true
				}
			default:
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			if whatsNewDataFound {
				whatsNew = acTz.Token().Data
				break mainLoop
				//whatsNewDataFound = false
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break mainLoop
		}
	}

	if wnOuterContainerFound && whatsNew == "" {
		return "", fmt.Errorf("the whats new container is found, however we couldn't extract the data from it")
	} else if whatsNew != "" {
		return whatsNew, nil
	} else {
		return "", nil
	}
}

// TODO, extractLastUpdated
// TODO, extractSize
// TODO, extractNumOfDownloads
// TODO, extractAppVer
// TODO, extractMinAndroidVer
// TODO, extractAgeRating

func handleDefaultStartTagToken(t html.Token) {
	switch t.Data {
	case "div":
		fallthrough
	case "h1":
		fallthrough
	case "span":
		fallthrough
	case "a":
		fallthrough
	case "button":
		fallthrough
	case "c-wiz":
		fallthrough
	case "content":
		// count only the tags that we actually open.
		openedTags++
	}
}

func handleDefaultEndTagToken(t html.Token) {
	switch t.Data {
	case "div":
		fallthrough
	case "h1":
		fallthrough
	case "span":
		fallthrough
	case "a":
		fallthrough
	case "button":
		fallthrough
	case "c-wiz":
		fallthrough
	case "content":
		// count only the tags that we actually open.
		openedTags--
	}
}

func formatHiResImgUrl(rawUrl string) (string, error) {
	if len(rawUrl) == 0 {
		return "", fmt.Errorf("can't format image url of lenght zero")
	} else if len(rawUrl) < 12 {
		// to prevent short urls from causing the app to panic
		return "", fmt.Errorf("can't format image url of unknown structure")
	}

	var urlFields = strings.Fields(rawUrl)
	var url = urlFields[0]
	var urlLen = len(url)
	// icon image size query first index
	var iconImgSizeQFi = strings.Index(url, "=s")
	// screenshot image size string
	var ssImgSizeStr = url[urlLen-11:]

	if iconImgSizeQFi > 0 {
		// remove the resizing query so the url now gets the original image.
		var newUrl = url[:iconImgSizeQFi]
		return newUrl, nil
	} else if strings.Contains(ssImgSizeStr, "w") && strings.Contains(ssImgSizeStr, "h") {
		// TODO, explore any possible res editing for this pattern
		return url, nil
	}

	// unexpected url pattern, so return the raw url as it is
	return rawUrl, nil
}

package appPage

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

const (
	playStoreUrl     = "https://play.google.com"
	devPageUrlPrefix = "/store/apps/dev"
	catPageUrlPrefix = "/store/apps/category/"
)

var totalTagsCounter uint64 = 0
var openedTags int64 = 0

func ExtractIconURL(acTz *html.Tokenizer) ([]string, error) {
	var iconUrlLoRes, iconUrlHiRes string
	var iconUrls = []string{"", ""}
	var containerFound bool
	var containerTagNum int64

mainLoop:
	for {
		if openedTags < containerTagNum {
			// this means that we have exited the container without finding the target element,
			// so we will break the loop as this is an indication on a expected wrong structure from the parser.
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
				// skip non candidate tags to improve the speed
				if len(t.Attr) == 0 {
					continue mainLoop
				}

				for _, attr := range t.Attr {
					if attr.Key == "class" && attr.Val == "xSyT2c" {
						containerFound = true
						containerTagNum = openedTags
						continue mainLoop
					}
				}

			case "img":
				// <img> tags doesn't count against opened tags.
				if containerFound && openedTags == containerTagNum {
					/*// skip non candidate tags to improve the speed
					if len(t.Attr) < 5 {
						continue mainLoop
					}*/

					var tmpLoRes string
					var tmpHiRes string
					var signs int8
					for _, attr := range t.Attr {
						switch {
						case attr.Key == "class" && attr.Val == "T75of sHb2Xb":
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
					if signs == 4 {
						tmpHiRes, e := formatHiResImgUrl(tmpHiRes)
						if e != nil {
							return nil, e
						}
						iconUrlLoRes = tmpLoRes
						iconUrlHiRes = tmpHiRes
						break mainLoop
					}
				}
			default:
				// count even any tag that we might close in parsing.
				handleDefaultStartTagToken(t)
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break mainLoop
		}
	}

	if iconUrlLoRes == "" || iconUrlHiRes == "" {
		return nil, fmt.Errorf("couldn't extract the app icon urls")
	}

	iconUrls[0] = iconUrlLoRes
	iconUrls[1] = iconUrlHiRes
	return iconUrls, nil
}

func ExtractAppName(acTz *html.Tokenizer) (string, error) {
	var appName string
	var containerFound bool
	var innerContainerFound bool
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
			case "h1":
				openedTags++
				// skip non candidate tags to improve the speed
				if len(t.Attr) < 2 {
					continue mainLoop
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
					containerFound = true
					containerTagNum = openedTags
					continue mainLoop
				}
			case "span":
				openedTags++
				if containerFound && openedTags > containerTagNum {
					// this is to checks if we have entered the span body or not,
					// to prevent taking any wrong text instead of the app name.
					innerContainerFound = true
				}
			default:
				// count even any tag that we might close in parsing.
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			if innerContainerFound {
				t := acTz.Token()
				appName = t.Data
				break mainLoop
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break mainLoop
		}
	}

	if appName == "" {
		return "", fmt.Errorf("couldn't extract the app name")
	}

	return appName, nil
}

// (devURL, devName, error)
func ExtractDevInfo(acTz *html.Tokenizer) (string, string, error) {
	var devName, devURL string
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
			case "a":
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
					containerFound = true
					containerTagNum = openedTags
				}
			default:
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			if containerFound && containerTagNum <= openedTags {
				t := acTz.Token()
				devName = t.Data
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

	devURL, e := updateDevUrl(devURL)
	if e != nil {
		return "", "", e
	}

	return devURL, devName, nil
}

// (catUrl, catName, err)
func ExtractCategoryInfo(acTz *html.Tokenizer) (string, string, error) {
	var catUrl, catName string
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
					containerFound = true
					containerTagNum = openedTags
				}
			default:
				// count even any tag that we might close in parsing.
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			// we can do this here(check for equality), as the text should be in the same level.
			if containerFound && containerTagNum == openedTags {
				t := acTz.Token()
				catName = t.Data
				break mainLoop
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break mainLoop
		}
	}

	if catUrl == "" || catName == "" {
		return "", "", fmt.Errorf("couldn't find the category url or the category name")
	}

	catUrl, e := updateCatUrl(catUrl)
	if e != nil {
		return "", "", e
	}

	return catUrl, catName, nil
}

// (containsAds, inAppPurchase, error)
func ExtractInAppOffering(acTz *html.Tokenizer) (string, error) {
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
					case attr.Key == "class" && attr.Val == "bSIuKf":
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

func ExtractPrice(acTz *html.Tokenizer) (string, error) {
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

func ExtractMediaUrls(acTz *html.Tokenizer) ([]string, []string, error) {
	var startImgUrl, videoUrl string
	var videoURLs = []string{"", ""}
	var imagesURLs = make([]string, 0, 30)
	var mediaContainerFound bool
	var mediaContainerTagNum int64
	var videoImgContainerFound bool
	var videoImgContainerTagNum int64
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
					case attr.Key == "class" && attr.Val == "SgoUSc":
						// the container of all the media is found, containing the trailer video with its start screen,
						// and all the containers for the screenshots.
						mediaContainerFound = true
						mediaContainerTagNum = openedTags

						// break out of the attr loop, cause all these cases are mutually exclusive.
						break attrLoop
					case mediaContainerFound && attr.Key == "class" && attr.Val == "MSLVtf Q4vdJd":
						videoImgContainerFound = true
						videoImgContainerTagNum = openedTags
						break attrLoop
					case videoImgContainerFound && attr.Key == "class" && attr.Val == "TdqJUe":
						videoContainerFound = true
						videoContainerTagNum = openedTags
						break attrLoop
					}
				}
			case "img":
				if videoImgContainerFound && openedTags == videoImgContainerTagNum {
					// skip non candidate tags to improve the speed
					/*if len(t.Attr) < 2 {
						continue videoLoop
					}*/

					var tmp string
					var signs int8
					for _, attr := range t.Attr {
						switch {
						case attr.Key == "class" && attr.Val == "T75of DYfLw":
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
						case attr.Key == "class" && attr.Val == "MMZjL lgooh  ":
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
						case attr.Key == "class" && attr.Val == "Q4vdJd":
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
					case attr.Key == "class" && attr.Val == "T75of DYfLw":
						signs++
					case attr.Key == "itemprop" && attr.Val == "image":
						signs++
					}
				}
				if signs == 4 {
					highRes, e := formatHiResImgUrl(tmpSrcset)
					if e != nil {
						highRes = tmpSrcset
					}
					imagesURLs = append(imagesURLs, tmpSrc, highRes)
				}
			default:
				handleDefaultStartTagToken(t)
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break screenshotsLoop
		}
	}

	if mediaContainerFound {
		if videoImgContainerFound && (startImgUrl == "" || videoUrl == "") {
			return nil, nil, fmt.Errorf("the video container is found, " +
				"however, we couldn't extract the required urls")

		} else if !videoImgContainerFound {
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

func ExtractDescription(acTz *html.Tokenizer) (string, error) {
	/*var descContainerFound bool
	var descContainerTagNum int64
	var desc = ""*/
	// the second method related vars
	var descContainerFound2 bool
	var descContainerTagNum2 int64
	var desc2 = ""

mainLoop:
	for {
		/*if openedTags < descContainerTagNum {
			break mainLoop
		}*/
		if openedTags < descContainerTagNum2 {
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

				if len(t.Attr) < 1 {
					continue mainLoop
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
				continue mainLoop
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
				break mainLoop
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
			break mainLoop
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

func ExtractRatingInfo(acTz *html.Tokenizer) (string, string, error) {
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

func ExtractRatingHistogram(acTz *html.Tokenizer) (map[string]string, error) {
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
					case attr.Key == "style" && attr.Val != "":
						signs++
						tmpValue = attr.Val
					case attr.Key == "class" && strings.Contains(attr.Val, "L2o20d"):
						signs++
						/*case attr.Key == "title" && attr.Val != "":
						signs++
						tmpValue = attr.Val*/
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
	} else if histogramContainerFound && len(histogram) != 5 {
		return nil, fmt.Errorf("the histogram container is found, however we couldn't extract the histogram data")
	} else {
		return nil, fmt.Errorf("the histogram container is not found")
	}
}

func ExtractWhatsNew(acTz *html.Tokenizer) (string, error) {
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
					case attr.Key == "jsshadow":
						fallthrough
					case attr.Key == "jsdata" && attr.Val != "":
						fallthrough
					case attr.Key == "jsmodel" && attr.Val == "hc6Ubd":
						signs++
					}
				}
				if signs == 4 {
					if tmp == "eG38Ge" {
						wnOuterContainerFound = true
						wnOuterContainerTagNum = openedTags
					} else if tmp == "HEOg8" {
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
			case "span":
				openedTags++

				if wnInnerContainerFound && openedTags == wnInnerContainerTagNum+1 && len(t.Attr) == 1 {
					whatsNewDataFound = true
				}
			default:
				handleDefaultStartTagToken(t)
			}
		case html.TextToken:
			if whatsNewDataFound {
				whatsNew = acTz.Token().Data
				break mainLoop
				// whatsNewDataFound = false
			}
		case html.EndTagToken:
			handleDefaultEndTagToken(acTz.Token())
		case html.ErrorToken:
			break mainLoop
		}
	}

	if wnInnerContainerFound && whatsNew == "" {
		return "", fmt.Errorf("the whats new container is found, but we couldn't extract the data from it")
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
		return "", fmt.Errorf("can't format image url of length zero")
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
		// TODO, explore any possible resolution changing for this pattern
		return url, nil
	}

	// unexpected url pattern, so return the raw url as it is
	return rawUrl, nil
}

func updateDevUrl(oldUrl string) (newUrl string, err error) {
	if strings.HasPrefix(oldUrl, devPageUrlPrefix) {
		newUrl = playStoreUrl + oldUrl
	} else {
		return "", fmt.Errorf("unexpected dev url")
	}
	return
}

func updateCatUrl(oldUrl string) (newUrl string, err error) {
	if strings.HasPrefix(oldUrl, catPageUrlPrefix) {
		newUrl = playStoreUrl + oldUrl
	} else {
		return "", fmt.Errorf("unexpected category url")
	}
	return
}

package apps

import (
	"fmt"
	"net/http"

	"github.com/asmsh/go-playstore-scraper/api/apps/fields"
	"github.com/asmsh/go-playstore-scraper/api/apps/internal/appPage"
	"github.com/asmsh/go-playstore-scraper/api/apps/internal/urls"
	"github.com/asmsh/go-playstore-scraper/api/apps/internal/validator"
	"golang.org/x/net/html"
)

func requestAppPage(url string) (*http.Response, error) {
	resp, e := http.Get(url)
	if e != nil {
		return nil, fmt.Errorf("error with requesting the url: \n" + e.Error())
	}

	// TODO, feat: handle the possible responses
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("error with requesting the url: \n"+"response status is: %s", resp.Status)
	}

	return resp, nil
}

func parseAppPage(appUrl *urls.AppUrl, appFields []fields.AppField) (*AppInfo, error) {
	var err error

	appFields, err = validator.ValidateAppFields(appFields)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the app info with error: %s", err.Error())
	}

	urlStr := appUrl.String()

	app := new(AppInfo)
	app.Country = appUrl.Country()
	app.Language = appUrl.Language()
	if appFields[0] == fields.AppId {
		app.AppId = appUrl.AppId()
		appFields = appFields[1:]
	}
	if len(appFields) > 0 && appFields[0] == fields.AppUrl {
		app.AppUrl = urlStr
		appFields = appFields[1:]
	}

	if len(appFields) == 0 {
		return app, nil
	}

	resp, err := requestAppPage(urlStr)
	if err != nil {
		return nil, fmt.Errorf("error with parsing the app page: " + err.Error())
	}
	defer resp.Body.Close()

	tz := html.NewTokenizer(resp.Body)
	err = parseAppContent(tz, app, appFields)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the app info with error: %s", err.Error())
	}

	return app, nil

}

func parseAppContent(acTz *html.Tokenizer, app *AppInfo, appFields []fields.AppField) error {
	for idx, currField := range appFields {
		var prevField, nextField fields.AppField

		// previous field
		if idx-1 >= 0 {
			prevField = appFields[idx-1]
		}
		// next field
		if idx+1 < len(appFields) {
			nextField = appFields[idx+1]
		}

		switch {
		case currField == fields.IconUrls:
			iconUrls, e := appPage.ExtractIconURL(acTz)
			if e != nil {
				return fmt.Errorf("error getting the icon url: " + e.Error())
			}
			app.IconUrls = iconUrls

		case currField == fields.AppName:
			appName, e := appPage.ExtractAppName(acTz)
			if e != nil {
				return fmt.Errorf("error getting the app name: " + e.Error())
			}
			app.AppName = appName

		case currField == fields.DevInfo:
			devUrl, devName, e := appPage.ExtractDevInfo(acTz)
			if e != nil {
				return fmt.Errorf("error getting the developer info: " + e.Error())
			}

			app.DevName = devName
			app.DevPageUrl = devUrl

		case currField == fields.Category:
			_, catName, e := appPage.ExtractCategoryInfo(acTz)
			if e != nil {
				return fmt.Errorf("error getting the category info: " + e.Error())
			}
			app.Category = catName

		case currField == fields.InAppExperience:
			offeringString, e := appPage.ExtractInAppOffering(acTz)
			if e != nil {
				return fmt.Errorf("error getting the in app offering info: " + e.Error())
			}
			app.InAppExperience = offeringString

		case currField == fields.Price:
			price, e := appPage.ExtractPrice(acTz)
			if e != nil {
				return fmt.Errorf("error getting the app price: " + e.Error())
			}
			app.Price = price

		case currField == fields.VideoTrailerUrls || currField == fields.ScreenShotsUrls:
			if prevField != fields.VideoTrailerUrls {
				videoUrls, imagesUrls, e := appPage.ExtractMediaUrls(acTz)
				if e != nil {
					return fmt.Errorf("error getting the app media urls: " + e.Error())
				}
				if currField == fields.VideoTrailerUrls {
					app.VideoTrailerUrls = videoUrls

					if nextField == fields.ScreenShotsUrls {
						app.ScreenShotsUrls = imagesUrls
					}
				}
				if currField == fields.ScreenShotsUrls {
					app.ScreenShotsUrls = imagesUrls
				}
			}

		case currField == fields.Description:
			appDesc, e := appPage.ExtractDescription(acTz)
			if e != nil {
				return fmt.Errorf("error getting the app's description: " + e.Error())
			}
			app.Description = appDesc

		case currField == fields.Rating || currField == fields.RatingCount:
			if prevField != fields.Rating {
				rating, ratingCount, e := appPage.ExtractRatingInfo(acTz)
				if e != nil {
					return fmt.Errorf("error getting the app's rating info: " + e.Error())
				}
				if currField == fields.Rating {
					app.Rating = rating

					if nextField == fields.RatingCount {
						app.RatingCount = ratingCount
					}
				}
				if currField == fields.RatingCount {
					app.RatingCount = ratingCount
				}
			}

		case currField == fields.RatingHistogram:
			histogram, e := appPage.ExtractRatingHistogram(acTz)
			if e != nil {
				return fmt.Errorf("error getting the app's histogram: " + e.Error())
			}
			app.RatingHistogram = histogram

		case currField == fields.WhatsNew:
			whatsNew, e := appPage.ExtractWhatsNew(acTz)
			if e != nil {
				return fmt.Errorf("error getting the app's whatsnew: " + e.Error())
			}
			app.WhatsNew = whatsNew
		}
	}

	return nil
}

package appField

import (
	"fmt"
	"sort"
)

// a structure represents one app info field (app name, dev name, ..)t
type AppField int

const (
	AllFields AppField = iota
	AppID
	AppUrl
	IconUrls
	AppName
	DevName
	DevPageUrl
	Category

	// TODO,
	//SecCategory

	InAppOffering

	// TODO,
	//ContainsAds
	//OffersInAppPurchases

	Price
	VideoTrailerUrls
	ScreenShotsUrls
	Description
	Rating
	RatingCount
	RatingHistogram
	WhatsNew

	// TODO,
	//LastUpdated
	//Size
	//NumOfDownloads
	//AppVer
	//MinAndroidVer
	//AgeRating
)

// ValidateAppFields checks if the provided AppField array contains only valid instances,
// and if not it will return an error.
// this function never returns an empty array (except with an non-nil error).
func ValidateAppFields(fields []AppField) ([]AppField, error) {
	// if no fields are provided, then retrieve all the fields.
	if len(fields) == 0 {
		return valuesArr[1:], nil
	}

	sort.Stable(appFields(fields))

	var retFields = []AppField{fields[0]}
	for _, f := range fields {
		if !isValidAppField(f) {
			return nil, fmt.Errorf("error validating the required app fields: unsupported app field")
		}

		// remove duplicate fields
		if f != retFields[len(retFields)-1] {
			retFields = append(retFields, f)
		}
	}

	// if the AllFields value is provided, then populate the returned array with all the fields.
	if fields[0] == AllFields {
		return valuesArr[1:], nil
	}

	return retFields, nil
}

// internal type used for implementing sorting
type appFields []AppField

func (f appFields) Len() int           { return len(f) }
func (f appFields) Less(i, j int) bool { return f[i] < f[j] }
func (f appFields) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

var valuesArr = []AppField{
	AllFields, AppID, AppUrl, IconUrls, AppName, DevName, DevPageUrl, Category, InAppOffering, Price,
	VideoTrailerUrls, ScreenShotsUrls, Description, Rating, RatingCount, RatingHistogram, WhatsNew,
}

func isValidAppField(f AppField) bool {
	for _, v := range valuesArr {
		if v == f {
			return true
		}
	}
	return false
}

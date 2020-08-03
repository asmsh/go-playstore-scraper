package validator

import (
	"fmt"
	"sort"

	"github.com/asmsh/go-playstore-scraper/api/apps/fields"
)

// ValidateAppFields checks if the provided AppInfoField array contains only valid instances,
// and if not it will return an error.
// this function never returns an empty array (except with an non-nil error).
func ValidateAppFields(appFields []fields.AppField) ([]fields.AppField, error) {
	// if no fields are provided, then retrieve all the fields.
	if len(appFields) == 0 {
		return allFields[1:], nil
	}

	sort.Stable(sortableAppFields(appFields))

	var retFields = make([]fields.AppField, 0, len(appFields))
	retFields = append(retFields, appFields[0])
	for idx := range appFields {
		// validate each field
		if !isValidAppInfoField(appFields[idx]) {
			return nil, fmt.Errorf("unsupported app field")
		}

		// remove duplicate fields
		if appFields[idx] != retFields[len(retFields)-1] {
			retFields = append(retFields, appFields[idx])
		}
	}

	// if the AllFields value is provided, then populate the returned array with all the fields.
	if appFields[0] == fields.AllFields {
		return allFields[1:], nil
	}

	return retFields, nil
}

func isValidAppInfoField(field fields.AppField) bool {
	if field < 0 {
		return false
	}
	if int(field) >= len(allFields) { // cause each field corresponds to a number, starting from 0
		return false
	}
	return true
}

// internal type used for implementing sorting of the app fields
type sortableAppFields []fields.AppField

func (f sortableAppFields) Len() int           { return len(f) }
func (f sortableAppFields) Less(i, j int) bool { return f[i] < f[j] }
func (f sortableAppFields) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

var allFields = []fields.AppField{
	fields.AllFields, fields.AppUrl, fields.AppId, fields.IconUrls,
	fields.AppName, fields.DevInfo, fields.Category, fields.FamilyCategory,
	fields.InAppExperience, fields.Price, fields.VideoTrailerUrls,
	fields.ScreenShotsUrls, fields.Description, fields.Rating,
	fields.RatingCount, fields.RatingHistogram, fields.WhatsNew,
	// FIXME, the following 6 fields aren't extracted (issue #1, #2, #3, #4, #5, #6)
	fields.LastUpdated, fields.Size, fields.NumOfDownloads,
	fields.Version, fields.MinAndroidVer, fields.AgeRating,
}

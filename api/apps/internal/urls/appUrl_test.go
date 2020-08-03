package urls_test

import (
	"testing"

	"github.com/asmsh/go-playstore-scraper/api/apps/internal/urls"
	"github.com/asmsh/go-playstore-scraper/locales"
)

type _appUrlArgs struct {
	appId    string
	country  locales.Country
	language locales.Language
}

type _appUrlTest struct {
	name    string
	wantErr bool
	args    _appUrlArgs
	wantUrl string
}

var _appUrlTests = []_appUrlTest{
	{
		name:    "test 1",
		wantErr: false,
		args: _appUrlArgs{
			appId:    "com",
			country:  locales.US,
			language: locales.English,
		},
		wantUrl: "https://play.google.com/store/apps/details?id=com&gl=us&hl=en",
	}, {
		name:    "test 2",
		wantErr: false,
		args: _appUrlArgs{
			appId:    "com.com",
			country:  locales.US,
			language: locales.DefLanguage,
		},
		wantUrl: "https://play.google.com/store/apps/details?id=com.com&gl=us",
	}, {
		name:    "test 3",
		wantErr: false,
		args: _appUrlArgs{
			appId:    "com.com",
			country:  locales.DefCountry,
			language: locales.DefLanguage,
		},
		wantUrl: "https://play.google.com/store/apps/details?id=com.com",
	},
}

func TestNewAppUrl(t *testing.T) {
	for _, tt := range _appUrlTests {
		t.Run(tt.name, func(t *testing.T) {
			appUrl, err := urls.NewAppUrl(tt.args.appId, tt.args.country, tt.args.language)
			if (err != nil) != tt.wantErr || appUrl == nil {
				t.Errorf("NewAppUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotId := appUrl.AppId(); gotId != tt.args.appId {
				t.Errorf("AppId() = %v, want %v", gotId, tt.args.appId)
			}
			if gotCo := appUrl.Country(); gotCo != tt.args.country {
				t.Errorf("Country() = %v, want %v", gotCo, tt.args.country)
			}
			if gotLa := appUrl.Language(); gotLa != tt.args.language {
				t.Errorf("Language() = %v, want %v", gotLa, tt.args.language)
			}

			if gotUrl := appUrl.String(); gotUrl != tt.wantUrl {
				t.Errorf("String() = %v, want %v", gotUrl, tt.wantUrl)
			}
		})
	}
}

func TestNewAppUrlFrom(t *testing.T) {
	for _, tt := range _appUrlTests {
		t.Run(tt.name, func(t *testing.T) {
			appUrl, err := urls.NewAppUrlFrom(tt.wantUrl)
			if (err != nil) != tt.wantErr || appUrl == nil {
				t.Errorf("NewAppUrlFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotId := appUrl.AppId(); gotId != tt.args.appId {
				t.Errorf("AppId() = %v, want %v", gotId, tt.args.appId)
			}
			if gotCo := appUrl.Country(); gotCo != tt.args.country {
				t.Errorf("Country() = %v, want %v", gotCo, tt.args.country)
			}
			if gotLa := appUrl.Language(); gotLa != tt.args.language {
				t.Errorf("Language() = %v, want %v", gotLa, tt.args.language)
			}

			if gotUrl := appUrl.String(); gotUrl != tt.wantUrl {
				t.Errorf("String() = %v, want %v", gotUrl, tt.wantUrl)
			}
		})
	}
}

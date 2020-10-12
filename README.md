# go-playstore-scraper [![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/asmsh/go-playstore-scraper)](https://pkg.go.dev/mod/github.com/asmsh/go-playstore-scraper)

It provides an API to retrieve the info of any particular app in the play store.

### Installing
> go get "github.com/asmsh/go-playstore-scraper"

### Notes
Still in beta.  
Check the changes before upgrading.    
See issues for not supported app fields.

### Legal
This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by Google or any of its affiliates or subsidiaries. This is an independent and unofficial API. Use at your own risk.

### Examples

* Using the app id(package name), extract all the app info. 
```go
package main

import (
	"fmt"
	"github.com/asmsh/go-playstore-scraper/api/apps"
)

func main() {
	app, e := apps.GetAppById("com.google.android.youtube")
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Printf("screenshots urls: %q\n", app.ScreenShotsUrls)
	}
}
```

* Using the app id, extract all the app info in a specific locale, by passing a specific
country and a language.
```go
package main

import (
	"fmt"
	"github.com/asmsh/go-playstore-scraper/api/apps"
	"github.com/asmsh/go-playstore-scraper/locales"
)

func main() {
	app, e := apps.GetAppByIdAdv("com.google.android.youtube", locales.US, locales.Arabic)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Printf("app description: %q\n", app.Description)
	}
}
```

* Using the app id, extract only specific fields of the app info.
```go
package main

import (
	"fmt"
	"github.com/asmsh/go-playstore-scraper/api/apps"
	"github.com/asmsh/go-playstore-scraper/api/apps/fields"
)

func main() {
	app, e := apps.GetAppById("com.google.android.youtube", fields.Rating, fields.RatingCount)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Printf("app rating: %q\n", app.Rating)
		fmt.Printf("app rating count: %q\n", app.RatingCount)
	}
}
``` 

* Using the url of the app, extract all the app info .
```go
package main

import (
	"fmt"
	"github.com/asmsh/go-playstore-scraper/api/apps"
)

func main() {
	url := "https://play.google.com/store/apps/details?id=com.google.android.youtube"
	app, e := apps.GetAppByUrl(url)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Printf("developer name: %q\n", app.DevName)
	}
}
``` 

* Using the url of the app, extract only specific fields of the app info.
```go
package main

import (
	"fmt"
	"github.com/asmsh/go-playstore-scraper/api/apps"
	"github.com/asmsh/go-playstore-scraper/api/apps/fields"
)

func main() {
	url := "https://play.google.com/store/apps/details?id=com.google.android.youtube"
	app, e := apps.GetAppByUrl(url, fields.InAppExperience)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Printf("in-app experience: %q\n", app.InAppExperience)
	}
}
``` 

* Using the url of the app, extract all the app info in a specific locale, by passing a specific 
country and a language.
```go
package main

import (
	"fmt"
	"github.com/asmsh/go-playstore-scraper/api/apps"
	"github.com/asmsh/go-playstore-scraper/locales"
)

func main() {
	url := "https://play.google.com/store/apps/details?id=com.google.android.youtube"
	app, e := apps.GetAppByUrlAdv(url, locales.US, locales.Arabic)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Printf("app category: %q\n", app.Category)
	}
}
```

##### more features are yet to come...

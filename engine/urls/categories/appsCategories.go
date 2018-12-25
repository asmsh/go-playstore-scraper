package categories

// Google Play Store apps main categories

// these will be used after the AppsStoreBaseURL constant

type Category string

// these consts will deal with building store urls that doesn't have specific category (like Top New Free Android Apps)
const (
	AllAppsAndGames Category = ""
	NoCategory      Category = ""
)

/*Apps categories constants*/
const (
	AndroidWearCategory       Category = "/category/ANDROID_WEAR"        //Android Wear
	ArtAndDesignCategory      Category = "/category/ART_AND_DESIGN"      //Art & Design
	AutoAndVehiclesCategory   Category = "/category/AUTO_AND_VEHICLES"   //Auto & Vehicles
	BeautyCategory            Category = "/category/BEAUTY"              //Beauty
	BooksAndReferenceCategory Category = "/category/BOOKS_AND_REFERENCE" //Books & Reference
	BusinessCategory          Category = "/category/BUSINESS"            //Business
	ComicsCategory            Category = "/category/COMICS"              //Comics
	CommunicationCategory     Category = "/category/COMMUNICATION"       //Communication
	DatingCategory            Category = "/category/DATING"              //Dating
	EducationCategory         Category = "/category/EDUCATION"           //Education
	EntertainmentCategory     Category = "/category/ENTERTAINMENT"       //Entertainment
	EventsCategory            Category = "/category/EVENTS"              //Events
	FinanceCategory           Category = "/category/FINANCE"             //Finance
	FoodAndDrinkCategory      Category = "/category/FOOD_AND_DRINK"      //Food & Drink
	HealthAndFitnessCategory  Category = "/category/HEALTH_AND_FITNESS"  //Health & Fitness
	HouseAndHomeCategory      Category = "/category/HOUSE_AND_HOME"      //House & Home
	LibrariesAndDemoCategory  Category = "/category/LIBRARIES_AND_DEMO"  //Libraries & Demo
	LifestyleCategory         Category = "/category/LIFESTYLE"           //Lifestyle
	MapsAndNavigationCategory Category = "/category/MAPS_AND_NAVIGATION" //Maps & Navigation
	MedicalCategory           Category = "/category/MEDICAL"             //Medical
	MusicAndAudioCategory     Category = "/category/MUSIC_AND_AUDIO"     //Music & Audio
	NewsAndMagazinesCategory  Category = "/category/NEWS_AND_MAGAZINES"  //News & Magazines
	ParentingCategory         Category = "/category/PARENTING"           //Parenting
	PersonalizationCategory   Category = "/category/PERSONALIZATION"     //Personalization
	PhotographyCategory       Category = "/category/PHOTOGRAPHY"         //Photography
	ProductivityCategory      Category = "/category/PRODUCTIVITY"        //Productivity
	ShoppingCategory          Category = "/category/SHOPPING"            //Shopping
	SocialCategory            Category = "/category/SOCIAL"              //Social
	SportsCategory            Category = "/category/SPORTS"              //Sports
	ToolsCategory             Category = "/category/TOOLS"               //Tools
	TravelAndLocalCategory    Category = "/category/TRAVEL_AND_LOCAL"    //Travel & Local
	VideoPlayersCategory      Category = "/category/VIDEO_PLAYERS"       //Video Players & Editors
	WeatherCategory           Category = "/category/WEATHER"             //Weather
)

var AppsCategories = []Category{
	AllAppsAndGames, NoCategory, AndroidWearCategory, ArtAndDesignCategory, AutoAndVehiclesCategory,
	BeautyCategory, BooksAndReferenceCategory, BusinessCategory, ComicsCategory, CommunicationCategory,
	DatingCategory, EducationCategory, EntertainmentCategory, EventsCategory, FinanceCategory, FoodAndDrinkCategory,
	HealthAndFitnessCategory, HouseAndHomeCategory, LibrariesAndDemoCategory, LifestyleCategory,
	MapsAndNavigationCategory, MedicalCategory, MusicAndAudioCategory, NewsAndMagazinesCategory, ParentingCategory,
	PersonalizationCategory, PhotographyCategory, ProductivityCategory, ShoppingCategory, SocialCategory, SportsCategory,
	ToolsCategory, TravelAndLocalCategory, VideoPlayersCategory, WeatherCategory,
}

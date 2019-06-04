package categories

// Google Play Store apps main categories
// these will be used after (AppsStoreBaseURL + categoryPath)

type Category string

// this constant is available in the builder and it's only here for reference
//const categoryPath   = "/category/"

// these constants will deal with building store urls that doesn't have specific category (like Top New Free Android Apps)
const (
	AllAppsAndGames Category = ""
	NoCategory      Category = ""
)

/*Apps categories constants*/
const (
	AndroidWearCategory       Category = "ANDROID_WEAR"        //Android Wear
	ArtAndDesignCategory      Category = "ART_AND_DESIGN"      //Art & Design
	AutoAndVehiclesCategory   Category = "AUTO_AND_VEHICLES"   //Auto & Vehicles
	BeautyCategory            Category = "BEAUTY"              //Beauty
	BooksAndReferenceCategory Category = "BOOKS_AND_REFERENCE" //Books & Reference
	BusinessCategory          Category = "BUSINESS"            //Business
	ComicsCategory            Category = "COMICS"              //Comics
	CommunicationCategory     Category = "COMMUNICATION"       //Communication
	DatingCategory            Category = "DATING"              //Dating
	EducationCategory         Category = "EDUCATION"           //Education
	EntertainmentCategory     Category = "ENTERTAINMENT"       //Entertainment
	EventsCategory            Category = "EVENTS"              //Events
	FinanceCategory           Category = "FINANCE"             //Finance
	FoodAndDrinkCategory      Category = "FOOD_AND_DRINK"      //Food & Drink
	HealthAndFitnessCategory  Category = "HEALTH_AND_FITNESS"  //Health & Fitness
	HouseAndHomeCategory      Category = "HOUSE_AND_HOME"      //House & Home
	LibrariesAndDemoCategory  Category = "LIBRARIES_AND_DEMO"  //Libraries & Demo
	LifestyleCategory         Category = "LIFESTYLE"           //Lifestyle
	MapsAndNavigationCategory Category = "MAPS_AND_NAVIGATION" //Maps & Navigation
	MedicalCategory           Category = "MEDICAL"             //Medical
	MusicAndAudioCategory     Category = "MUSIC_AND_AUDIO"     //Music & Audio
	NewsAndMagazinesCategory  Category = "NEWS_AND_MAGAZINES"  //News & Magazines
	ParentingCategory         Category = "PARENTING"           //Parenting
	PersonalizationCategory   Category = "PERSONALIZATION"     //Personalization
	PhotographyCategory       Category = "PHOTOGRAPHY"         //Photography
	ProductivityCategory      Category = "PRODUCTIVITY"        //Productivity
	ShoppingCategory          Category = "SHOPPING"            //Shopping
	SocialCategory            Category = "SOCIAL"              //Social
	SportsCategory            Category = "SPORTS"              //Sports
	ToolsCategory             Category = "TOOLS"               //Tools
	TravelAndLocalCategory    Category = "TRAVEL_AND_LOCAL"    //Travel & Local
	VideoPlayersCategory      Category = "VIDEO_PLAYERS"       //Video Players & Editors
	WeatherCategory           Category = "WEATHER"             //Weather
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

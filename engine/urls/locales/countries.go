package locales

type Country string

// this is used only when we want to left the country out of the url being built.
const CountryNone Country = ""

// almost all countries'code
const (
	CountryAlbania              Country = "dz"
	CountryAlgeria              Country = "dz"
	CountryAngola               Country = "ao"
	CountryAntiguaAndBarbuda    Country = "ag"
	CountryArgentina            Country = "ar"
	CountryArmenia              Country = "am"
	CountryAruba                Country = "aw"
	CountryAustralia            Country = "au"
	CountryAustria              Country = "at"
	CountryAzerbaijan           Country = "az"
	CountryBahamas              Country = "bs"
	CountryBahrain              Country = "bh"
	CountryBangladesh           Country = "bd"
	CountryBelarus              Country = "by"
	CountryBelgium              Country = "be"
	CountryBelize               Country = "bz"
	CountryBenin                Country = "bj"
	CountryBolivia              Country = "bo"
	CountryBosniaAndHerzegovina Country = "ba"
	CountryBotswana             Country = "bw"
	CountryBrazil               Country = "br"
	CountryBulgaria             Country = "bg"
	CountryBurkina              Country = "bf"
	CountryCambodia             Country = "kh"
	CountryCameroon             Country = "cm"
	CountryCanada               Country = "ca"
	CountryCapeVerde            Country = "cv"
	CountryChile                Country = "cl"
	CountryColombia             Country = "co"
	CountryCostaRica            Country = "cr"
	CountryCoteDIvore           Country = "ci" // Cote d'Ivore
	CountryCroatia              Country = "hr"
	CountryCyprus               Country = "cy"
	CountryCzechRepublic        Country = "cz"
	CountryDenmark              Country = "dk"
	CountryDominicanRepublic    Country = "do"
	CountryEcuador              Country = "ec"
	CountryEgypt                Country = "eg"
	CountryElSalvador           Country = "sv"
	CountryEstonia              Country = "ee"
	CountryFiji                 Country = "fj"
	CountryFinland              Country = "fi"
	CountryFrance               Country = "fr"
	CountryGabon                Country = "ga"
	CountryGermany              Country = "de"
	CountryGhana                Country = "gh"
	CountryGreece               Country = "gr"
	CountryGuatemala            Country = "gt"
	CountryGuineaBissau         Country = "gw"
	CountryHaiti                Country = "ht"
	CountryHonduras             Country = "hn"
	CountryHongKong             Country = "hk"
	CountryHungary              Country = "hu"
	CountryIceland              Country = "is"
	CountryIndia                Country = "in"
	CountryIndonesia            Country = "id"
	CountryIreland              Country = "ie"
	CountryItaly                Country = "it"
	CountryJamaica              Country = "jm"
	CountryJapan                Country = "jb"
	CountryJordan               Country = "jo"
	CountryKazakhstan           Country = "kz"
	CountryKenya                Country = "ke"
	CountryKuwait               Country = "kw"
	CountryKyrgyzstan           Country = "kg"
	CountryLaos                 Country = "la"
	CountryLatvia               Country = "lv"
	CountryLebanon              Country = "lb"
	CountryLiechtenstein        Country = "li"
	CountryLithuania            Country = "lt"
	CountryLuxembourg           Country = "lu"
	CountryMacedonia            Country = "mk"
	CountryMalaysia             Country = "my"
	CountryMali                 Country = "ml"
	CountryMalta                Country = "mt"
	CountryMauritius            Country = "mu"
	CountryMexico               Country = "mx"
	CountryMoldova              Country = "md"
	CountryMorocco              Country = "ma"
	CountryMozambique           Country = "mz"
	CountryNamibia              Country = "na"
	CountryNepal                Country = "np"
	CountryNetherlands          Country = "nl"
	CountryNetherlandsAntilles  Country = "an"
	CountryNewZealand           Country = "nz"
	CountryNicaragua            Country = "ni"
	CountryNiger                Country = "ne"
	CountryNigeria              Country = "ng"
	CountryNorway               Country = "no"
	CountryOman                 Country = "om"
	CountryPakistan             Country = "pk"
	CountryPanama               Country = "pa"
	CountryPapuaNewGuinea       Country = "pg"
	CountryParaguay             Country = "py"
	CountryPeru                 Country = "pe"
	CountryPhilippines          Country = "ph"
	CountryPoland               Country = "pl"
	CountryPortugal             Country = "pt"
	CountryQatar                Country = "qa"
	CountryRomania              Country = "ro"
	CountryRussia               Country = "ru"
	CountryRwanda               Country = "rw"
	CountrySaudiArabia          Country = "sa"
	CountrySenegal              Country = "sn"
	CountrySerbia               Country = "rs"
	CountrySingapore            Country = "sg"
	CountrySlovakia             Country = "sk"
	CountrySlovenia             Country = "si"
	CountrySouthAfrica          Country = "za"
	CountrySouthKorea           Country = "kr"
	CountrySpain                Country = "es"
	CountrySriLanka             Country = "lk"
	CountrySweden               Country = "se"
	CountrySwitzerland          Country = "ch"
	CountryTaiwan               Country = "tw"
	CountryTajikistan           Country = "tj"
	CountryTanzania             Country = "tz"
	CountryThailand             Country = "th"
	CountryTogo                 Country = "tg"
	CountryTrinidadAndTobago    Country = "tt"
	CountryTunisia              Country = "tn"
	CountryTurkey               Country = "tr"
	CountryTurkmenistan         Country = "tm"
	CountryUganda               Country = "ug"
	CountryUkraine              Country = "ua"
	CountryUAE                  Country = "ae" // United Arab Emirates
	CountryUK                   Country = "gb" // United Kingdom
	CountryUS                   Country = "us" // United States (including Puerto Rico, American Samoa, Guam, Marshall Islands, Northern Mariana Islands, Palau and US Virgin Islands)
	CountryUruguay              Country = "uy"
	CountryUzbekistan           Country = "uz"
	CountryVenezuela            Country = "ve"
	CountryVietnam              Country = "vn"
	CountryYemen                Country = "ye"
	CountryZambia               Country = "zm"
	CountryZimbabwe             Country = "zw"
)

// Countries an array containing all the above code
var Countries = []Country{
	CountryNone, CountryAlbania, CountryAlgeria, CountryAngola, CountryAntiguaAndBarbuda, CountryArgentina, CountryArmenia,
	CountryAruba, CountryAustralia, CountryAustria, CountryAzerbaijan, CountryBahamas, CountryBahrain, CountryBangladesh,
	CountryBelarus, CountryBelgium, CountryBelize, CountryBenin, CountryBolivia, CountryBosniaAndHerzegovina, CountryBotswana,
	CountryBrazil, CountryBulgaria, CountryBurkina, CountryCambodia, CountryCameroon, CountryCanada, CountryCapeVerde, CountryChile,
	CountryColombia, CountryCostaRica, CountryCoteDIvore, CountryCroatia, CountryCyprus, CountryCzechRepublic,
	CountryDenmark, CountryDominicanRepublic, CountryEcuador, CountryEgypt, CountryElSalvador, CountryEstonia, CountryFiji,
	CountryFinland, CountryFrance, CountryGabon, CountryGermany, CountryGhana, CountryGreece, CountryGuatemala,
	CountryGuineaBissau, CountryHaiti, CountryHonduras, CountryHongKong, CountryHungary, CountryIceland, CountryIndia,
	CountryIndonesia, CountryIreland, CountryItaly, CountryJamaica, CountryJapan, CountryJordan, CountryKazakhstan,
	CountryKenya, CountryKuwait, CountryKyrgyzstan, CountryLaos, CountryLatvia, CountryLebanon, CountryLiechtenstein,
	CountryLithuania, CountryLuxembourg, CountryMacedonia, CountryMalaysia, CountryMali, CountryMalta, CountryMauritius,
	CountryMexico, CountryMoldova, CountryMorocco, CountryMozambique, CountryNamibia, CountryNepal, CountryNetherlands,
	CountryNetherlandsAntilles, CountryNewZealand, CountryNicaragua, CountryNiger, CountryNigeria, CountryNorway, CountryOman,
	CountryPakistan, CountryPanama, CountryPapuaNewGuinea, CountryParaguay, CountryPeru, CountryPhilippines, CountryPoland,
	CountryPortugal, CountryQatar, CountryRomania, CountryRussia, CountryRwanda, CountrySaudiArabia, CountrySenegal,
	CountrySerbia, CountrySingapore, CountrySlovakia, CountrySlovenia, CountrySouthAfrica, CountrySouthKorea, CountrySpain,
	CountrySriLanka, CountrySweden, CountrySwitzerland, CountryTaiwan, CountryTajikistan, CountryTanzania, CountryThailand,
	CountryTogo, CountryTrinidadAndTobago, CountryTunisia, CountryTurkey, CountryTurkmenistan, CountryUAE, CountryUganda,
	CountryUK, CountryUkraine, CountryUruguay, CountryUS, CountryUzbekistan, CountryVenezuela, CountryVietnam, CountryYemen,
	CountryZambia, CountryZimbabwe,
}

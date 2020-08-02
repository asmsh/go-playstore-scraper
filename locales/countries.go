package locales

// Country is the country code which the playstore supports.
type Country string

const DefCountry Country = ""

const (
	Albania              Country = "al"
	Algeria              Country = "dz"
	Angola               Country = "ao"
	AntiguaAndBarbuda    Country = "ag"
	Argentina            Country = "ar"
	Armenia              Country = "am"
	Aruba                Country = "aw"
	Australia            Country = "au"
	Austria              Country = "at"
	Azerbaijan           Country = "az"
	Bahamas              Country = "bs"
	Bahrain              Country = "bh"
	Bangladesh           Country = "bd"
	Belarus              Country = "by"
	Belgium              Country = "be"
	Belize               Country = "bz"
	Benin                Country = "bj"
	Bolivia              Country = "bo"
	BosniaAndHerzegovina Country = "ba"
	Botswana             Country = "bw"
	Brazil               Country = "br"
	Bulgaria             Country = "bg"
	BurkinaFaso          Country = "bf"
	Cambodia             Country = "kh"
	Cameroon             Country = "cm"
	Canada               Country = "ca"
	CapeVerde            Country = "cv"
	Chile                Country = "cl"
	Colombia             Country = "co"
	CostaRica            Country = "cr"
	// CoteDIvoire is Cote d'Ivoire(Ivory Coast)
	CoteDIvoire         Country = "ci"
	Croatia             Country = "hr"
	Cyprus              Country = "cy"
	CzechRepublic       Country = "cz"
	Denmark             Country = "dk"
	DominicanRepublic   Country = "do"
	Ecuador             Country = "ec"
	Egypt               Country = "eg"
	ElSalvador          Country = "sv"
	Estonia             Country = "ee"
	Fiji                Country = "fj"
	Finland             Country = "fi"
	France              Country = "fr"
	Gabon               Country = "ga"
	Germany             Country = "de"
	Ghana               Country = "gh"
	Greece              Country = "gr"
	Guatemala           Country = "gt"
	GuineaBissau        Country = "gw"
	Haiti               Country = "ht"
	Honduras            Country = "hn"
	HongKong            Country = "hk"
	Hungary             Country = "hu"
	Iceland             Country = "is"
	India               Country = "in"
	Indonesia           Country = "id"
	Ireland             Country = "ie"
	Italy               Country = "it"
	Jamaica             Country = "jm"
	Japan               Country = "jb"
	Jordan              Country = "jo"
	Kazakhstan          Country = "kz"
	Kenya               Country = "ke"
	Kuwait              Country = "kw"
	Kyrgyzstan          Country = "kg"
	Laos                Country = "la"
	Latvia              Country = "lv"
	Lebanon             Country = "lb"
	Liechtenstein       Country = "li"
	Lithuania           Country = "lt"
	Luxembourg          Country = "lu"
	Macedonia           Country = "mk"
	Malaysia            Country = "my"
	Mali                Country = "ml"
	Malta               Country = "mt"
	Mauritius           Country = "mu"
	Mexico              Country = "mx"
	Moldova             Country = "md"
	Morocco             Country = "ma"
	Mozambique          Country = "mz"
	Namibia             Country = "na"
	Nepal               Country = "np"
	Netherlands         Country = "nl"
	NetherlandsAntilles Country = "an"
	NewZealand          Country = "nz"
	Nicaragua           Country = "ni"
	Niger               Country = "ne"
	Nigeria             Country = "ng"
	Norway              Country = "no"
	Oman                Country = "om"
	Pakistan            Country = "pk"
	Panama              Country = "pa"
	PapuaNewGuinea      Country = "pg"
	Paraguay            Country = "py"
	Peru                Country = "pe"
	Philippines         Country = "ph"
	Poland              Country = "pl"
	Portugal            Country = "pt"
	Qatar               Country = "qa"
	Romania             Country = "ro"
	Russia              Country = "ru"
	Rwanda              Country = "rw"
	SaudiArabia         Country = "sa"
	Senegal             Country = "sn"
	Serbia              Country = "rs"
	Singapore           Country = "sg"
	Slovakia            Country = "sk"
	Slovenia            Country = "si"
	SouthAfrica         Country = "za"
	SouthKorea          Country = "kr"
	Spain               Country = "es"
	SriLanka            Country = "lk"
	Sweden              Country = "se"
	Switzerland         Country = "ch"
	Taiwan              Country = "tw"
	Tajikistan          Country = "tj"
	Tanzania            Country = "tz"
	Thailand            Country = "th"
	Togo                Country = "tg"
	TrinidadAndTobago   Country = "tt"
	Tunisia             Country = "tn"
	Turkey              Country = "tr"
	Turkmenistan        Country = "tm"
	Uganda              Country = "ug"
	Ukraine             Country = "ua"
	// UAE is United Arab Emirates
	UAE Country = "ae"
	// UK is United Kingdom
	UK Country = "gb"
	// US is United States (including Puerto Rico, American Samoa, Guam,
	// Marshall Islands, Northern Mariana Islands, Palau and US Virgin Islands)
	US         Country = "us"
	Uruguay    Country = "uy"
	Uzbekistan Country = "uz"
	Venezuela  Country = "ve"
	Vietnam    Country = "vn"
	Yemen      Country = "ye"
	Zambia     Country = "zm"
	Zimbabwe   Country = "zw"
)

// AllCountries is an array of all the country codes available in this package
var AllCountries = [...]Country{
	DefCountry, Albania, Algeria, Angola,
	AntiguaAndBarbuda, Argentina, Armenia, Aruba,
	Australia, Austria, Azerbaijan, Bahamas,
	Bahrain, Bangladesh, Belarus, Belgium,
	Belize, Benin, Bolivia, BosniaAndHerzegovina,
	Botswana, Brazil, Bulgaria, BurkinaFaso,
	Cambodia, Cameroon, Canada, CapeVerde,
	Chile, Colombia, CostaRica, CoteDIvoire,
	Croatia, Cyprus, CzechRepublic, Denmark,
	DominicanRepublic, Ecuador, Egypt, ElSalvador,
	Estonia, Fiji, Finland, France, Gabon,
	Germany, Ghana, Greece, Guatemala,
	GuineaBissau, Haiti, Honduras, HongKong,
	Hungary, Iceland, India, Indonesia, Ireland,
	Italy, Jamaica, Japan, Jordan, Kazakhstan,
	Kenya, Kuwait, Kyrgyzstan, Laos, Latvia,
	Lebanon, Liechtenstein, Lithuania, Luxembourg,
	Macedonia, Malaysia, Mali, Malta, Mauritius,
	Mexico, Moldova, Morocco, Mozambique, Namibia,
	Nepal, Netherlands, NetherlandsAntilles, NewZealand,
	Nicaragua, Niger, Nigeria, Norway, Oman,
	Pakistan, Panama, PapuaNewGuinea, Paraguay,
	Peru, Philippines, Poland, Portugal, Qatar,
	Romania, Russia, Rwanda, SaudiArabia, Senegal,
	Serbia, Singapore, Slovakia, Slovenia, SouthAfrica,
	SouthKorea, Spain, SriLanka, Sweden, Switzerland,
	Taiwan, Tajikistan, Tanzania, Thailand, Togo,
	TrinidadAndTobago, Tunisia, Turkey, Turkmenistan,
	UAE, Uganda, UK, Ukraine, Uruguay,
	US, Uzbekistan, Venezuela, Vietnam, Yemen,
	Zambia, Zimbabwe,
}

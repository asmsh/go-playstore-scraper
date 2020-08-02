package locales

func IsValidCountry(co Country) bool {
	_, ok := allCounsMap[co]
	return ok
}

func IsValidLanguage(lang Language) bool {
	_, ok := allLangsMap[lang]
	return ok
}

var allCounsMap = map[Country]struct{}{
	DefCountry: {}, Albania: {}, Algeria: {}, Angola: {}, AntiguaAndBarbuda: {},
	Argentina: {}, Armenia: {}, Aruba: {}, Australia: {}, Austria: {}, Azerbaijan: {},
	Bahamas: {}, Bahrain: {}, Bangladesh: {}, Belarus: {}, Belgium: {}, Belize: {},
	Benin: {}, Bolivia: {}, BosniaAndHerzegovina: {}, Botswana: {}, Brazil: {},
	Bulgaria: {}, BurkinaFaso: {}, Cambodia: {}, Cameroon: {}, Canada: {}, CapeVerde: {},
	Chile: {}, Colombia: {}, CostaRica: {}, CoteDIvoire: {}, Croatia: {}, Cyprus: {},
	CzechRepublic: {}, Denmark: {}, DominicanRepublic: {}, Ecuador: {}, Egypt: {},
	ElSalvador: {}, Estonia: {}, Fiji: {}, Finland: {}, France: {}, Gabon: {},
	Germany: {}, Ghana: {}, Greece: {}, Guatemala: {}, GuineaBissau: {}, Haiti: {},
	Honduras: {}, HongKong: {}, Hungary: {}, Iceland: {}, India: {}, Indonesia: {},
	Ireland: {}, Italy: {}, Jamaica: {}, Japan: {}, Jordan: {}, Kazakhstan: {}, Kenya: {},
	Kuwait: {}, Kyrgyzstan: {}, Laos: {}, Latvia: {}, Lebanon: {}, Liechtenstein: {},
	Lithuania: {}, Luxembourg: {}, Macedonia: {}, Malaysia: {}, Mali: {}, Malta: {},
	Mauritius: {}, Mexico: {}, Moldova: {}, Morocco: {}, Mozambique: {}, Namibia: {},
	Nepal: {}, Netherlands: {}, NetherlandsAntilles: {}, NewZealand: {}, Nicaragua: {},
	Niger: {}, Nigeria: {}, Norway: {}, Oman: {}, Pakistan: {}, Panama: {}, PapuaNewGuinea: {},
	Paraguay: {}, Peru: {}, Philippines: {}, Poland: {}, Portugal: {}, Qatar: {},
	Romania: {}, Russia: {}, Rwanda: {}, SaudiArabia: {}, Senegal: {}, Serbia: {},
	Singapore: {}, Slovakia: {}, Slovenia: {}, SouthAfrica: {}, SouthKorea: {},
	Spain: {}, SriLanka: {}, Sweden: {}, Switzerland: {}, Taiwan: {}, Tajikistan: {},
	Tanzania: {}, Thailand: {}, Togo: {}, TrinidadAndTobago: {}, Tunisia: {},
	Turkey: {}, Turkmenistan: {}, UAE: {}, Uganda: {}, UK: {}, Ukraine: {}, Uruguay: {},
	US: {}, Uzbekistan: {}, Venezuela: {}, Vietnam: {}, Yemen: {}, Zambia: {}, Zimbabwe: {},
}

var allLangsMap = map[Language]struct{}{
	DefLanguage: {}, Afrikaans: {}, Amharic: {}, Arabic: {}, Azeri: {},
	Belarusian: {}, Bulgarian: {}, Bengali: {}, Bosnian: {}, Catalan: {},
	Czech: {}, Danish: {}, German: {}, GermanAustria: {}, GermanSwitzerland: {},
	Greek: {}, English: {}, EnglishAustralia: {}, EnglishCanada: {}, EnglishUK: {},
	EnglishIreland: {}, EnglishIndia: {}, EnglishSingapore: {}, EnglishUS: {},
	EnglishSouthAfrica: {}, Spanish: {}, SpanishLatinAmerica: {}, SpanishArgentina: {},
	SpanishBolivia: {}, SpanishChile: {}, SpanishColombia: {}, SpanishCostaRica: {},
	SpanishDominicanRepublic: {}, SpanishEcuador: {}, SpanishSpain: {}, SpanishGuatemala: {},
	SpanishHonduras: {}, SpanishMexico: {}, SpanishNicaragua: {}, SpanishPanama: {},
	SpanishPeru: {}, SpanishPuertoRico: {}, SpanishParaguay: {}, SpanishElSalvador: {},
	SpanishUS: {}, SpanishUruguay: {}, SpanishVenezuela: {}, Estonian: {}, Basque: {},
	Farsi: {}, Finnish: {}, Filipino: {}, French: {}, FrenchCanada: {}, FrenchFrance: {},
	FrenchSwitzerland: {}, Galician: {}, SwissGerman: {}, Gujarati: {}, Hebrew: {},
	Hindi: {}, Croatian: {}, Hungarian: {}, Armenian: {}, Indonesian: {}, Indonesian2: {},
	Icelandic: {}, Italian: {}, Japanese: {}, Georgian: {}, Kazakh: {}, Khmer: {},
	Kannada: {}, Korean: {}, Kyrgyz: {}, Lingala: {}, Lao: {}, Lithuanian: {},
	Latvian: {}, Macedonian: {}, Malayalam: {}, Mongolian: {}, Marathi: {}, Malay: {},
	Burmese: {}, NorwegianBokmal: {}, Nepali: {}, Dutch: {}, Norwegian: {}, Punjabi: {},
	Polish: {}, Portuguese: {}, PortugueseBrazil: {}, PortuguesePortugal: {}, Romanian: {},
	Russian: {}, Sinhala: {}, Slovak: {}, Slovenian: {}, Albanian: {}, Serbian: {},
	SerbianLatin: {}, Swedish: {}, Swahili: {}, Tamil: {}, Telugu: {}, Thai: {},
	Tagalog: {}, Turkish: {}, Ukrainian: {}, Urdu: {}, Uzbek: {}, Vietnamese: {},
	Chinese: {}, ChineseSimplified: {}, ChineseHongKong: {}, ChineseTraditional: {}, Zulu: {},
}

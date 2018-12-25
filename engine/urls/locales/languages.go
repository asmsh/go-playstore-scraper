package locales

type Language string

// this is used only when we want to left out the language from the url being built.
const LanguageNone Language = ""

const (
	LanguageAfrikaans                 Language = "af"
	LanguageAmharic                   Language = "am" // the same code as a foreign country, Armenia
	LanguageArabic                    Language = "ar" // the same code as a foreign country, Argentina
	LanguageAzeri                     Language = "az" // the same code as its country, Azerbaijan
	LanguageBelarusian                Language = "be" // the same code as a foreign country, Belgium
	LanguageBulgarian                 Language = "bg" // the same code as a foreign country, Bulgaria
	LanguageBengali                   Language = "bn"
	LanguageBosnian                   Language = "bs" // the same code as a foreign country, Bahamas
	LanguageCatalan                   Language = "ca" // the same code as a foreign country, Canada
	LanguageCzech                     Language = "cs"
	LanguageDanish                    Language = "da"
	LanguageGerman                    Language = "de" // the same code as its country, Germany
	LanguageGerman_Austria            Language = "de_AT"
	LanguageGerman_Switzerland        Language = "de_CH"
	LanguageGreek                     Language = "el"
	LanguageEnglish                   Language = "en"
	LanguageEnglish_Australia         Language = "en_AU"
	LanguageEnglish_Canada            Language = "en_CA"
	LanguageEnglish_UK                Language = "en_GB"
	LanguageEnglish_Ireland           Language = "en_IE"
	LanguageEnglish_India             Language = "en_IN"
	LanguageEnglish_Singapore         Language = "en_SG"
	LanguageEnglish_US                Language = "en_US"
	LanguageEnglish_SouthAfrica       Language = "en_ZA"
	LanguageSpanish                   Language = "es" // the same code as its country, Spain
	LanguageSpanish_LatinAmerica      Language = "es_419"
	LanguageSpanish_Argentina         Language = "es_AR"
	LanguageSpanish_Bolivia           Language = "es_BO"
	LanguageSpanish_Chile             Language = "es_CL"
	LanguageSpanish_Colombia          Language = "es_CO"
	LanguageSpanish_CostaRica         Language = "es_CR"
	LanguageSpanish_DominicanRepublic Language = "es_DO"
	LanguageSpanish_Ecuador           Language = "es_EC"
	LanguageSpanish_Spain             Language = "es-ES"
	LanguageSpanish_Guatemala         Language = "es_GT"
	LanguageSpanish_Honduras          Language = "es_HN"
	LanguageSpanish_Mexico            Language = "es_MX"
	LanguageSpanish_Nicaragua         Language = "es_NI"
	LanguageSpanish_Panama            Language = "es_PA"
	LanguageSpanish_Peru              Language = "es_PE"
	LanguageSpanish_PuertoRico        Language = "es_PR"
	LanguageSpanish_Paraguay          Language = "es_PY"
	LanguageSpanish_ElSalvador        Language = "es_SV"
	LanguageSpanish_US                Language = "es_US"
	LanguageSpanish_Uruguay           Language = "es_UY"
	LanguageSpanish_Venezuela         Language = "es_VE"
	LanguageEstonian                  Language = "et"
	LanguageBasque                    Language = "eu"
	LanguageFarsi                     Language = "fa"
	LanguageFinnish                   Language = "fi" // the same code as its country, Finland
	LanguageFilipino                  Language = "fil"
	LanguageFrench                    Language = "fr" // the same code as its country, France
	LanguageFrench_Canada             Language = "fr_CA"
	LanguageFrench_France             Language = "fr_FR"
	LanguageFrench_Switzerland        Language = "fr_CH"
	LanguageGalician                  Language = "gl"
	LanguageSwissGerman               Language = "gsw" // TODO, any better name?
	LanguageGujarati                  Language = "gu"
	LanguageHebrew                    Language = "he"
	LanguageHindi                     Language = "hi"
	LanguageCroatian                  Language = "hr" // the same code as its country, Croatia
	LanguageHungarian                 Language = "hu" // the same code as its country, Hungary
	LanguageArmenian                  Language = "hy"
	LanguageIndonesian                Language = "id" // the same code as its country, Indonesia
	LanguageIndonesian_2              Language = "in" // the same code as a foreign country, India
	LanguageIcelandic                 Language = "is" // the same code as its country, Iceland
	LanguageItalian                   Language = "it" // the same code as its country, Italy
	LanguageJapanese                  Language = "ja"
	LanguageGeorgian                  Language = "ka"
	LanguageKazakh                    Language = "kk"
	LanguageKhmer                     Language = "km"
	LanguageKannada                   Language = "kn"
	LanguageKorean                    Language = "ko"
	LanguageKyrgyz                    Language = "ky"
	LanguageLingala                   Language = "ln"
	LanguageLao                       Language = "lo"
	LanguageLithuanian                Language = "lt" // the same code as its country, Lithuania
	LanguageLatvian                   Language = "lv" // the same code as its country, Latvia
	LanguageMacedonian                Language = "mk" // the same code as its country, Macedonia
	LanguageMalayalam                 Language = "ml" // the same code as a foreign country, Mali
	LanguageMongolian                 Language = "mn"
	LanguageMarathi                   Language = "mr"
	LanguageMalay                     Language = "ms"
	LanguageBurmese                   Language = "my" // the same code as a foreign country, Malaysia
	LanguageNorwegian_Bokmal          Language = "nb"
	LanguageNepali                    Language = "ne" // the same code as a foreign country, Niger
	LanguageDutch                     Language = "nl" // the same code as its country, Netherlands
	LanguageNorwegian                 Language = "no" // the same code as its country, Norway
	LanguagePunjabi                   Language = "pa" // the same code as a foreign country, Panama
	LanguagePolish                    Language = "pl" // the same code as its country, Poland
	LanguagePortuguese                Language = "pt" // the same code as its country, Portugal
	LanguagePortuguese_Brazil         Language = "pt_BR"
	LanguagePortuguese_Portugal       Language = "pt_PT"
	LanguageRomanian                  Language = "ro" // the same code as its country, Romania
	LanguageRussian                   Language = "ru" // the same code as its country, Russia
	LanguageSinhala                   Language = "si" // the same code as a foreign country, Slovenia
	LanguageSlovak                    Language = "sk" // the same code as its country, Slovakia
	LanguageSlovenian                 Language = "sl"
	LanguageAlbanian                  Language = "sq"
	LanguageSerbian                   Language = "sr"
	LanguageSerbian_Latin             Language = "sr_LATN"
	LanguageSwedish                   Language = "sv" // the same code as a foreign country, ElSalvador
	LanguageSwahili                   Language = "sw"
	LanguageTamil                     Language = "ta"
	LanguageTelugu                    Language = "te"
	LanguageThai                      Language = "th" // the same code as its country, Thailand
	LanguageTagalog                   Language = "tl"
	LanguageTurkish                   Language = "tr" // the same code as its country, Turkey
	LanguageUkrainian                 Language = "uk"
	LanguageUrdu                      Language = "ur"
	LanguageUzbek                     Language = "uz" // the same code as its country, Uzbekistan
	LanguageVietnamese                Language = "vi"
	LanguageChinese                   Language = "zh"
	LanguageChinese_Simplified        Language = "zh_CN"
	LanguageChinese_HongKong          Language = "zh_HK"
	LanguageChinese_Traditional       Language = "zh_TW"
	LanguageZulu                      Language = "zu"
)

var Languages = []Language{
	LanguageNone, LanguageAfrikaans, LanguageAmharic, LanguageArabic, LanguageAzeri, LanguageBelarusian, LanguageBulgarian,
	LanguageBengali, LanguageBosnian, LanguageCatalan, LanguageCzech, LanguageDanish, LanguageGerman, LanguageGerman_Austria,
	LanguageGerman_Switzerland, LanguageGreek, LanguageEnglish, LanguageEnglish_Australia, LanguageEnglish_Canada,
	LanguageEnglish_UK, LanguageEnglish_Ireland, LanguageEnglish_India, LanguageEnglish_Singapore, LanguageEnglish_US,
	LanguageEnglish_SouthAfrica, LanguageSpanish, LanguageSpanish_LatinAmerica, LanguageSpanish_Argentina,
	LanguageSpanish_Bolivia, LanguageSpanish_Chile, LanguageSpanish_Colombia, LanguageSpanish_CostaRica,
	LanguageSpanish_DominicanRepublic, LanguageSpanish_Ecuador, LanguageSpanish_Spain, LanguageSpanish_Guatemala,
	LanguageSpanish_Honduras, LanguageSpanish_Mexico, LanguageSpanish_Nicaragua, LanguageSpanish_Panama, LanguageSpanish_Peru,
	LanguageSpanish_PuertoRico, LanguageSpanish_Paraguay, LanguageSpanish_ElSalvador, LanguageSpanish_US, LanguageSpanish_Uruguay,
	LanguageSpanish_Venezuela, LanguageEstonian, LanguageBasque, LanguageFarsi, LanguageFinnish, LanguageFilipino,
	LanguageFrench, LanguageFrench_Canada, LanguageFrench_France, LanguageFrench_Switzerland, LanguageGalician,
	LanguageSwissGerman, LanguageGujarati, LanguageHebrew, LanguageHindi, LanguageCroatian, LanguageHungarian, LanguageArmenian,
	LanguageIndonesian, LanguageIndonesian_2, LanguageIcelandic, LanguageItalian, LanguageJapanese, LanguageGeorgian,
	LanguageKazakh, LanguageKhmer, LanguageKannada, LanguageKorean, LanguageKyrgyz, LanguageLingala, LanguageLao,
	LanguageLithuanian, LanguageLatvian, LanguageMacedonian, LanguageMalayalam, LanguageMongolian, LanguageMarathi,
	LanguageMalay, LanguageBurmese, LanguageNorwegian_Bokmal, LanguageNepali, LanguageDutch, LanguageNorwegian,
	LanguagePunjabi, LanguagePolish, LanguagePortuguese, LanguagePortuguese_Brazil, LanguagePortuguese_Portugal,
	LanguageRomanian, LanguageRussian, LanguageSinhala, LanguageSlovak, LanguageSlovenian, LanguageAlbanian,
	LanguageSerbian, LanguageSerbian_Latin, LanguageSwedish, LanguageSwahili, LanguageTamil, LanguageTelugu, LanguageThai,
	LanguageTagalog, LanguageTurkish, LanguageUkrainian, LanguageUrdu, LanguageUzbek, LanguageVietnamese, LanguageChinese,
	LanguageChinese_Simplified, LanguageChinese_HongKong, LanguageChinese_Traditional, LanguageZulu,
}

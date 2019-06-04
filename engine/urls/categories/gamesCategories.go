package categories

// Google Play Store games main categories
// these will be used after (AppsStoreBaseURL + categoryPath)

// this constant is available in the builder and it's only here for reference
//const categoryPath   = "/category/"

/*Games categories constants*/
const (
	GamesDefaultHome        Category = "GAME"              //Home Page for Games(this isn't an actual category)
	GameActionCategory      Category = "GAME_ACTION"       //Action
	GameAdventureCategory   Category = "GAME_ADVENTURE"    //Adventure
	GameArcadeCategory      Category = "GAME_ARCADE"       //Arcade
	GameBoardCategory       Category = "GAME_BOARD"        //Board
	GameCardCategory        Category = "GAME_CARD"         //Card
	GameCasinoCategory      Category = "GAME_CASINO"       //Casino
	GameCasualCategory      Category = "GAME_CASUAL"       //Casual
	GameEducationalCategory Category = "GAME_EDUCATIONAL"  //Educational
	GameMusicCategory       Category = "GAME_MUSIC"        //Music
	GamePuzzleCategory      Category = "GAME_PUZZLE"       //Puzzle
	GameRacingCategory      Category = "GAME_RACING"       //Racing
	GameRolePlayingCategory Category = "GAME_ROLE_PLAYING" //Role Playing
	GameSimulationCategory  Category = "GAME_SIMULATION"   //Simulation
	GameSportsCategory      Category = "GAME_SPORTS"       //Sports
	GameStrategyCategory    Category = "GAME_STRATEGY"     //Strategy
	GameTriviaCategory      Category = "GAME_TRIVIA"       //Trivia
	GameWordCategory        Category = "GAME_WORD"         //Word
)

var GamesCategories = []Category{
	GamesDefaultHome, GameActionCategory, GameAdventureCategory, GameArcadeCategory, GameBoardCategory,
	GameCardCategory, GameCasinoCategory, GameCasualCategory, GameEducationalCategory, GameMusicCategory,
	GamePuzzleCategory, GameRacingCategory, GameRolePlayingCategory, GameSimulationCategory, GameSportsCategory,
	GameStrategyCategory, GameTriviaCategory, GameWordCategory,
}

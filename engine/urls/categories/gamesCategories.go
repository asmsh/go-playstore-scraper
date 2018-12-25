package categories

// Google Play Store games main categories
// these will be used after the AppsStoreBaseURL constant

/*Games categories constants*/
const (
	GamesDefaultHome        Category = "/category/GAME"              //Home Page for Games(this isn't an actual category)
	GameActionCategory      Category = "/category/GAME_ACTION"       //Action
	GameAdventureCategory   Category = "/category/GAME_ADVENTURE"    //Adventure
	GameArcadeCategory      Category = "/category/GAME_ARCADE"       //Arcade
	GameBoardCategory       Category = "/category/GAME_BOARD"        //Board
	GameCardCategory        Category = "/category/GAME_CARD"         //Card
	GameCasinoCategory      Category = "/category/GAME_CASINO"       //Casino
	GameCasualCategory      Category = "/category/GAME_CASUAL"       //Casual
	GameEducationalCategory Category = "/category/GAME_EDUCATIONAL"  //Educational
	GameMusicCategory       Category = "/category/GAME_MUSIC"        //Music
	GamePuzzleCategory      Category = "/category/GAME_PUZZLE"       //Puzzle
	GameRacingCategory      Category = "/category/GAME_RACING"       //Racing
	GameRolePlayingCategory Category = "/category/GAME_ROLE_PLAYING" //Role Playing
	GameSimulationCategory  Category = "/category/GAME_SIMULATION"   //Simulation
	GameSportsCategory      Category = "/category/GAME_SPORTS"       //Sports
	GameStrategyCategory    Category = "/category/GAME_STRATEGY"     //Strategy
	GameTriviaCategory      Category = "/category/GAME_TRIVIA"       //Trivia
	GameWordCategory        Category = "/category/GAME_WORD"         //Word
)

var GamesCategories = []Category{
	GamesDefaultHome, GameActionCategory, GameAdventureCategory, GameArcadeCategory, GameBoardCategory,
	GameCardCategory, GameCasinoCategory, GameCasualCategory, GameEducationalCategory, GameMusicCategory,
	GamePuzzleCategory, GameRacingCategory, GameRolePlayingCategory, GameSimulationCategory, GameSportsCategory,
	GameStrategyCategory, GameTriviaCategory, GameWordCategory,
}

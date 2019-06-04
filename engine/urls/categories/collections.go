package categories

// Collections subcategories for each main category
// these constants will be used after (AppsStoreBaseURL + categoryPath + AppCategory/GameCategory + collectionPath)

type Collection string

// these constants are available in the builder and they are only here for reference
//const categoryPath   = "/category/"
//const collectionPath = "/collection/"

/*Collections constants*/
const (
	TopPaidCollection     Collection = "topselling_paid"     //Top Paid
	TopFreeCollection     Collection = "topselling_free"     //Top Free
	TopGrossingCollection Collection = "topgrossing"         //Top Grossing
	TopNewPaidCollection  Collection = "topselling_new_paid" //Top New Paid
	TopNewFreeCollection  Collection = "topselling_new_free" //Top New Free
	TrendingCollection    Collection = "movers_shakers"      //Trending Apps
)

var Collections = []Collection{
	TopPaidCollection, TopFreeCollection, TopGrossingCollection,
	TopNewPaidCollection, TopNewFreeCollection, TrendingCollection,
}

// The only guaranteed collections are, TopPaidCollection and TopFreeCollection, these two are available in every category

package categories

// Collections subcategories for each main category
// these constants will be used after the (AppsStoreBaseURL + AppCategory/GameCategory)
type Collection string

/*Collections constants*/
const (
	TopPaidCollection     Collection = "/collection/topselling_paid"     //Top Paid
	TopFreeCollection     Collection = "/collection/topselling_free"     //Top Free
	TopGrossingCollection Collection = "/collection/topgrossing"         //Top Grossing
	TopNewPaidCollection  Collection = "/collection/topselling_new_paid" //Top New Paid
	TopNewFreeCollection  Collection = "/collection/topselling_new_free" //Top New Free
	TrendingCollection    Collection = "/collection/movers_shakers"      //Trending Apps
)

var Collections = []Collection{
	TopPaidCollection, TopFreeCollection, TopGrossingCollection,
	TopNewPaidCollection, TopNewFreeCollection, TrendingCollection,
}

// The only guaranteed collections are, TopPaidCollection and TopFreeCollection, these two are available in every category

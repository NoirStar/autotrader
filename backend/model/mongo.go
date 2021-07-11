package model

// User model
type User struct {
	ID       string  `bson:"id"`
	PW       string  `bson:"pw"`
	Email    string  `bson:"email"`
	Nickname string  `bson:"nickname"`
	Birth    string  `bson:"birth"`
	Level    int     `bson:"level"`
	Money    float64 `bson:"money"`
}

// Market model
type Market map[string]*MarketData

// MarketData model
type MarketData struct {
	AskCount float64 `bson:"ask_count,omitempty"`
	AskTotal float64 `bson:"ask_total,omitempty"`
	BidCount float64 `bson:"bid_count,omitempty"`
	BidTotal float64 `bson:"bid_total,omitempty"`
}

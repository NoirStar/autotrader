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
type Market struct {
	AskCount float64 `json:"ask_count"`
	AskTotal float64 `json:"ask_total"`
	BidCount float64 `json:"bid_count"`
	BidTotal float64 `json:"bid_total"`
}

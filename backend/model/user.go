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

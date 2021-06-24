package model

// User model
type User struct {
	ID       string  `bson:"id"`
	Password string  `bson:"password"`
	Email    string  `bson:"email"`
	Nickname string  `bson:"nickname"`
	Birth    string  `bson:"birth"`
	Sex      string  `bson:"sex"`
	Level    int     `bson:"level"`
	Money    float64 `bson:"money"`
}

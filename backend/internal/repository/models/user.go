package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	Login     string             `bson:"login"`
	Password  string             `bson:"password"`
	IsBlocked bool               `bson:"is_blocked"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Username     string             `bson:"username"`
	Login        string             `bson:"login"`
	Password     string             `bson:"password"`
	PasswordSalt string             `bson:"password_salt"`
	Preference   *UserPreference    `bson:"preference"`
}

type UserPreference struct {
	IsBlocked            bool `bson:"is_blocked"`
	IsPasswordConstraint bool `bson:"is_password_constraint"`
}

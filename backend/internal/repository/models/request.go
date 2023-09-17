package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetUserRequest struct {
	ID       primitive.ObjectID
	Login    string
	Password string
}

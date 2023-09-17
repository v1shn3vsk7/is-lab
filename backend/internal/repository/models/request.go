package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetUserRequest struct {
	ID       primitive.ObjectID
	Login    string
	Password string
}

type UpdateUserPasswordRequest struct {
	ID       primitive.ObjectID
	Password string
}

type UpdateUserRequest struct {
	Username             string
	Login                string
	Password             string
	IsBlocked            bool
	IsPasswordConstraint bool
}

package repository

import (
	"github.com/v1shn3vsk7/is-lab/internal/repository/models"
	"go.mongodb.org/mongo-driver/bson"
)

func getUserFilter(req *models.GetUserRequest) bson.M {
	if req.Login != "" && req.Password != "" {
		return bson.M{
			"login": bson.M{
				"$eq": req.Login,
			},
			"password": bson.M{
				"$eq": req.Password,
			},
		}
	}

	return bson.M{
		"_id": bson.M{
			"$eq": req.ID,
		},
	}

}

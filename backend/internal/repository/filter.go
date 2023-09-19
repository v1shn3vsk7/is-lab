package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/v1shn3vsk7/is-lab/internal/repository/models"
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

func FilterByID(id primitive.ObjectID) bson.M {
	return bson.M{
		"_id": bson.M{
			"$eq": id,
		},
	}
}

func UpdateField(field string, value interface{}) bson.M {
	return bson.M{
		"$set": bson.M{
			field: value,
		},
	}
}

// мне лень нормальный фильтр сделать
func UpdateUserFilter(req *models.UpdateUserRequest) bson.M {
	if req.IsBlocked != nil {
		return UpdateField("preference.is_blocked", req.IsBlocked)
	}

	return UpdateField("preference.is_password_constraint", req.IsPasswordConstraint)
}

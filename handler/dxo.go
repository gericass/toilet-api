package handler

import (
	"github.com/gericass/toilet-api/handler/response"
	"github.com/gericass/toilet-api/data/local"
	"github.com/gericass/toilet-api/data/remote"
	"database/sql"
)

func ConvertUser(user *local.User) *response.User {
	u := &response.User{
		Name:     user.Name,
		UID:      user.UID,
		IconPath: user.IconPath,
	}
	return u
}

func ConvertReviews(reviews []*local.Review, db *sql.DB) *response.Reviews {
	var rs []*response.Review
	for _, v := range reviews {
		user := &local.User{ID: v.UserId}
		user.FindUserById(db)
		r := &response.Review{
			ToiletId:  v.ToiletId,
			User:      ConvertUser(user),
			Valuation: v.Valuation,
			CreatedAt: v.CreatedAt,
		}
		rs = append(rs, r)
	}
	return &response.Reviews{Reviews: rs, Status: "OK"}
}

func ConvertToilets(toilets []*local.Toilet) *response.Toilets {
	var respToilets []*response.Toilet
	for _, v := range toilets {
		t := &response.Toilet{
			Name:        v.Name,
			GoogleId:    v.GoogleId,
			Geolocation: v.Geolocation,
			Image:       v.ImagePath,
			Description: v.Description,
			Valuation:   v.Valuation,
			UpdatedAt:   v.UpdatedAt,
		}
		respToilets = append(respToilets, t)
	}
	return &response.Toilets{Toilets: respToilets, Status: "OK"}
}

func ConvertPlaceDetailToToilet(pd *remote.PlaceDetail) *local.Toilet {
	t := &local.Toilet{
		Name:        pd.Result.Name,
		GoogleId:    pd.Result.PlaceID,
		Lat:         pd.Result.Geometry.Location.Lat,
		Lng:         pd.Result.Geometry.Location.Lng,
		Geolocation: pd.Result.FormattedAddress,
		ImagePath:   pd.Result.Icon,
	}
	return t
}

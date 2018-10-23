package handler

import (
	"github.com/gericass/toilet-api/handler/response"
	"github.com/gericass/toilet-api/data/local"
	"github.com/gericass/toilet-api/data/remote"
)

func ConvertReviews(reviews []*local.Review) []*response.Review {
	var rs []*response.Review
	for _, v := range reviews {
		r := &response.Review{
			ToiletId:  v.ToiletId,
			UserId:    v.UserId,
			Valuation: v.Valuation,
			CreatedAt: v.CreatedAt,
		}
		rs = append(rs, r)
	}
	return rs
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
	return &response.Toilets{Toilets: respToilets}
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

package handler

import (
	"github.com/labstack/echo"
	"github.com/gericass/toilet-api/util"
	"net/http"
	"github.com/gericass/toilet-api/data/local"
	"github.com/gericass/toilet-api/handler/response"
	"github.com/gericass/toilet-api/handler/request"
	"github.com/gericass/toilet-api/data/remote"
)

func convertToilets(toilets []*local.Toilet) *response.Toilets {
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

func convertPlaceDetailToToilet(pd *remote.PlaceDetail) *local.Toilet {
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

func GetFavoriteHandler(c echo.Context) error {
	cc := c.(*CustomContext)
	token, err := util.GetToken(c)
	if err != nil {
		return c.String(http.StatusForbidden, "Forbidden")
	}
	user := &local.User{UID: token.UID}
	userId, err := user.GetUserId(cc.DB)
	usersToilets := &local.UsersToilets{UserId: userId}
	usersToiletsList, err := usersToilets.FindToiletsByUserId(cc.DB)
	if err != nil {
		return err
	}
	var toilets []*local.Toilet
	for _, v := range usersToiletsList {
		t := &local.Toilet{ID: v.ToiletId}
		t.FindToiletByGoogleId(cc.DB)
		toilets = append(toilets, t)
	}
	resp := convertToilets(toilets)
	return c.JSON(http.StatusOK, resp)
}

func PostFavoriteHandler(c echo.Context) error {
	cc := c.(*CustomContext)
	requestFavorite := new(request.Favorite)
	err := c.Bind(requestFavorite)
	if err != nil {
		return err
	}
	toilet := &local.Toilet{GoogleId: requestFavorite.GoogleId}
	toiletExists, err := toilet.Exists(cc.DB)
	if err != nil {
		return err
	}
	if !toiletExists {
		placeDetail, err := remote.GetPlaceDetail(toilet.GoogleId)
		if err != nil {
			return err
		}
		t := convertPlaceDetailToToilet(placeDetail)
		err = t.Insert(cc.DB)
		if err != nil {
			return err
		}
	}

	err = toilet.FindToiletByGoogleId(cc.DB)
	if err != nil {
		return err
	}
	token, err := util.GetToken(c)
	if err != nil {
		return err
	}
	user := &local.User{UID: token.UID}
	userId, err := user.GetUserId(cc.DB)
	if err != nil {
		return err
	}
	usersToilets := &local.UsersToilets{ToiletId: toilet.ID, UserId: userId}
	exists, err := usersToilets.Exists(cc.DB)
	if err != nil {
		return err
	}
	if exists {
		return c.String(http.StatusConflict, "already exists")
	}
	err = usersToilets.Insert(cc.DB)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "")
}

func DeleteFavoriteHandler(c echo.Context) error {
	cc := c.(*CustomContext)
	placeId := c.Param("placeId")
	toilet := &local.Toilet{GoogleId: placeId}
	toiletId, err := toilet.GetToiletId(cc.DB)
	if err != nil {
		return err
	}
	token, err := util.GetToken(c)
	if err != nil {
		return err
	}
	user := &local.User{UID: token.UID}
	userId, err := user.GetUserId(cc.DB)
	if err != nil {
		return err
	}
	usersToilets := &local.UsersToilets{ToiletId: toiletId, UserId: userId}
	err = usersToilets.RemoveUsersToilets(cc.DB)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "")
}

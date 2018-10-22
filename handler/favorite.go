package handler

import (
	"github.com/labstack/echo"
	"github.com/gericass/toilet-api/util"
	"net/http"
	"github.com/gericass/toilet-api/data/local"
	"github.com/gericass/toilet-api/handler/response"
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

	return nil
}

func DeleteFavoriteHandler(c echo.Context) error {

	return nil
}

package handler

import (
	"github.com/labstack/echo"
	"github.com/gericass/toilet-api/handler/request"
	"github.com/gericass/toilet-api/data/remote"
	"github.com/gericass/toilet-api/handler/response"
	"net/http"
)

func createToilets(places *remote.Place) *response.Toilets {
	var ts []*response.Toilet
	for _, v := range places.Results {
		t := &response.Toilet{Name: v.Name, GoogleId: v.PlaceID, Geolocation: v.Vicinity,}
		ts = append(ts, t)
	}
	return &response.Toilets{Toilets: ts}
}

func SearchHandler(c echo.Context) error {
	keyword := new(request.Keyword)
	c.Bind(keyword)
	places, err := remote.SearchPlaces(keyword.Keyword)
	if err != nil {
		return err
	}
	toilets := createToilets(places)
	return c.JSON(http.StatusOK, toilets)
}

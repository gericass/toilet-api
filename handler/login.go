package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type user struct {
	GoogleId string `json:"google_id" form:"google_id" query:"google_id"`
	Name     string `json:"name" form:"name" query:"name"`
	Icon     string `json:"icon" form:"icon" query:"icon"`
}

func LoginHandler(c echo.Context) error {
	i := new(user)
	err := c.Bind(i)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "OK")
}

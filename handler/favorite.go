package handler

import "github.com/labstack/echo"

func FavoriteHandler(c echo.Context) error {
	c.Request().Header.Get("X-Toilet-Token")
	return nil
}

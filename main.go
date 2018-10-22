package main

import (
	"github.com/labstack/echo"
	"github.com/gericass/toilet-api/handler"
	"github.com/gericass/toilet-api/data/local"
	"github.com/labstack/echo/middleware"
	"google.golang.org/api/option"
	"firebase.google.com/go"
	"errors"
	"context"
	"net/http"
)

func dbMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := local.ConnectDB()
		if err != nil {
			return err
		}
		defer db.Close()
		cc := &handler.CustomContext{c, db}
		return h(cc)
	}
}

func validateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		forbidden := c.String(http.StatusForbidden, "Forbidden")
		token := c.Request().Header.Get("X-Toilet-Token")
		if token == "" {
			return errors.New("token empty")
		}
		opt := option.WithCredentialsFile("toilet-review-220105-1876f144320f.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return forbidden
		}
		client, err := app.Auth(context.Background())
		if err != nil {
			return forbidden
		}
		verifyToken, err := client.VerifyIDToken(context.Background(), token)
		if err != nil || verifyToken == nil {
			return forbidden
		}
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(dbMiddleware)
	e.Use(validateToken)
	e.Use(middleware.Logger())

	e.POST("/login", handler.LoginHandler)
	e.GET("/search", handler.SearchHandler)

	e.Start(":8000")
}

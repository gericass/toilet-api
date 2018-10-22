package util

import (
	"github.com/labstack/echo"
	"google.golang.org/api/option"
	"firebase.google.com/go"
	"golang.org/x/net/context"
	"firebase.google.com/go/auth"
)

func GetToken(c echo.Context) (*auth.Token, error) {
	token := c.Request().Header.Get("X-Toilet-Token")
	opt := option.WithCredentialsFile("toilet-review-220105-1876f144320f.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	verifyToken, err := client.VerifyIDToken(context.Background(), token)
	if err != nil || verifyToken == nil {
		return nil, err
	}
	return verifyToken, nil
}

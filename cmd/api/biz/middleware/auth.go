package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"sign-lottery/pkg/errmsg"
	"sign-lottery/pkg/jwt"
	"strings"
)

func UserJwtMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.Abort()
			c.JSON(http.StatusOK, errmsg.TokenNotExist)
			return
		}
		checkToken := strings.Split(auth, " ")
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, errmsg.TokenIsError)
			c.Abort()
			return
		}
		token, err := jwt.ParseUserToken(checkToken[1])
		if err != nil {
			c.JSON(http.StatusOK, errmsg.TokenIsError)
			c.Abort()
			return
		}
		c.Set("email", token.Email)
		c.Set("id", token.Id)
		c.Next(ctx)
	}
}

func AdminJwtMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.Abort()
			c.JSON(http.StatusOK, errmsg.TokenNotExist)
			return
		}
		checkToken := strings.Split(auth, " ")
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, errmsg.TokenIsError)
			c.Abort()
			return
		}
		token, err := jwt.ParseAdminToken(checkToken[1])
		if err != nil {
			c.JSON(http.StatusOK, errmsg.TokenIsError)
			c.Abort()
			return
		}
		c.Set("name", token.Name)
		c.Next(ctx)
	}
}

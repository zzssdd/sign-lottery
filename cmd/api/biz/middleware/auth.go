package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"sign-lottery/dao/cache"
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
		Cache := cache.NewCache()
		if !Cache.User.ExistToken(ctx, checkToken[1]) {
			c.JSON(http.StatusOK, errmsg.TokenNotExist)
			c.Abort()
			return
		}
		real_token, err := Cache.User.GetToken(ctx, checkToken[1])
		if err != nil {
			c.JSON(http.StatusOK, errmsg.TokenNotExist)
			c.Abort()
			return
		}
		token, err := jwt.ParseUserToken(real_token)
		if err == fmt.Errorf(errmsg.GetMsg(errmsg.TokenExpired)) {
			userToken, err := jwt.GenNewUserToken(real_token)
			if err != nil {
				c.JSON(http.StatusOK, errmsg.Error)
				c.Abort()
				return
			}
			err = Cache.User.StoreToken(ctx, checkToken[1], userToken)
			if err != nil {
				c.JSON(http.StatusOK, errmsg.Error)
				c.Abort()
				return
			}
		} else if err != nil {
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
		Cache := cache.NewCache()
		if !Cache.User.ExistToken(ctx, checkToken[1]) {
			c.JSON(http.StatusOK, errmsg.TokenNotExist)
			c.Abort()
			return
		}
		real_token, err := Cache.User.GetToken(ctx, checkToken[1])
		if err != nil {
			c.JSON(http.StatusOK, errmsg.TokenNotExist)
			c.Abort()
			return
		}
		token, err := jwt.ParseAdminToken(checkToken[1])
		if err == fmt.Errorf(errmsg.GetMsg(errmsg.TokenExpired)) {
			adminToken, err := jwt.GenAdminToken(real_token)
			if err != nil {
				c.JSON(http.StatusOK, errmsg.Error)
				c.Abort()
				return
			}
			err = Cache.User.StoreToken(ctx, checkToken[1], adminToken)
			if err != nil {
				c.JSON(http.StatusOK, errmsg.Error)
				c.Abort()
				return
			}
		} else if err != nil {
			c.JSON(http.StatusOK, errmsg.TokenIsError)
			c.Abort()
			return
		}
		c.Set("name", token.Name)
		c.Next(ctx)
	}
}

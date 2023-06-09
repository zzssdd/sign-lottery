package jwt

import (
	"errors"
	"fmt"
	"sign-lottery/pkg/constants"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var Secret = []byte("fe2fsaw3")
var AdminSecret = []byte("asdawer2")

type MyClaims struct {
	Id    int64  `json:"id"`
	Email string `json:"name"`
	jwt.StandardClaims
}

type AdminClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenToken(id int64, email string) (token string, err error) {
	claim := &MyClaims{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constants.TokenExpireTime * time.Hour).Unix(),
			Issuer:    "yogen",
			Subject:   "sign-lottery",
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(Secret)
	if err != nil {
		Log.Errorln("signed token err:", err)
	}
	return
}

func GenAdminToken(name string) (token string, err error) {
	claim := &AdminClaims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constants.TokenExpireTime * time.Hour).Unix(),
			Issuer:    "yogen-admin",
			Subject:   "sign-lottery",
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(AdminSecret)
	if err != nil {
		Log.Errorln("signed token err:", err)
	}
	return
}

func ParseUserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			err = fmt.Errorf(errmsg.GetMsg(errmsg.TokenExpired))
		}
		return nil, err
	}
	if token == nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claim, nil
	}
	return nil, fmt.Errorf("token不合法")
}

func GenNewUserToken(tokenString string) (string, error) {
	var claim MyClaims
	_, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return "", err
	}
	return GenToken(claim.Id, claim.Email)
}

func GenNewAdminToken(tokenString string) (string, error) {
	var claim AdminClaims
	_, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return AdminSecret, nil
	})
	if err != nil {
		return "", err
	}
	return GenAdminToken(claim.Name)
}

func ParseAdminToken(tokenString string) (*AdminClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return AdminSecret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			err = fmt.Errorf(errmsg.GetMsg(errmsg.TokenExpired))
		}
		return nil, err
	}
	if token == nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*AdminClaims); ok && token.Valid {
		return claim, nil
	}
	return nil, fmt.Errorf("token不合法")
}

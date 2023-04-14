package jwt

import (
	. "sign-lottery/pkg/log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var Secret = []byte("fe2fsaw3")
var AdminSecret = []byte("asdawer2")

type MyClaims struct {
	id    int64  `json:"id"`
	email string `json:"name"`
	jwt.StandardClaims
}

type AdminClaims struct {
	name string `json:"name"`
	jwt.StandardClaims
}

func GenToken(id int64, email string) (token string, err error) {
	claim := &MyClaims{
		id:    id,
		email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
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
		name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
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

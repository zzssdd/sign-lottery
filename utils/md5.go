package utils

import (
	"crypto/md5"
	"encoding/hex"
	"sign-lottery/pkg/constants"
)

func Crypto(password string) string {
	h := md5.New()
	h.Write([]byte(constants.Salt))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

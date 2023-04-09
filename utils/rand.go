package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func RandCode(l int) string {
	rand.Seed(time.Now().UnixNano())
	ret := ""
	for i := 0; i < l; i++ {
		ret += strconv.Itoa(rand.Intn(10))
	}
	return ret
}

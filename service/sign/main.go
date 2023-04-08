package main

import (
	"log"
	sign "sign-lottery/kitex_gen/sign/signservice"
	sign2 "sign-lottery/service/sign"
)

func main() {
	svr := sign.NewServer(new(sign2.SignServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

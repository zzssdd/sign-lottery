package main

import (
	"log"
	sign "sign-lottery/kitex_gen/sign/signservice"
)

func main() {
	svr := sign.NewServer(new(SignServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

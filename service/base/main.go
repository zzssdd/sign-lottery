package main

import (
	"log"
	user "sign-lottery/kitex_gen/user/baseservice"
	"sign-lottery/service/base"
)

func main() {
	svr := user.NewServer(new(base.BaseServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

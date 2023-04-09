package main

import (
	"log"
	user "sign-lottery/kitex_gen/user/baseservice"
	main2 "sign-lottery/service/base"
)

func main() {
	svr := user.NewServer(new(main2.BaseServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

package main

import (
	"log"
	lottery "sign-lottery/kitex_gen/lottery/lotteryservice"
	main2 "sign-lottery/service/lottery"
)

func main() {
	svr := lottery.NewServer(new(main2.LotteryServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

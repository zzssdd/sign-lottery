package lottery

import (
	"log"
	lottery "sign-lottery/kitex_gen/lottery/lotteryservice"
)

func main() {
	svr := lottery.NewServer(new(LotteryServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

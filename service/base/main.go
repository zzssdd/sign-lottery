package base

import (
	"log"
	user "sign-lottery/kitex_gen/user/baseservice"
)

func main() {
	svr := user.NewServer(new(BaseServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

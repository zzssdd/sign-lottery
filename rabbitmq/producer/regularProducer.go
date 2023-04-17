package producer

import (
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/comm"
)

type RegularProducer struct {
}

func (r *RegularProducer) ProducerRegular(msg []byte, timeStamp string) error {
	rabbit := comm.NewRabbit("regular", "sign-lottery", "regular")
	err := rabbit.RegularSetUp()
	if err != nil {
		Log.Errorln("set up regular rabbitmq err:", err)
		return err
	}
	err = rabbit.RegularSetUp()
	if err != nil {
		Log.Errorln("set up rabbitmq err:", err)
		return err
	}
	defer rabbit.Destory()
	err = rabbit.RegularSend(msg, timeStamp)
	if err != nil {
		Log.Errorln("send regular task to rabbitmq err:", err)
		return err
	}
	return nil
}

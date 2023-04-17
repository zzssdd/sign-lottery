package consumer

import (
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/comm"
)

type RegularConsumer struct {
}

func (r *RegularConsumer) ConsumerRegular(regularChan chan<- string) error {
	rabbit := comm.NewRabbit("regular", "sign-lottery", "regular")
	err := rabbit.RegularSetUp()
	if err != nil {
		Log.Errorln("set up regular rabbitmq err:", err)
		return err
	}
	defer rabbit.Destory()
	msgChan, err := rabbit.Ch.Consume(
		rabbit.QueueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		Log.Errorln("get msg from regular rabbit err:", err)
		return err
	}
	for msg := range msgChan {
		regularChan <- string(msg.Body)
	}
	return nil
}

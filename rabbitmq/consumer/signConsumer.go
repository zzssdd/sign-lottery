package consumer

import (
	"encoding/json"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/comm"
	"sign-lottery/rabbitmq/model"
)

type SignConsumer struct{}

func (s *SignConsumer) ConsumerSign(signChan chan<- model.Sign) error {
	rabbit := comm.NewRabbit("sign", "sign-lottery", "sign")
	err := rabbit.SetUp()
	if err != nil {
		Log.Errorln("rabbitmq set up err:", err)
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
		Log.Errorln("get msg from rabbit err:", err)
		return err
	}
	var signInfo model.Sign
	for msg := range msgChan {
		err = json.Unmarshal(msg.Body, &signInfo)
		if err != nil {
			Log.Errorln("unmarshal sign json err:", err)
			return err
		}
		signChan <- signInfo
		msg.Ack(true)
		rabbit.Ch.Ack(msg.DeliveryTag, false)
	}
	return nil
}

package consumer

import (
	"encoding/json"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/comm"
	"sign-lottery/rabbitmq/model"
)

type ChooseConsumer struct{}

func (c *ChooseConsumer) ConsumerChoose(chooseChan chan<- model.Choose) error {
	rabbit := comm.NewRabbit("choose", "sign-lottery", "choose")
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
	var chooseInfo model.Choose
	for msg := range msgChan {
		err = json.Unmarshal(msg.Body, &chooseInfo)
		if err != nil {
			Log.Errorln("unmarshal choose json err:", err)
			return err
		}
		chooseChan <- chooseInfo
		msg.Ack(true)
		rabbit.Ch.Ack(msg.DeliveryTag, false)
	}
	return nil
}

package consumer

import (
	"encoding/json"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/comm"
	"sign-lottery/rabbitmq/model"
)

type OrderConsumer struct {
}

func (o *OrderConsumer) ProducerOrder(orderChan chan<- model.Order) error {
	rabbit := comm.NewRabbit("order", "sign-lottery", "order")
	err := rabbit.SetUp()
	if err != nil {
		Log.Errorln("set up rabbitmq err:", err)
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
	var orderInfo model.Order
	for msg := range msgChan {
		err = json.Unmarshal(msg.Body, &orderInfo)
		if err != nil {
			Log.Errorln("unmarshal order json err:", err)
			continue
		}
		orderChan <- orderInfo
		msg.Ack(true)
		rabbit.Ch.Ack(msg.DeliveryTag, false)
	}
	return nil
}

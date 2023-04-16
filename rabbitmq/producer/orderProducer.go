package producer

import (
	"encoding/json"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/comm"
	"sign-lottery/rabbitmq/model"
)

type OrderProducer struct {
}

func (o *OrderProducer) ProducerOrder(order *model.Order) error {
	rabbit := comm.NewRabbit("order", "sign-lottery", "order")
	err := rabbit.SetUp()
	if err != nil {
		Log.Errorln("set up rabbitmq err:", err)
		return err
	}
	defer rabbit.Destory()
	byte_data, err := json.Marshal(*order)
	if err != nil {
		Log.Errorln("transfer order to json err:", err)
		return err
	}
	err = rabbit.Send(byte_data)
	if err != nil {
		Log.Errorln("send order to rabbitmq err:", err)
		return err
	}
	return nil
}

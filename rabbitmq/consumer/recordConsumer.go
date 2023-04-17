package consumer

import (
	"encoding/json"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/comm"
	"sign-lottery/rabbitmq/model"
)

type RecordConsumer struct {
}

func (r *RecordConsumer) ConsumerRecord(recordChan chan<- model.Record) error {
	rabbit := comm.NewRabbit("record", "sign-lottery", "record")
	err := rabbit.SetUp()
	if err != nil {
		Log.Errorln("set up record rabbitmq err:", err)
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
	var record model.Record
	for msg := range msgChan {
		err = json.Unmarshal(msg.Body, &record)
		if err != nil {
			Log.Errorln("unmarshal record json err:", err)
			continue
		}
		recordChan <- record
		msg.Ack(true)
		rabbit.Ch.Ack(msg.DeliveryTag, false)
	}
	return nil
}

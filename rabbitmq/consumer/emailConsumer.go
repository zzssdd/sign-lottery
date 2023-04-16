package consumer

import (
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/comm"
)

type EmailConsumer struct{}

func (e *EmailConsumer) ConsumerEmail(emailChan chan<- string) error {
	rabbit := comm.NewRabbit("email", "sign-lottery", "email")
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
	for msg := range msgChan {
		msg.Ack(true)
		rabbit.Ch.Ack(msg.DeliveryTag, false)
		emailChan <- string(msg.Body)
	}
	return nil
}

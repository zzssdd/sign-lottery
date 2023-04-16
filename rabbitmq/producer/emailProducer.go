package producer

import (
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/comm"
)

type EmailProducer struct{}

func (e *EmailProducer) ProducerEmail(email string) error {
	rabbit := comm.NewRabbit("email", "sign-lottery", "email")
	err := rabbit.SetUp()
	if err != nil {
		Log.Errorln("rabbit set up err:", err)
		return err
	}
	defer rabbit.Destory()
	err = rabbit.Send([]byte(email))
	if err != nil {
		Log.Errorln("send email msg err:", err)
		return err
	}
	return nil
}

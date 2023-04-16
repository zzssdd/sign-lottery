package producer

import (
	"encoding/json"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/comm"
	"sign-lottery/rabbitmq/model"
)

type SignProducer struct{}

func (s *SignProducer) ProducerSign(sign *model.Sign) error {
	rabbit := comm.NewRabbit("sign", "sign-lottery", "choose")
	err := rabbit.SetUp()
	if err != nil {
		Log.Errorln("rabbit set up err:", err)
		return err
	}
	defer rabbit.Destory()
	byte_data, err := json.Marshal(*sign)
	if err != nil {
		Log.Errorln("transfer sign into json err:", err)
		return err
	}
	err = rabbit.Send(byte_data)
	if err != nil {
		Log.Errorln("send sign msg err:", err)
		return err
	}
	return nil
}

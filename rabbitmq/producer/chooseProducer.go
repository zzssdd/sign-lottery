package producer

import (
	"encoding/json"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/comm"
	"sign-lottery/rabbitmq/model"
)

type ChooseProcuer struct{}

func (c *ChooseProcuer) ProducerChoose(choose *model.Choose) error {
	rabbit := comm.NewRabbit("choose", "sign-lottery", "choose")
	err := rabbit.SetUp()
	if err != nil {
		Log.Errorln("rabbit set up err:", err)
		return err
	}
	defer rabbit.Destory()
	byte_data, err := json.Marshal(*choose)
	if err != nil {
		Log.Errorln("transfer choose into json err:", err)
		return err
	}
	err = rabbit.Send(byte_data)
	if err != nil {
		Log.Errorln("send choose msg err:", err)
		return err
	}
	return nil
}

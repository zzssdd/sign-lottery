package producer

import (
	"encoding/json"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/comm"
	"sign-lottery/rabbitmq/model"
)

type RecordProducer struct {
}

func (r *RecordProducer) ProducerRecord(record *model.Record) error {
	rabbit := comm.NewRabbit("record", "sign-lottery", "record")
	err := rabbit.SetUp()
	if err != nil {
		Log.Errorln("set up rabbit err:", err)
		return err
	}
	byte_data, err := json.Marshal(record)
	if err != nil {
		Log.Errorln("transfer record to json err:", err)
		return err
	}
	err = rabbit.Send(byte_data)
	if err != nil {
		Log.Errorln("send record to rabbitmq err:", err)
		return err
	}
	return nil
}

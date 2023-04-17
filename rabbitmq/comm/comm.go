package comm

import (
	"github.com/streadway/amqp"
	"sign-lottery/pkg/constants"
	. "sign-lottery/pkg/log"
)

type RabbitMQ struct {
	Conn         *amqp.Connection
	Ch           *amqp.Channel
	QueueName    string
	ExchangeName string
	Key          string
}

func NewRabbit(queueName string, exchangeName string, key string) *RabbitMQ {
	conn, err := amqp.Dial(constants.RabbitMqDSN)
	if err != nil {
		Log.Errorln("connect rabbitmq err:", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		Log.Errorln("create channel err:", err)
	}
	return &RabbitMQ{
		Conn:         conn,
		Ch:           ch,
		QueueName:    queueName,
		ExchangeName: exchangeName,
		Key:          key,
	}
}

func (r *RabbitMQ) SetUp() error {
	q, err := r.Ch.QueueDeclare(r.QueueName, false, false, false, false, nil)
	if err != nil {
		Log.Errorln("create queue err:", err)
		return err
	}
	err = r.Ch.ExchangeDeclare(r.ExchangeName, amqp.ExchangeDirect, false, true, false, true, nil)
	if err != nil {
		Log.Errorln("create exchange err:", err)
		return err
	}
	err = r.Ch.QueueBind(q.Name, r.Key, r.ExchangeName, true, nil)
	if err != nil {
		Log.Errorln("bind queue err:", err)
		return err
	}
	return nil
}

func (r *RabbitMQ) RegularSetUp() error {
	q, err := r.Ch.QueueDeclare(r.QueueName, true, false, false, false, nil)
	if err != nil {
		Log.Errorln("create regular queue err:", err)
		return err
	}
	err = r.Ch.ExchangeDeclare(r.ExchangeName, "x-delayed-message", true, false, false, false, amqp.Table{
		"x-delayed-type": amqp.ExchangeDirect,
	})
	if err != nil {
		Log.Errorln("create regular exchange err:", err)
		return err
	}
	err = r.Ch.QueueBind(q.Name, r.Key, r.ExchangeName, false, nil)
	if err != nil {
		Log.Errorln("bind regular queue err:", err)
		return err
	}
	return nil
}

func (r *RabbitMQ) RegularSend(data []byte, timeStamp string) error {
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
		Headers: map[string]interface{}{
			"x-delay": timeStamp,
		},
	}
	return r.Ch.Publish(r.ExchangeName, r.Key, false, false, msg)
}

func (r *RabbitMQ) Send(data []byte) error {
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
	}
	return r.Ch.Publish(r.ExchangeName, r.Key, false, false, msg)
}

func (r *RabbitMQ) Destory() {
	r.Ch.Close()
	r.Conn.Close()
}

package consumer

type Consumer struct {
	Email  EmailConsumer
	Choose ChooseConsumer
	Sign   SignConsumer
	Order  OrderConsumer
}

func NewConsumer() *Consumer {
	return &Consumer{
		Email:  EmailConsumer{},
		Choose: ChooseConsumer{},
		Sign:   SignConsumer{},
		Order:  OrderConsumer{},
	}
}

package consumer

type Consumer struct {
	Email   EmailConsumer
	Choose  ChooseConsumer
	Sign    SignConsumer
	Order   OrderConsumer
	Regular RegularConsumer
	Record  RecordConsumer
}

func NewConsumer() *Consumer {
	return &Consumer{
		Email:   EmailConsumer{},
		Choose:  ChooseConsumer{},
		Sign:    SignConsumer{},
		Order:   OrderConsumer{},
		Regular: RegularConsumer{},
		Record:  RecordConsumer{},
	}
}

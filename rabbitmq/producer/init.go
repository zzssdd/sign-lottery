package producer

type Producer struct {
	Email  EmailProducer
	Choose ChooseProcuer
	Sign   SignProducer
	Order  OrderProducer
}

func NewProcuer() *Producer {
	return &Producer{
		Email:  EmailProducer{},
		Choose: ChooseProcuer{},
		Sign:   SignProducer{},
		Order:  OrderProducer{},
	}
}

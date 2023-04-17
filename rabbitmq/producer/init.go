package producer

type Producer struct {
	Email   EmailProducer
	Choose  ChooseProcuer
	Sign    SignProducer
	Order   OrderProducer
	Regular RegularProducer
	Record  RecordProducer
}

func NewProcuer() *Producer {
	return &Producer{
		Email:   EmailProducer{},
		Choose:  ChooseProcuer{},
		Sign:    SignProducer{},
		Order:   OrderProducer{},
		Regular: RegularProducer{},
		Record:  RecordProducer{},
	}
}

package payload

func NewBrokerService() Service {
	return &broker{}
}

func (b *broker) Broker() (*JsonResponse, error) {
	res := &JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	return res, nil
}

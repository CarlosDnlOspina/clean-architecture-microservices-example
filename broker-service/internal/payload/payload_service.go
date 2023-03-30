package payload

import "context"

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

func (b *broker) Authenticate(ctx context.Context, authPayload *AuthPayload) (*JsonResponse, error) {
	res := &JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	return res, nil
}

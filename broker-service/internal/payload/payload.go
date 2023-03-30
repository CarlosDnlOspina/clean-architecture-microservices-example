package payload

import "context"

type broker struct{}

type AuthPayload struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type Repository interface {
	Broker() (*JsonResponse, error)
}

type Service interface {
	Broker() (*JsonResponse, error)
	Authenticate(ctx context.Context, authPayload *AuthPayload) (*JsonResponse, error)
}

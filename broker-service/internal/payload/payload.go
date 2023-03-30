package payload

type broker struct{}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type RequestPayload struct {
	Auth AuthPayload `json:"auth,omitempty"`
}

type Repository interface {
	Broker() (*JsonResponse, error)
}

type Service interface {
	Broker() (*JsonResponse, error)
}

package main

import (
	"broker/internal/payload"
	"broker/router"
)

func main() {

	payloadSvc := payload.NewBrokerService()
	payloadHandler := payload.NewHandler(payloadSvc)
	router.InitRouter(payloadHandler)

	router.Start("0.0.0.0:8080")
}

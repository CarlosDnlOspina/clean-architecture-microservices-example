package main

import (
	"chat/config"
	"chat/db/cockroach"
	"chat/internal/user"
	"chat/router"
	"fmt"
	"log"
)

func main() {

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	log.Println(cfg.PG.URL)
	pg, err := cockroach.New(cfg.PG.URL, cockroach.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatalf("could not initialize database connection:: %s", err)
	}

	fmt.Printf("connectes")
	userRep := user.NewRepository(pg)
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")

}

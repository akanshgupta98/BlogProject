package main

import (
	"log"

	"github.com/akanshgupta98/BlogProject/gateway/internal/config"
	"github.com/akanshgupta98/BlogProject/gateway/internal/server"
)

func main() {

	cfg := config.New()
	srv := server.NewServer(cfg)
	srv.RegisterRoutes()
	err := srv.Serve()

	if err != nil {
		log.Fatalf("unable to start server")
	}

}

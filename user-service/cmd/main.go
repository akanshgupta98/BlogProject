package main

import (
	"log"

	"github.com/akanshgupta98/BlogProject/user-service/internal/config"
	"github.com/akanshgupta98/BlogProject/user-service/internal/server"
	"github.com/akanshgupta98/BlogProject/user-service/internal/service"
)

func main() {

	cfg := config.NewConfig()

	s := server.NewServer(cfg)
	if err := service.InitService(); err != nil {
		log.Fatal("unable to init service layer", err)
	}
	log.Println("started grpc server")
	// go func() {
	if err := s.Serve(); err != nil {
		log.Fatal("unable to start server", err)
	}
	// }()

}

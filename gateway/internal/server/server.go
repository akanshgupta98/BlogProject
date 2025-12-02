package server

import (
	"fmt"
	"log"

	"github.com/akanshgupta98/BlogProject/gateway/internal/config"
	v1 "github.com/akanshgupta98/BlogProject/gateway/internal/server/handlers/v1"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Mux  *gin.Engine
	Addr string
}

func NewServer(cfg config.Config) *Server {

	mux := gin.Default()
	addr := cfg.Srv.IP + ":" + cfg.Srv.Port
	return &Server{
		Mux:  mux,
		Addr: addr,
	}

}

func (s *Server) Serve() error {
	log.Printf("starting server on addr: %s ...\n", s.Addr)
	err := s.Mux.Run(s.Addr)
	if err != nil {
		return fmt.Errorf("unable to start the server: %s", err.Error())
	}
	return nil
}

func (s *Server) RegisterRoutes() {
	gv1 := s.Mux.Group("/api/v1/")
	v1.RegisterRoutes(gv1)
	s.Mux.GET("/", HelloWorld)
}

package server

import (
	"log"
	"net"

	pb "github.com/akanshgupta98/BlogProject/proto/protogen/user"
	"github.com/akanshgupta98/BlogProject/user-service/internal/config"
	"google.golang.org/grpc"
)

type GRPC_Server struct {
	Addr string
}
type server struct {
	pb.UserServiceServer
}

func NewServer(cfg config.Config) *GRPC_Server {

	addr := cfg.Grpc.IP + cfg.Grpc.Port
	return &GRPC_Server{
		Addr: addr,
	}

}

func (g *GRPC_Server) Serve() error {
	lis, err := net.Listen("tcp", g.Addr)
	if err != nil {
		log.Printf("error starting grpc server: %s", err.Error())
		return err
	}

	s := grpc.NewServer()

	// Register service.
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Printf("error starting grpc server: %s", err.Error())
		return err
	}
	return nil

}

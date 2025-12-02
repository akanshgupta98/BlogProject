package server

import (
	"context"

	pb "github.com/akanshgupta98/BlogProject/proto/protogen/user"
	"github.com/akanshgupta98/BlogProject/user-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) CreateUser(ctx context.Context, req *pb.UserData) (*pb.UserID, error) {

	// Create User.
	user := service.UserData{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Username: req.Username,
	}
	// Validations

	switch {
	case user.Email == "":
		return nil, status.Errorf(codes.InvalidArgument, "empty email not allowed")
	case user.Name == "":
		return nil, status.Errorf(codes.InvalidArgument, "empty username not allowed")
	case user.Username == "":
		return nil, status.Errorf(codes.InvalidArgument, "empty name not allowed")

	}

	// Send to service layer
	resp, err := service.CreateUser(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create user:%s", err.Error())
	}

	result := pb.UserID{
		Id: uint64(resp),
	}

	return &result, nil

}

func (s *server) FetchUserByID(ctx context.Context, id *pb.UserID) (*pb.UserData, error) {
	// Send to service layer
	// uid := id.Id
	userData := service.UserData{
		ID: int(id.Id),
	}
	resp, err := service.FetchUserByID(userData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to fetch user:%s", err.Error())
	}

	result := &pb.UserData{
		Email:    resp.Email,
		Id:       uint64(resp.ID),
		Name:     resp.Name,
		Phone:    resp.Phone,
		Username: resp.Username,
	}

	return result, nil

}

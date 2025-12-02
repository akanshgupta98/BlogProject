package v1

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "github.com/akanshgupta98/BlogProject/proto/protogen/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/users", FetchUsers)
	rg.POST("/register", RegisterUser)
	rg.GET("/users/:id", FetchUserByID)
}

func DoFetchUser(c pb.UserServiceClient, data *pb.UserID) (*pb.UserData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := c.FetchUserByID(ctx, data)
	return resp, err

}
func sendGrpcFetchByIdReq(uid int) (*pb.UserData, error) {
	client, err := grpc.NewClient(":50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("unable to create grpc client", err.Error())
		return nil, err
	}
	defer client.Close()

	c := pb.NewUserServiceClient(client)
	data := &pb.UserID{
		Id: uint64(uid),
	}
	resp, err := DoFetchUser(c, data)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Println("GRPC ERROR", e.Code(), e.Message())
		} else {
			log.Println("NON grpc error.")

		}
		return nil, err

	}
	return resp, nil

}
func FetchUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		log.Println("invalid id")
		ctx.JSON(http.StatusBadGateway, nil)
		return
	}
	resp, err := sendGrpcFetchByIdReq(uid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	data := UserReq{
		ID:       int(resp.Id),
		Name:     resp.Name,
		Email:    resp.Email,
		Username: resp.Username,
		Phone:    resp.Phone,
	}
	ctx.JSON(http.StatusOK, data)

}

func FetchUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"repsone": "Dummy users",
	})

}

type UserReq struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	Name     string `json:"name"`
}

func DoCreateUser(c pb.UserServiceClient, data *pb.UserData) (*pb.UserID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := c.CreateUser(ctx, data)
	return resp, err

}
func sendGrpcReq(data *pb.UserData) (*pb.UserID, error) {
	client, err := grpc.NewClient(":50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("unable to create grpc client", err.Error())
		return nil, err
	}
	defer client.Close()

	c := pb.NewUserServiceClient(client)

	resp, err := DoCreateUser(c, data)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Println("GRPC ERROR", e.Code(), e.Message())
		} else {
			log.Println("NON grpc error.")

		}
		return nil, err

	}
	return resp, nil

}

func RegisterUser(ctx *gin.Context) {
	// Extract request payload.
	reqPayload := UserReq{}
	if err := ctx.BindJSON(&reqPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Convert into required format for user-service grpc
	user := &pb.UserData{
		Name:     reqPayload.Name,
		Email:    reqPayload.Email,
		Phone:    reqPayload.Phone,
		Username: reqPayload.Username,
	}
	// Send register user request to user-service
	resp, err := sendGrpcReq(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"id": resp,
	})

}

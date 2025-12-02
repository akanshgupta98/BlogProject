package service

import (
	"log"

	"github.com/akanshgupta98/BlogProject/user-service/internal/repository"
)

type UserData struct {
	ID       int
	Username string
	Email    string
	Phone    string
	Name     string
}

var s Service

type Service struct {
	repo *repository.Repo
}

func InitService() error {
	r, err := repository.RepoInit()
	if err != nil {
		log.Println("Error initializing repo")
		return err
	}

	s = Service{
		repo: r,
	}
	return nil
}

func CreateUser(user UserData) (int, error) {
	userData := repository.UserTable{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Username: user.Username,
	}

	resp, err := s.repo.CreateUser(userData)
	if err != nil {
		return 0, err
	}
	return resp, nil
}
func FetchUserByID(user UserData) (UserData, error) {
	userData := repository.UserTable{
		ID: user.ID,
	}
	result := UserData{}

	resp, err := s.repo.FetchUserByID(userData)
	if err != nil {
		return result, err
	}
	result.Email = resp.Email
	result.Name = resp.Name
	result.Phone = resp.Phone
	result.ID = resp.ID
	result.Username = resp.Username

	return result, nil
}

package usecase

import (
	"context"

	"github.com/bouroo/go-project-structure/api/auth/repository"
	"github.com/bouroo/go-project-structure/pkg/model"
	pb "github.com/bouroo/go-project-structure/pkg/proto/user"
	"golang.org/x/crypto/bcrypt"
)

func Signin(email string, password string) (token string, err error) {
	grpcClient, err := repository.NewUserAccountServiceClient()
	if err != nil {
		return
	}

	resp, err := grpcClient.ReadUserAccount(context.Background(), &pb.UserAccount{Email: email, Password: password})
	if err != nil {
		return
	}

	token, err = repository.GenAccessToken(resp.Id, resp.Email)

	return
}

func Singup(email string, password string) (userAccount model.UserAccount, err error) {
	grpcClient, err := repository.NewUserAccountServiceClient()
	if err != nil {
		return
	}
	resp, err := grpcClient.CreateUserAccount(context.Background(), &pb.UserAccount{Email: email, Password: password})
	if err != nil {
		return
	}
	userAccount = model.UserAccount{
		ID:       resp.Id,
		Email:    resp.Email,
		Password: resp.Password,
	}
	return
}

func ChangePassword(userID string, password string) (err error) {
	grpcClient, err := repository.NewUserAccountServiceClient()
	if err != nil {
		return
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	_, err = grpcClient.UpdateUserAccount(context.Background(), &pb.UserAccount{Id: userID, Password: string(passwordBytes)})
	return
}

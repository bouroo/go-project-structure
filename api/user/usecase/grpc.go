package usecase

import (
	"context"

	"github.com/bouroo/go-project-structure/api/user/repository"
	"github.com/bouroo/go-project-structure/pkg/entity"
	pb "github.com/bouroo/go-project-structure/pkg/proto/user"
)

type UserAccountServiceServer struct {
	pb.UnimplementedUserAccountServiceServer
}

func (s *UserAccountServiceServer) CreateUserAccount(ctx context.Context, request *pb.UserAccount) (response *pb.UserAccount, err error) {
	user := &entity.UserAccount{
		Email:    request.Email,
		Password: request.Password,
	}
	err = repository.CreateUserAccount(user)
	response = &pb.UserAccount{
		Id:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
	return
}

func (s *UserAccountServiceServer) ReadUserAccount(ctx context.Context, request *pb.UserAccount) (response *pb.UserAccount, err error) {
	user, err := repository.ReadUserAccount(request.Id)
	response = &pb.UserAccount{
		Id:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}
	return
}

func (s *UserAccountServiceServer) UpdateUserAccount(ctx context.Context, request *pb.UserAccount) (response *pb.UserAccount, err error) {
	user := entity.UserAccount{
		Email:    request.Email,
		Password: request.Password,
	}
	err = repository.UpdateUserAccount(request.Id, user)
	response = &pb.UserAccount{
		Id:    user.ID,
		Email: user.Email,
	}
	return
}

func (s *UserAccountServiceServer) DeleteUserAccount(ctx context.Context, request *pb.UserAccount) (response *pb.UserAccount, err error) {
	err = repository.DeleteUserAccount(request.Id)
	return
}
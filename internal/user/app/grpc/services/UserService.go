package grpc_services

import (
	"context"
	"delivery/internal/user/domain/entities"
	user_grpc "delivery/internal/user/infra/grpc/user"
	"delivery/internal/user/infra/models"
	"delivery/internal/user/infra/repositories"
	pkgUtils "delivery/pkg/utils"
	"fmt"

	"google.golang.org/protobuf/types/known/structpb"
)

type UserService struct {
	user_grpc.UnimplementedUserServiceServer
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

var _ entities.UserEntity = (*models.User)(nil)

func (us *UserService) Create(ctx context.Context, in *user_grpc.CreateUserRequest) (*user_grpc.Response, error) {
	res, err := us.userRepo.Create(ctx, &models.User{
		Email:    in.Email,
		Password: pkgUtils.Sha512(in.Password),
		Role:     in.Role,
	})

	if err != nil {
		return &user_grpc.Response{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	stuctRes, err := structpb.NewStruct(map[string]any{
		"id":    res.GetID(),
		"email": res.GetEmail(),
		"role":  res.GetRole(),
	})

	if err != nil {
		return &user_grpc.Response{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Result:  stuctRes,
	}, nil
}
func (us *UserService) Delete(ctx context.Context, in *user_grpc.DeleteUserRequest) (*user_grpc.Response, error) {
	if ok, err := us.userRepo.Delete(ctx, in.ID); !ok || err != nil {
		return &user_grpc.Response{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	stuctRes, err := structpb.NewStruct(map[string]any{
		"id": in.ID,
	})

	if err != nil {
		return &user_grpc.Response{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Result:  stuctRes,
	}, nil

}
func (us *UserService) GetMany(ctx context.Context, in *user_grpc.DynamicGETQueryRequest) (*user_grpc.Response, error) {
	fmt.Println(in.Params)

	return &user_grpc.Response{
		Code:    200,
		Message: "success",
		Result:  nil,
	}, nil

}
func (us *UserService) GetOne(ctx context.Context, in *user_grpc.EmptyUserRequest) (*user_grpc.Response, error) {
	return nil, nil
}
func (us *UserService) Search(ctx context.Context, in *user_grpc.EmptyUserRequest) (*user_grpc.Response, error) {
	return nil, nil
}
func (us *UserService) Update(ctx context.Context, in *user_grpc.UpdateUserRequest) (*user_grpc.Response, error) {
	return nil, nil
}

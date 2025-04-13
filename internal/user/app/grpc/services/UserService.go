package grpc_services

import (
	"context"
	user_grpc "delivery/internal/user/infra/grpc/user"
	"delivery/internal/user/infra/models"
	"delivery/internal/user/infra/repositories"
	pkgUtils "delivery/pkg/utils"

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
		Result:  []*structpb.Struct{stuctRes},
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

	stuctRes, err := structpb.NewValue(struct {
		id string
	}{
		id: in.ID,
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
func (us *UserService) GetMany(ctx context.Context, in *user_grpc.EmptyUserRequest) (*user_grpc.Response, error) {

	if in.Offset <= 0 {
		in.Offset = 10
	}
	if in.Page <= 0 {
		in.Page = 1
	}

	res, err := us.userRepo.GetMany(ctx, in.Page, in.Offset)
	if err != nil {
		return &user_grpc.Response{
			Code:    200,
			Message: "success",
			Result:  []*structpb.Struct{},
		}, nil
	}

	resultMap := make([]any, in.Offset)
	i := 0
	for _, r := range res {
		resultMap[i] = struct {
			id, email, role string
		}{
			id:    r.GetID(),
			email: r.GetEmail(),
			role:  r.GetRole(),
		}
		i++
	}

	structpb.NewList()

	data, err := structpb.NewList(resultMap[:i])

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
		Result:  resultMap,
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

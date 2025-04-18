package grpc_services

import (
	"context"
	profile_grpc "delivery/internal/user/infra/grpc/profile"
	"delivery/internal/user/infra/models"
	"delivery/internal/user/infra/repositories"
	"strconv"

	"google.golang.org/protobuf/types/known/structpb"
)

type ProfileService struct {
	repository *repositories.ProfileRepository
	profile_grpc.UnimplementedProfileServiceServer
}

func NewProfileService(repository *repositories.ProfileRepository) *ProfileService {
	return &ProfileService{
		repository: repository,
	}
}

func (_this *ProfileService) Create(ctx context.Context, in *profile_grpc.CreateProfileRequest) (*profile_grpc.ProfileResponse, error) {
	res, err := _this.repository.Create(ctx, &models.Profile{
		UserID:   in.UserID,
		FullName: in.FullName,
		Avatar:   in.Avatar,
		Bio:      in.Bio,
	})
	if err != nil {
		return &profile_grpc.ProfileResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	stuctRes, err := structpb.NewValue(map[string]any{
		"id":        res.GetID(),
		"Full_name": res.GetFullName(),
		"bio":       res.GetBio(),
		"avatr":     res.GetAvatar(),
	})

	if err != nil {
		return &profile_grpc.ProfileResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	return &profile_grpc.ProfileResponse{
		Code:    200,
		Message: "success",
		Result:  []*structpb.Value{stuctRes},
	}, nil

}

func (_this *ProfileService) Delete(ctx context.Context, in *profile_grpc.DeleteProfileRequest) (*profile_grpc.ProfileResponse, error) {
	isDone, err := _this.repository.Delete(ctx, in.GetID())
	if err != nil || !isDone {
		return &profile_grpc.ProfileResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	return &profile_grpc.ProfileResponse{
		Code:    200,
		Message: "success",
		Result:  nil,
	}, nil
}

func (_this *ProfileService) Update(ctx context.Context, in *profile_grpc.UpdateProfileRequest) (*profile_grpc.ProfileResponse, error) {
	res, err := _this.repository.Update(ctx, &models.Profile{
		ID:       in.ID,
		FullName: in.Fields["full_name"],
		Avatar:   in.Fields["avatar"],
		Bio:      in.Fields["bio"],
	})
	if err != nil {
		return &profile_grpc.ProfileResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	stuctRes, err := structpb.NewValue(map[string]any{
		"id":        res.GetID(),
		"Full_name": res.GetFullName(),
		"bio":       res.GetBio(),
		"avatr":     res.GetAvatar(),
	})

	if err != nil {
		return &profile_grpc.ProfileResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	return &profile_grpc.ProfileResponse{
		Code:    200,
		Message: "success",
		Result:  []*structpb.Value{stuctRes},
	}, nil
}

func (_this *ProfileService) GetOne(ctx context.Context, in *profile_grpc.EmptyProfileResponse) (*profile_grpc.ProfileResponse, error) {
	res, err := _this.repository.GetOne(ctx, "in.Page")
	if err != nil {
		return &profile_grpc.ProfileResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	stuctRes, err := structpb.NewValue(map[string]any{
		"id":        res.GetID(),
		"Full_name": res.GetFullName(),
		"bio":       res.GetBio(),
		"avatr":     res.GetAvatar(),
	})

	if err != nil {
		return &profile_grpc.ProfileResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	return &profile_grpc.ProfileResponse{
		Code:    200,
		Message: "success",
		Result:  []*structpb.Value{stuctRes},
	}, nil
}

func (_this *ProfileService) GetMany(ctx context.Context, in *profile_grpc.GetRequest) (*profile_grpc.ProfileResponse, error) {

	params := in.GetQueryParams().GetFields()

	page := params["page"].GetStringValue()
	offset := params["offset"].GetStringValue()
	id := params["id"].GetStringValue()

	// Convert string values to other types if needed
	idInt, _ := strconv.Atoi(id)

	res, err := _this.repository.GetMany(ctx, int(in.Offset), int(in.Page))
	if err != nil {
		return &profile_grpc.ProfileResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	items := make([]*structpb.Value, len(res))
	i := 0
	for _, item := range res {
		itm, err := structpb.NewValue(map[string]any{
			"id":        item.GetID(),
			"Full_name": item.GetFullName(),
			"bio":       item.GetBio(),
			"avatr":     item.GetAvatar(),
		})
		if err != nil {
			continue
		}

		items[i] = itm
		i++
	}

	return &profile_grpc.ProfileResponse{
		Code:    200,
		Message: "success",
		Result:  items[:i],
	}, nil
}

func (_this *ProfileService) Search(ctx context.Context, in *profile_grpc.GetRequest) (*profile_grpc.ProfileResponse, error) {
	res, err := _this.repository.Search(ctx, "in.GetPage()", int(in.Offset), int(in.Page))
	if err != nil {
		return &profile_grpc.ProfileResponse{
			Code:    400,
			Message: err.Error(),
			Result:  nil,
		}, err
	}

	items := make([]*structpb.Value, len(res))
	i := 0
	for _, item := range res {
		itm, err := structpb.NewValue(map[string]any{
			"id":        item.GetID(),
			"Full_name": item.GetFullName(),
			"bio":       item.GetBio(),
			"avatr":     item.GetAvatar(),
		})
		if err != nil {
			continue
		}

		items[i] = itm
		i++
	}

	return &profile_grpc.ProfileResponse{
		Code:    200,
		Message: "success",
		Result:  items[:i],
	}, nil
}

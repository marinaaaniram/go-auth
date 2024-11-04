package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/marinaaaniram/go-auth/internal/model"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

func FromUserToDesc(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id: user.ID,
		UserInfo: &desc.UserInfo{
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: updatedAt,
		},
	}
}

func FromDescCreateToUser(req *desc.CreateRequest) *model.User {
	return &model.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     req.GetRole(),
	}
}

func FromDescGetToUser(req *desc.GetRequest) *model.User {
	return &model.User{
		ID: req.GetId(),
	}
}

func FromDescUpdateToUser(req *desc.UpdateRequest) *model.User {
	var name string
	if req.GetName() != nil {
		name = req.GetName().GetValue()
	}

	return &model.User{
		ID:   req.GetId(),
		Name: name,
		Role: req.GetRole(),
	}
}

func FromDescDeleteToUser(req *desc.DeleteRequest) *model.User {
	return &model.User{
		ID: req.GetId(),
	}
}

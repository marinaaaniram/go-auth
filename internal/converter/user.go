package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/marinaaaniram/go-auth/internal/constant"
	"github.com/marinaaaniram/go-auth/internal/model"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Convert model.UserRole format to desc.RoleEnum
func roleModelToDesc(role constant.UserRole) desc.RoleEnum {
	switch role {
	case constant.AdminUserRole:
		return desc.RoleEnum_ADMIN
	case constant.UserUserRole:
		return desc.RoleEnum_USER
	default:
		return desc.RoleEnum_UNKNOWN
	}
}

// Convert desc.RoleEnum format to model.UserRole
func roleDescToModel(role desc.RoleEnum) constant.UserRole {
	switch role {
	case desc.RoleEnum_ADMIN:
		return constant.AdminUserRole
	case desc.RoleEnum_USER:
		return constant.UserUserRole
	default:
		return constant.UnknowUserRole
	}
}

// Convert User internal model to desc model
func FromUserToDesc(user *model.User) *desc.User {
	if user == nil {
		return nil
	}

	descUser := &desc.User{
		Id: user.ID,
		UserInfo: &desc.UserInfo{
			Name:      user.Name,
			Email:     user.Email,
			Role:      roleModelToDesc(user.Role),
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}

	if user.UpdatedAt != nil {
		descUser.UserInfo.UpdatedAt = timestamppb.New(*user.UpdatedAt)
	}

	return descUser
}

// Convert User internal model to desc model
func FromUserIdToDescCreate(id int64) *desc.CreateResponse {
	return &desc.CreateResponse{
		Id: id,
	}
}

// Convert desc CreateRequest fields to internal User model
func FromDescCreateToUser(req *desc.CreateRequest) *model.User {
	if req == nil {
		return nil
	}

	return &model.User{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     roleDescToModel(req.GetRole()),
	}
}

// Convert desc UpdateRequest fields to internal User model
func FromDescUpdateToUser(req *desc.UpdateRequest) *model.User {
	if req == nil {
		return nil
	}

	var name string
	if req.GetName() != nil {
		name = req.GetName().GetValue()
	}

	return &model.User{
		ID:   req.GetId(),
		Name: name,
		Role: roleDescToModel(req.GetRole()),
	}
}

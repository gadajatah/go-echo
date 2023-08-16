package services

import (
	"context"
	"crud-test/models/web"
)

type UserService interface {
	Create(ctx context.Context, req web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, req web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}

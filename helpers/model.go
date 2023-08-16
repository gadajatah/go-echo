package helpers

import (
	"crud-test/helpers"
	"crud-test/models/domain"
	"crud-test/models/web"
)

func ToUserResponse(user domain.Users) web.UserResponse {
	return web.UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}
}

func ToUserResponses(users []domain.Users) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, helpers.ToUserResponse(user))
	}
	return userResponses
}

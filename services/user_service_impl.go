package services

import (
	"context"
	"crud-test/helpers"
	"crud-test/models/domain"
	"crud-test/models/web"
	"crud-test/repositories"
	"database/sql"
)

type UserServiceImpl struct {
	UsersRepository repositories.UsersRepository
	DB              *sql.DB
}

func (service *UserServiceImpl) Create(ctx context.Context, req web.UserCreateRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	user := domain.Users(
		Name: req.Name,
	)

	user := service.UsersRepository.Save(ctx, tx, user)
	return helpers.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, req web.UserUpdateRequest) web.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	user, err := service.UsersRepository.FindById(ctx, tx, req.Id)
	helpers.PanicIfError(err)

	user.Name = req.Name

	user := service.UsersRepository.Update(ctx, tx, user)
	return helpers.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	user, err := service.UsersRepository.FindById(ctx, tx, req.Id)
	helpers.PanicIfError(err)

	service.UsersRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	user, err := service.UsersRepository.FindById(ctx, tx, userId)
	helpers.PanicIfError(err)

	return helpers.ToUserResponse(user)
}

func (service *UserServiceImpl) FindByAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	users := service.UsersRepository.FindAll(ctx, tx)
	return helpers.ToUserResponses(users)
}

package repositories

import (
	"context"
	"crud-test/helpers"
	"crud-test/models/domain"
	"database/sql"
	"errors"
)

type UsersRepositoryImpl struct {
	//
}

func (repository *UsersRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, users domain.Users) domain.Users {
	sql := "insert into users(name) values (?)"
	result, err := tx.ExecContext(ctx, sql, users.Name)
	helpers.PanicIfError(err)

	id, err := result.LastInsertId()
	helpers.PanicIfError(err)

	users.Id = int(id)
	return users
}

func (repository *UsersRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, users domain.Users) domain.Users {
	sql := "update users set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, users.Name, users.Id)
	helpers.PanicIfError(err)

	return users
}

func (repository *UsersRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, users domain.Users) {
	sql := "delete from users where id = ?"
	_, err := tx.ExecContext(ctx, sql, users.Id)
	helpers.PanicIfError(err)
}

func (repository *UsersRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, users.Id int) (domain.Users, error) {
	sql := "select id, name from users where id = ?"
	rows, err := tx.QueryContext(ctx, sql, users.Id)
	helpers.PanicIfError(err)

	users := domain.Users{}
	if rows.Next() {
		rows.Scan(&users.Id, &users.Name)
		return users, nil
	} else {
		return users, errors.New("User is not found.")
	}
}

func (repository *UsersRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, users domain.Users) {
	sql := "select id, name from users"
	rows, err := tx.QueryContext(ctx, sql)
	helpers.PanicIfError(err)

	var users []domain.Users
	for rows.Next() {
		user := domain.Users{}
		err := rows.Scan(&users.Id, &users.Name)
		helpers.PanicIfError(err)

		users = append(users, user)
	}
	return users
}

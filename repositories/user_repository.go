package repositories

import (
	"context"
	"crud-test/models/domain"
	"database/sql"
)

type UsersRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
	Update(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
	Delete(ctx context.Context, tx *sql.Tx, user domain.Users)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.Users, error)
	FindAll(ctx context.Context, tx *sql.Tx, user domain.Users) []domain.Users
}

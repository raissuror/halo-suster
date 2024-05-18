package repository

import (
	"context"
	"database/sql"
	"halo-suster/model/domain"
)

type UserRepo interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindByNip(ctx context.Context, tx *sql.Tx, phoneNumber string) (domain.User, error)
}

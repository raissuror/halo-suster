package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"halo-suster/model/domain"
	"time"
)

type UserRepoImpl struct {
}

func NewUserRepo() UserRepo {
	return &UserRepoImpl{}
}

func (repository *UserRepoImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	var exists bool
	errValidation := tx.QueryRow("SELECT exists(SELECT 1 FROM users WHERE nip=$1)", user.Nip).Scan(&exists)
	if errValidation != nil {
		fmt.Print("NIP not found")
	}
	if exists {
		fmt.Printf("NIP %s", user.Nip)
		return user, errors.New("NIP already exists")
	}
	sql := "INSERT INTO users(nip, name, password, created_at, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING id"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	insertedId := 0
	err := tx.QueryRowContext(ctx, sql, user.Nip, user.Name, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&insertedId)
	if err != nil {
		return user, errors.New("error save user")
	}
	user.Id = insertedId
	return user, nil
}

func (repository *UserRepoImpl) FindByNip(ctx context.Context, tx *sql.Tx, nip string) (domain.User, error) {
	var user domain.User
	sql := "SELECT id, name, nip, password, created_at, updated_at FROM users WHERE nip=$1"
	err := tx.QueryRowContext(ctx, sql, nip).Scan(&user.Id, &user.Name, &user.Nip, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

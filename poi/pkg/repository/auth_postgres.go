package repository

import (
	"context"
	"fmt"
	"restapi/internal/entity"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, age, gender, password) values($1, $2, $3, $4, $5) RETURNING id", UserTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Age, user.Gender, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id FROM  %s WHERE email=$1 AND password=$2", UserTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}

func (r *AuthPostgres) GetUsers(ctx context.Context) (users []*entity.User, err error) {
	query := fmt.Sprintf("SELECT * FROM %s", UserTable)
	err = r.db.Select(&users, query)

	return users, err
}

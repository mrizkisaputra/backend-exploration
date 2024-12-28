package repositories

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"golang-postgresql/model"
)

var (
	createUserQuery = "INSERT INTO users(name, email, age) VALUES($1, $2, $3)"
	findAllQuery    = "SELECT name, email, age FROM users"
)

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) FindAll(ctx context.Context) ([]model.User, error) {
	records, err := u.db.QueryContext(ctx, findAllQuery)
	if err != nil {
		return nil, err
	}

	var users []model.User
	for records.Next() {
		var user model.User

		if err := records.Scan(&user.Name, &user.Email, &user.Age); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func NewUserRepository(db *sql.DB) UserPostgresRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Insert(ctx context.Context, user model.User) error {
	tx, err := u.db.Begin()
	defer tx.Rollback()
	if err != nil {
		return errors.Wrap(err, "UserRepository.Insert.db.Begin")
	}

	_, err = tx.ExecContext(ctx, createUserQuery, user.Name, user.Email, user.Age)
	if err != nil {
		return errors.Wrap(err, "UserRepository.Insert.ExecContext")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "UserRepository.Insert.tx.Commit")
	}

	return nil
}

func (u *userRepository) Update(ctx context.Context, user model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) Delete(ctx context.Context, user model.User) error {
	//TODO implement me
	panic("implement me")
}

package repository

import (
	"context"
	"database/sql"
	"errors"
	"hms/user-service/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	query := `
		INSERT INTO users (email, password_hash, role)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	return r.db.QueryRow(query, user.Email, user.PasswordHash, user.Role).Scan(&user.ID)
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	query := `
		SELECT id, email, password_hash, role, is_active
		FROM users
		WHERE email = $1
	`
	user := &model.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.IsActive,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func GetUserFromSession(sessionID string) (int64, error) {
	val, err := RedisClient.Get(context.Background(), sessionID).Result()
	if err != nil {
		return 0, errors.New("session not found")
	}

	// convert string to int64
	return parseInt64(val), nil
}

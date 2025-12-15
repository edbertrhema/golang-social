package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (u *UserStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDurations)
	defer cancel()

	err := u.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		switch {
		case errors.As(err, new(*pq.Error)):
			// Cast the generic error to the specific pq.Error type
			pqErr := err.(*pq.Error)

			// 2. Check for the unique violation error code (23505)
			if pqErr.Code == "23505" {
				// You can inspect pqErr.Constraint to see which constraint failed (e.g., 'users_username_key')
				return ErrDuplicateKey
			}
		default:
			return err
		}
	}

	return nil
}

func (u UserStore) GetByID(ctx context.Context, id int) (*User, error) {
	query := `SELECT id, username, email, created_at FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDurations)
	defer cancel()

	var user User

	err := u.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Follower struct {
	UserID     int `json:"user_id"`
	FollowerID int `json:"follower_id"`
	CreatedAd  int `json:"created_added"`
}
type FollowerStore struct {
	db *sql.DB
}

func (f *FollowerStore) Follow(ctx context.Context, followerId, userID int) error {
	query := `INSERT INTO followers (user_id, follower_id) VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDurations)
	defer cancel()

	_, err := f.db.ExecContext(ctx, query, userID, followerId)
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
	return nil
}

func (f *FollowerStore) Unfollow(ctx context.Context, followerId, userID int) error {
	query := `DELETE FROM followers WHERE user_id = $1 AND follower_id = $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDurations)
	defer cancel()

	_, err := f.db.ExecContext(ctx, query, userID, followerId)
	if err != nil {
		return err
	}
	return nil
}

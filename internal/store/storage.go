package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound           = errors.New("Resource not found")
	ErrDuplicateKey       = errors.New("Duplicate Key")
	QueryTimeoutDurations = 5 * time.Second
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int) (*Post, error)
		Delete(context.Context, int) error
		Update(context.Context, *Post) error
		GetUserFeed(context.Context, int) ([]*PostWithMetadata, error)
	}
	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int) (*User, error)
	}
	Comments interface {
		GetByPostID(context.Context, int) ([]Comment, error)
	}
	Followers interface {
		Follow(context.Context, int, int) error
		Unfollow(context.Context, int, int) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}

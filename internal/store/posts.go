package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Post struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int       `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdateAt  string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO posts (content, title, user_id, tags) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	err := s.db.QueryRowContext(ctx, query, post.Content, post.Title, post.UserID, pq.Array(post.Tags)).Scan(&post.ID, &post.CreatedAt, &post.UpdateAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetByID(ctx context.Context, id int) (*Post, error) {
	query := `SELECT id,user_id,title,content, created_at, updated_at, tags  FROM posts WHERE id = $1`

	var post Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdateAt,
		pq.Array(&post.Tags),
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (s *PostStore) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM posts WHERE id = $1`

	result, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if row == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error {
	query := `UPDATE posts SET content = $1, title = $2 WHERE id = $3`

	result, err := s.db.ExecContext(ctx, query, post.Content, post.Title, post.ID)
	if err != nil {
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return ErrNotFound
	}
	return nil

}

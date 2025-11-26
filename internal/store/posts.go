package store

import (
	"context"
	"database/sql"
)

type Post struct {
	ID        int      `json: "user_id"`
	Content   string   `json: "content"`
	Title     string   `json: "title"`
	UserID    int      `json: "user_id"`
	Tags      []string `json: "tags"`
	CreatedAt string   `json: "created_at"`
	UpdateAt  string   `json: "update_at"`
}

type PostsStore struct {
	db *sql.DB
}

func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO posts (content, title, user_id, tags) VALUES ($1, $2, $3, $4) RETURNING id, created_at, update_at`

	s.db.QueryContext(ctx, query, post.Content, post.Title, post.UserID, post.Tags)
}

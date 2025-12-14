package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	Id        int    `json:"id"`
	PostID    int    `json:"post_id"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) GetByPostID(ctx context.Context, PostId int) ([]Comment, error) {
	query := `	SELECT c.id ,c."content",c.created_at, u.id, u.username  FROM comments c left join users u on c.user_id = u.id where c.post_id = $1 order by c.created_at desc`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDurations)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, PostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var c Comment
		c.User = User{}
		err := rows.Scan(&c.Id, &c.Content, &c.CreatedAt, &c.User.ID, &c.User.Username)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

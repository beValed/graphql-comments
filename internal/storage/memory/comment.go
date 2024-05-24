package memory

import (
	"graphql-comments/internal/app/comment"
	"sync"
)

type CommentRepository struct {
	mu       sync.Mutex
	comments map[uint]comment.Comment
	nextID   uint
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{
		comments: make(map[uint]comment.Comment),
		nextID:   1,
	}
}

func (r *CommentRepository) Create(c comment.Comment) (comment.Comment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	c.ID = r.nextID
	r.comments[r.nextID] = c
	r.nextID++
	return c, nil
}

func (r *CommentRepository) GetByPostID(postID uint) ([]comment.Comment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var comments []comment.Comment
	for _, c := range r.comments {
		if c.PostID == postID {
			comments = append(comments, c)
		}
	}
	return comments, nil
}

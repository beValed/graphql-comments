package postgres

import (
	"gorm.io/gorm"
	"graphql-comments/internal/app/comment"
)

type PostgresCommentRepository struct {
	db *gorm.DB
}

func NewPostgresCommentRepository(db *gorm.DB) *PostgresCommentRepository {
	return &PostgresCommentRepository{db: db}
}

func (r *PostgresCommentRepository) Create(c comment.Comment) (comment.Comment, error) {
	if err := r.db.Create(&c).Error; err != nil {
		return comment.Comment{}, err
	}
	return c, nil
}

func (r *PostgresCommentRepository) GetByPostID(postID uint) ([]comment.Comment, error) {
	var comments []comment.Comment
	if err := r.db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

package postgres

import (
	"gorm.io/gorm"
	"graphql-comments/internal/app/post"
)

type PostgresPostRepository struct {
	db *gorm.DB
}

func NewPostgresPostRepository(db *gorm.DB) *PostgresPostRepository {
	return &PostgresPostRepository{db: db}
}

func (r *PostgresPostRepository) Create(p post.Post) (post.Post, error) {
	if err := r.db.Create(&p).Error; err != nil {
		return post.Post{}, err
	}
	return p, nil
}

func (r *PostgresPostRepository) GetByID(id uint) (post.Post, error) {
	var p post.Post
	if err := r.db.First(&p, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return post.Post{}, post.ErrPostNotFound
		}
		return post.Post{}, err
	}
	return p, nil
}

func (r *PostgresPostRepository) GetAll() ([]post.Post, error) {
	var posts []post.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostgresPostRepository) Update(p post.Post) (post.Post, error) {
	if err := r.db.Save(&p).Error; err != nil {
		return post.Post{}, err
	}
	return p, nil
}

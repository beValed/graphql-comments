package post

import (
	"errors"
	"time"
)

var (
	ErrPostNotFound = errors.New("post not found")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePost(title, content string) (Post, error) {
	post := Post{
		Title:           title,
		Content:         content,
		CommentsEnabled: true,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	return s.repo.Create(post)
}

func (s *Service) GetPostByID(id uint) (Post, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetPosts() ([]Post, error) {
	return s.repo.GetAll()
}

func (s *Service) ToggleComments(id uint, enabled bool) (Post, error) {
	post, err := s.repo.GetByID(id)
	if err != nil {
		return Post{}, err
	}
	post.CommentsEnabled = enabled
	post.UpdatedAt = time.Now()
	return s.repo.Update(post)
}

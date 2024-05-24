package comment

import (
	"errors"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateComment(postID uint, parentID *uint, content string) (Comment, error) {
	if len(content) > 2000 {
		return Comment{}, errors.New("content length exceeds 2000 characters")
	}
	comment := Comment{PostID: postID, ParentID: parentID, Content: content}
	return s.repo.Create(comment)
}

func (s *Service) GetCommentsByPostID(postID uint) ([]Comment, error) {
	return s.repo.GetByPostID(postID)
}

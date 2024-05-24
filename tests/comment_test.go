package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"graphql-comments/internal/app/comment"
	"graphql-comments/internal/storage/memory"
)

func TestCreateComment(t *testing.T) {
	commentService := comment.NewService(nil)

	validComment, err := commentService.CreateComment(1, nil, "This is a valid comment")
	assert.NoError(t, err)
	assert.NotNil(t, validComment)

	invalidComment, err := commentService.CreateComment(1, nil, string(make([]byte, 2001)))
	assert.Error(t, err)
	assert.Nil(t, invalidComment)
}

func TestGetCommentsByPostID(t *testing.T) {
	repo := memory.NewCommentRepository()
	service := comment.NewService(repo)

	_, _ = service.CreateComment(1, nil, "Test Comment 1")
	_, _ = service.CreateComment(1, nil, "Test Comment 2")

	comments, err := service.GetCommentsByPostID(1)
	assert.NoError(t, err)
	assert.Len(t, comments, 2)
}

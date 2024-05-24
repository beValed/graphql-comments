package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"graphql-comments/internal/app/post"
	"graphql-comments/internal/storage/memory"
)

func TestCreatePost(t *testing.T) {
	repo := memory.NewPostRepository()
	service := post.NewService(repo)

	newPost, err := service.CreatePost("Test Title", "Test Content")
	assert.NoError(t, err)
	assert.NotNil(t, newPost)
	assert.Equal(t, "Test Title", newPost.Title)
	assert.Equal(t, "Test Content", newPost.Content)
}

func TestGetPosts(t *testing.T) {
	repo := memory.NewPostRepository()
	service := post.NewService(repo)

	_, _ = service.CreatePost("Test Title 1", "Test Content 1")
	_, _ = service.CreatePost("Test Title 2", "Test Content 2")

	posts, err := service.GetPosts()
	assert.NoError(t, err)
	assert.Len(t, posts, 2)
}

func TestToggleComments(t *testing.T) {
	repo := memory.NewPostRepository()
	service := post.NewService(repo)

	post, _ := service.CreatePost("Test Title", "Test Content")
	updatedPost, err := service.ToggleComments(post.ID, false)
	assert.NoError(t, err)
	assert.False(t, updatedPost.CommentsEnabled)
}

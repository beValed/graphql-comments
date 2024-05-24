package graph

import (
	"context"
	"graphql-comments/internal/app/comment"
	"graphql-comments/internal/app/post"
	"graphql-comments/internal/storage/memory"
	"graphql-comments/internal/storage/postgres"

	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm"
)

type Resolver struct {
	PostService    *post.Service
	CommentService *comment.Service
	DB             *gorm.DB
}

func NewResolver(db *gorm.DB, useInMemory bool) *Resolver {
	var postService *post.Service
	var commentService *comment.Service

	if useInMemory {
		postRepo := memory.NewPostRepository()
		commentRepo := memory.NewCommentRepository()
		postService = post.NewService(postRepo)
		commentService = comment.NewService(commentRepo)
	} else {
		postRepo := postgres.NewPostgresPostRepository(db)
		commentRepo := postgres.NewPostgresCommentRepository(db)
		postService = post.NewService(postRepo)
		commentService = comment.NewService(commentRepo)
	}

	return &Resolver{
		DB:             db,
		PostService:    postService,
		CommentService: commentService,
	}
}

type Config struct {
	Resolvers *Resolver
}

func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{Resolvers: cfg.Resolvers})
}

func (r *Resolver) Posts(ctx context.Context) ([]*post.Post, error) {
	posts, err := r.PostService.GetPosts()
	if err != nil {
		return nil, err
	}
	var gqlPosts []*post.Post
	for _, p := range posts {
		gqlPosts = append(gqlPosts, &p)
	}
	return gqlPosts, nil
}

func (r *Resolver) Post(ctx context.Context, id uint) (*post.Post, error) {
	post, err := r.PostService.GetPostByID(id)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *Resolver) CreatePost(ctx context.Context, title string, content string) (*post.Post, error) {
	post, err := r.PostService.CreatePost(title, content)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *Resolver) CreateComment(ctx context.Context, postID uint, parentID *uint, content string) (*comment.Comment, error) {
	comment, err := r.CommentService.CreateComment(postID, parentID, content)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *Resolver) ToggleComments(ctx context.Context, postID uint, enabled bool) (*post.Post, error) {
	post, err := r.PostService.ToggleComments(postID, enabled)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

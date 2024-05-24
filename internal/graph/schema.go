package graph

import (
	"context"
	"time"
)

type Post struct {
	ID              uint      `json:"id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	CommentsEnabled bool      `json:"commentsEnabled"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type Comment struct {
	ID        uint      `json:"id"`
	PostID    uint      `json:"postId"`
	ParentID  *uint     `json:"parentId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MutationResolver struct{ *Resolver }

func (r *MutationResolver) CreatePost(ctx context.Context, title string, content string) (*Post, error) {
	post, err := r.PostService.CreatePost(title, content)
	if err != nil {
		return nil, err
	}
	return &Post{
		ID:              post.ID,
		Title:           post.Title,
		Content:         post.Content,
		CommentsEnabled: post.CommentsEnabled,
		CreatedAt:       post.CreatedAt,
		UpdatedAt:       post.UpdatedAt,
	}, nil
}

func (r *MutationResolver) ToggleComments(ctx context.Context, postID uint, enable bool) (*Post, error) {
	post, err := r.PostService.ToggleComments(postID, enable)
	if err != nil {
		return nil, err
	}
	return &Post{
		ID:              post.ID,
		Title:           post.Title,
		Content:         post.Content,
		CommentsEnabled: post.CommentsEnabled,
		CreatedAt:       post.CreatedAt,
		UpdatedAt:       post.UpdatedAt,
	}, nil
}

func (r *MutationResolver) CreateComment(ctx context.Context, postID uint, parentID *uint, content string) (*Comment, error) {
	comment, err := r.CommentService.CreateComment(postID, parentID, content)
	if err != nil {
		return nil, err
	}
	return &Comment{
		ID:        comment.ID,
		PostID:    comment.PostID,
		ParentID:  comment.ParentID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}, nil
}

type QueryResolver struct{ *Resolver }

func (r *QueryResolver) Posts(ctx context.Context) ([]*Post, error) {
	posts, err := r.PostService.GetPosts()
	if err != nil {
		return nil, err
	}

	result := make([]*Post, len(posts))
	for i, post := range posts {
		result[i] = &Post{
			ID:              post.ID,
			Title:           post.Title,
			Content:         post.Content,
			CommentsEnabled: post.CommentsEnabled,
			CreatedAt:       post.CreatedAt,
			UpdatedAt:       post.UpdatedAt,
		}
	}

	return result, nil
}

func (r *QueryResolver) Post(ctx context.Context, id uint) (*Post, error) {
	post, err := r.PostService.GetPostByID(id)
	if err != nil {
		return nil, err
	}
	return &Post{
		ID:              post.ID,
		Title:           post.Title,
		Content:         post.Content,
		CommentsEnabled: post.CommentsEnabled,
		CreatedAt:       post.CreatedAt,
		UpdatedAt:       post.UpdatedAt,
	}, nil
}

func (r *QueryResolver) Comments(ctx context.Context, postID uint) ([]*Comment, error) {
	comments, err := r.CommentService.GetCommentsByPostID(postID)
	if err != nil {
		return nil, err
	}

	result := make([]*Comment, len(comments))
	for i, comment := range comments {
		result[i] = &Comment{
			ID:        comment.ID,
			PostID:    comment.PostID,
			ParentID:  comment.ParentID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}
	}

	return result, nil
}

type SubscriptionResolver struct{ *Resolver }

func (r *SubscriptionResolver) CommentAdded(ctx context.Context, postID uint) (<-chan *Comment, error) {
	commentChan := make(chan *Comment)

	go func() {
		defer close(commentChan)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				comments, err := r.CommentService.GetCommentsByPostID(postID)
				if err == nil {
					for _, comment := range comments {
						commentChan <- &Comment{
							ID:        comment.ID,
							PostID:    comment.PostID,
							ParentID:  comment.ParentID,
							Content:   comment.Content,
							CreatedAt: comment.CreatedAt,
							UpdatedAt: comment.UpdatedAt,
						}
					}
				}
				time.Sleep(5 * time.Second)
			}
		}
	}()

	return commentChan, nil
}

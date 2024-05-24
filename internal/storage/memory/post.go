package memory

import (
	"graphql-comments/internal/app/post"
	"sync"
)

type PostRepository struct {
	mu     sync.Mutex
	posts  map[uint]post.Post
	nextID uint
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		posts:  make(map[uint]post.Post),
		nextID: 1,
	}
}

func (r *PostRepository) Create(p post.Post) (post.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p.ID = r.nextID
	r.posts[r.nextID] = p
	r.nextID++
	return p, nil
}

func (r *PostRepository) GetByID(id uint) (post.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, exists := r.posts[id]
	if !exists {
		return post.Post{}, post.ErrPostNotFound
	}
	return p, nil
}

func (r *PostRepository) GetAll() ([]post.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	posts := make([]post.Post, 0, len(r.posts))
	for _, p := range r.posts {
		posts = append(posts, p)
	}
	return posts, nil
}

func (r *PostRepository) Update(p post.Post) (post.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.posts[p.ID]
	if !exists {
		return post.Post{}, post.ErrPostNotFound
	}

	r.posts[p.ID] = p
	return p, nil
}

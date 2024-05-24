package post

type Repository interface {
	Create(Post) (Post, error)
	GetByID(uint) (Post, error)
	GetAll() ([]Post, error)
	Update(Post) (Post, error)
}

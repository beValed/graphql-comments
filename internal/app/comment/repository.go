package comment

type Repository interface {
	Create(comment Comment) (Comment, error)
	GetByPostID(postID uint) ([]Comment, error)
}

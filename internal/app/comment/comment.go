package comment

import "time"

type Comment struct {
	ID        uint
	PostID    uint
	ParentID  *uint
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

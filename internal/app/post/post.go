package post

import "time"

type Post struct {
	ID              uint
	Title           string
	Content         string
	CommentsEnabled bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

package postmodule

import "time"

type Post struct {
	EntityName      string
	Id              int32
	UserId          int32
	Content         string
	CreatedAt       time.Time
	LatestUpdatedAt time.Time
	DeletedAt       time.Time
}

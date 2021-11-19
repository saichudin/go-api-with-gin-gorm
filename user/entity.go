package user

import "time"

type User struct {
	Id			int
	Username	string
	Email		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
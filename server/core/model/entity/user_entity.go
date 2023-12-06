package entity

import "time"

type User struct {
	Id        *int
	Name      *string
	Age       *int
	Country   *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (User) TableName() string {
	return "user"
}

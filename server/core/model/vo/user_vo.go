package vo

import "time"

type User struct {
	Id        *int       `json:"id"`
	Name      *string    `json:"name"`
	Age       *int       `json:"age"`
	Country   *string    `json:"country"`
	CreatedAt *time.Time `json:"created_at"`
}

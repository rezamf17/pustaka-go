package book

import "time"

type Entity struct {
	Id          int
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

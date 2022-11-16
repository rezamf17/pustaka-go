package book

import (
	"encoding/json"
)

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,numeric"`
	Description string      `json:"description" binding:"required"`
	Rating      int         `json:"rating" binding:"required,numeric"`
	Discount    int         `json:"discount" binding:"required,numeric"`
}

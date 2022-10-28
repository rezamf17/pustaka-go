package book

import (
	"encoding/json"
)

type PostBooks struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,numeric"`
}

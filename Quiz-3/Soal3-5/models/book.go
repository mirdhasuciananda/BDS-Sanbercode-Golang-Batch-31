package models

import "time"

type Book struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	ReleaseYear uint      `json:"release_year"`
	Price       string    `json:"price"`
	TotalPage   uint      `json:"total_page"`
	Thickness   string    `json:"thickness"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CategoryId  uint      `json:"category_id"`
}

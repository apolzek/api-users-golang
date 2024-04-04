package response

import "time"

type CategoryResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ManyCategoriesResponse struct {
	Categories []CategoryResponse `json:"categories"`
}

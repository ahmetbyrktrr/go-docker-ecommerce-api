package model

type Product struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	Size       string  `json:"size"`
	Color      string  `json:"color"`
	CategoryID int64   `json:"category_id"`
}

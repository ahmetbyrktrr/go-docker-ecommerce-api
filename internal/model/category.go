package model

type Category struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products,omitempty"`
}

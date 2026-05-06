package models

type Product struct {
	ID          int      `db:"id" json:"id"`
	Title       string   `db:"title" json:"title"`
	Description string   `db:"description" json:"description"`
	Price       float64  `db:"price" json:"price"`
	Currency    string   `db:"currency" json:"currency"`
	Stock       int      `db:"stock" json:"stock"`
	CategoryID  *int64   `db:"category_id" json:"category_id"`
	ImageURL    string   `db:"image_url" json:"image_url"`
	Category    *Category `db:"-" json:"category,omitempty"`
}


package models

type UpdateProduct struct {
	Title       string `json:"title"`
	ID          uint   `json:"id"`
	CategoryID  int    `json:"category_id"`
	SellerID    int    `json:"seller_id"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
}

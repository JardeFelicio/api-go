package model

type Product struct {
	Id          int     `json: "productId"`
	Name        string  `json: "name"`
	Description string  `json: "description"`
	Price       float64 `json: "price"`
	Active      bool    `json: "active"`
	Stock       int     `json: "stock"`
	Category_id int     `json: "category_id"`
}

package model

var prodCategories = [10]string{
		"Clothing",
    "Accessories",
    "Footwear",
    "Beverages"}


type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Sku string `json:"sku"`
	Category string `json:"category"`
	ImageUrl string `json:"imageUrl"`
	Notes string `json:"notes"`
	Price int `json:"price"`
	Stock int `json:"stock"`
	Location string `json:"location"`
	IsAvailable bool `json:"isAvailable"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
}

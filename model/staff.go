package model

type Staff struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"email"`
	Password    string `json:"password"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
}

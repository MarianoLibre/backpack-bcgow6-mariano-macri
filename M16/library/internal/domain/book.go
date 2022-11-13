package domain

type Book struct {
	Id       int    `json:"id"`
	Title    string `json:"title" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

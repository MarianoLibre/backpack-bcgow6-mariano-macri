package domain

type Book struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Quantity int `json:"quantity"`
}

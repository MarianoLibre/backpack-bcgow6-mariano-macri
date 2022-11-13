package domain

type Loan struct {
	Id     int `json:"id"`
	BookId int `json:"book_id" binding:"required"`
	UserId int `json:"user_id" binding:"required"`
}

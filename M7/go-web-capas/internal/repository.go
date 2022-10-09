package internal


type Product struct {
    Id          int         `json:"id"`
    Name        string      `json:"name"`
    Colour      string      `json:"colour"`
    Price       float64     `json:"price"`
    Stock       int         `json:"stock"`
    Code        string      `json:"code"`
    Published   bool        `json:"published"`
    CreatedAt   string      `json:"created-at"`
}

var ps []Product
var lastID int

type Repository interface{
   GetAll() ([]Product, error)
   Store(id int, name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error)
   LastID() (int, error)
}

type repository struct {}

func NewRepository() Repository {
   return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
   return ps, nil
}

func (r *repository) LastID() (int, error) {
   return lastID, nil
}

func (r *repository) Store(id int, name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error) {
   p := Product{id, name, colour, price, stock, code, published, createdAt}
   ps = append(ps, p)
   lastID = p.Id
   return p, nil
}

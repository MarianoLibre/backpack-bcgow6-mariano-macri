package products

import (
	"fmt"

	"github.com/MarianoLibre/go-web-capas/pkg/store"
)

type Product struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Colour    string  `json:"colour"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      string  `json:"code"`
	Published bool    `json:"published"`
	CreatedAt string  `json:"created-at"`
}

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error)
	LastID() (int, error)
	Update(id int, name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error)
	UpdateName(id int, name string) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	lastId := 0
	for _, p := range ps {
		if p.Id > lastId {
			lastId = p.Id
		}
	}

	return lastId, nil
}

func (r *repository) Store(id int, name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error) {
	p := Product{id, name, colour, price, stock, code, published, createdAt}

	var ps []Product
	r.db.Read(&ps)
	ps = append(ps, p)
	//fmt.Println("REPOSITORY>>> ", id, name, colour, code, createdAt)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) Update(id int, name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error) {
	p := Product{Name: name, Colour: colour, Code: code, CreatedAt: createdAt, Stock: stock, Price: price, Published: published}
	updated := false
	var ps []Product
	r.db.Read(&ps)
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Product '%d' not found", id)
	}
	return p, nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var p Product
	updated := false
	var ps []Product
	r.db.Read(&ps)
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Name = name
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Product '%d' not found", id)
	}
	return p, nil
}

func (r *repository) UpdateNameAndPrice(id int, name string, price float64) (Product, error) {
	var p Product
	updated := false
	var ps []Product
	r.db.Read(&ps)
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Name = name
			ps[i].Price = price
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Product '%d' not found", id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	var ps []Product
	r.db.Read(&ps)
	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Product '%d' not found", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	if r.db.Write(ps) != nil {
		return fmt.Errorf("error while deleting product '%d'", id)
	}
	return nil
}

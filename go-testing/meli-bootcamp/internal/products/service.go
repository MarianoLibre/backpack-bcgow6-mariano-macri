package products

//import "fmt"

type Service interface {
	GetAll() ([]Product, error)
	Store(name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error)
	Update(id int, name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, name, colour, code, createdAt, stock, price, published)
	if err != nil {
		return Product{}, err
	}

	//fmt.Println("SERVICE>>> !", producto)
	return producto, nil
}

func (s *service) Update(id int, name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error) {
	return s.repository.Update(id, name, colour, code, createdAt, stock, price, published)
}

func (s *service) UpdateNameAndPrice(id int, name string, price float64) (Product, error) {
	return s.repository.UpdateNameAndPrice(id, name, price)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

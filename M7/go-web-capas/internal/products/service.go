package internal

type Service interface {
   GetAll() ([]Product, error)
   Store(name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error)
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
   if err != nil{
       return nil, err
   }

   return ps, nil
}

func (s *service) Store(name, colour, code, createdAt string, stock int, price float64, published bool) (Product, error) {
   lastID, err := s.repository.LastID()
   if err != nil{
       return Product{}, err
   }

   lastID++

   producto, err := s.repository.Store(lastID, name, colour, code, createdAt, stock, price, published)
   if err != nil{
       return Product{}, err
   }

   return producto, nil
}

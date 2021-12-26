package productos

type Service interface {
	GetAll() ([]Producto, error)
	Store(p Producto) (Producto, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetAll() ([]Producto, error) {
	productos, err := s.repository.GetAll()
	return productos, err
}

func (s *service) Store(p Producto) (Producto, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Producto{}, err
	}

	lastID++
	p.Id = lastID
	producto, err := s.repository.Store(p)
	if err != nil {
		return Producto{}, err
	}

	return producto, nil
}

package productos

type Service interface {
	GetAll() ([]Producto, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll() ([]Producto, error) {
	productos, err := s.repo.GetAll()
	return productos, err
}

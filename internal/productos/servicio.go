package productos

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAll() ([]Producto, error)
	Store(p Producto) (Producto, error)
	Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Producto, error)
	Delete(id int) error
	UpdateNamePrice(ctx *gin.Context, nombre string, precio float64) (Producto, error)
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

func (s *service) Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Producto, error) {
	prod := Producto{Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaDeCreacion: fechaDeCreacion}
	p, err := s.repository.Update(id, prod)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateNamePrice(ctx *gin.Context, nombre string, precio float64) (Producto, error) {
	return s.repository.UpdateNamePrice(ctx, nombre, precio)
}

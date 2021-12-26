package productos

import (
	"encoding/json"
	"os"
)

type Repository interface {
	GetAll() ([]Producto, error)
	Store(Producto) (Producto, error)
	LastID() (int, error)
}

type Producto struct {
	Id              int     `json:"id"`
	Nombre          string  `json:"nombre" validate:"required"`
	Color           string  `json:"color" validate:"required"`
	Precio          float64 `json:"precio" validate:"required"`
	Stock           int     `json:"stock" validate:"required"`
	Codigo          string  `json:"codigo" validate:"required"`
	Publicado       bool    `json:"publicado" validate:"required"`
	FechaDeCreacion string  `json:"fechaDeCreacion" validate:"required"`
}

var listaProductos []Producto
var lastID int

type repository struct{}

func NewRepository() Repository {
	products, err := readFile()
	if err != nil {
		return &repository{}
	}
	listaProductos = products
	return &repository{}
}

func (r *repository) GetAll() ([]Producto, error) {
	return listaProductos, nil
}

func (r *repository) Store(p Producto) (Producto, error) {
	listaProductos = append(listaProductos, p)
	return p, nil
}

func (r *repository) LastID() (int, error) {
	lastID = listaProductos[len(listaProductos)-1].Id
	return lastID, nil
}

func readFile() ([]Producto, error) {
	productos := []Producto{}
	jsonData, err := os.ReadFile("./products.json")
	if err != nil {
		return productos, err
	}

	errJson := json.Unmarshal(jsonData, &productos)
	if errJson != nil {
		return productos, err
	}

	listaProductos = productos
	return productos, nil
}

package productos

import (
	"encoding/json"
	"errors"
	"os"
)

type Repository interface {
	GetAll() ([]Producto, error)
	Store(Producto) (Producto, error)
	LastID() (int, error)
	Update(id int, p Producto) (Producto, error)
	Delete(id int) error
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

func (r *repository) Update(id int, p Producto) (Producto, error) {
	var actualizado bool
	for i, producto := range listaProductos {
		if producto.Id == id {
			p.Id = id
			listaProductos[i] = p
			actualizado = true
		}
	}
	if !actualizado {
		return p, errors.New("no se ha encontrado el registro especificado")
	}
	return p, nil
}

func (r *repository) LastID() (int, error) {
	lastID = listaProductos[len(listaProductos)-1].Id
	return lastID, nil
}

func (r *repository) Delete(id int) error {
	var deleted bool
	for i, prod := range listaProductos {
		if prod.Id == id {
			listaProductos[i].Publicado = false
			deleted = true
		}
	}
	if !deleted {
		return errors.New("no existe el elemento especificado")
	}
	return nil
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

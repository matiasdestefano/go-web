package productos

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matiasdestefano/go-web/pkg/store"
)

type Repository interface {
	GetAll() ([]Producto, error)
	Store(Producto) (Producto, error)
	LastID() (int, error)
	Update(id int, p Producto) (Producto, error)
	Delete(id int) error
	UpdateNamePrice(ctx *gin.Context, nombre string, precio float64) (Producto, error)
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

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Producto, error) {
	var listaProductos []Producto
	err := r.db.Read(&listaProductos)
	if err != nil {
		return []Producto{}, err
	}
	return listaProductos, nil
}

func (r *repository) Store(p Producto) (Producto, error) {
	var listaProductos []Producto
	err := r.db.Read(&listaProductos)
	if err != nil {
		return Producto{}, err
	}
	p.Id, err = r.LastID()
	p.Id++
	if err != nil {
		return Producto{}, err
	}
	listaProductos = append(listaProductos, p)
	err = r.db.Write(&listaProductos)
	if err != nil {
		return Producto{}, err
	}
	return p, nil
}

func (r *repository) Update(id int, p Producto) (Producto, error) {
	var actualizado bool
	var listaProductos []Producto
	err := r.db.Read(&listaProductos)
	if err != nil {
		return Producto{}, err
	}
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

	err = r.db.Write(&listaProductos)
	if err != nil {
		return Producto{}, err
	}
	return p, nil
}

func (r *repository) LastID() (int, error) {
	var listaProductos []Producto
	err := r.db.Read(&listaProductos)
	if err != nil {
		return 0, err
	}
	return listaProductos[len(listaProductos)-1].Id, nil
}

func (r *repository) Delete(id int) error {
	var deleted bool
	var listaProductos []Producto
	err := r.db.Read(&listaProductos)
	if err != nil {
		return err
	}
	for i, prod := range listaProductos {
		if prod.Id == id {
			listaProductos[i].Publicado = false
			deleted = true
		}
	}
	if !deleted {
		return errors.New("no existe el elemento especificado")
	}
	err = r.db.Write(&listaProductos)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateNamePrice(ctx *gin.Context, nombre string, precio float64) (Producto, error) {
	var listaProductos []Producto
	var updated bool
	var prod Producto
	err := r.db.Read(&listaProductos)
	if err != nil {
		return Producto{}, err
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	for i, p := range listaProductos {
		if p.Id == id {
			listaProductos[i].Nombre = nombre
			listaProductos[i].Precio = precio
			prod = listaProductos[i]
			updated = true
		}
	}
	if !updated {
		return Producto{}, errors.New("el elemento especificado no existe")
	}
	err = r.db.Write(&listaProductos)
	if err != nil {
		return Producto{}, err
	}
	return prod, nil
}

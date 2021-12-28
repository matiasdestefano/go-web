package controlador

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/matiasdestefano/go-web/internal/productos"
	"github.com/matiasdestefano/go-web/pkg/web"
)

type request struct {
	Id              int     `json:"id"`
	Nombre          string  `json:"nombre" validate:"required"`
	Color           string  `json:"color" validate:"required"`
	Precio          float64 `json:"precio" validate:"required"`
	Stock           int     `json:"stock" validate:"required"`
	Codigo          string  `json:"codigo" validate:"required"`
	Publicado       bool    `json:"publicado" validate:"required"`
	FechaDeCreacion string  `json:"fechaDeCreacion" validate:"required"`
}

type Producto struct {
	service productos.Service
}

func NewProducto(p productos.Service) *Producto {
	return &Producto{
		service: p,
	}
}

func (prod *Producto) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validateToken(ctx.GetHeader("token"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}

		listaProductos, err := prod.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, listaProductos, ""))
	}
}

func (prod *Producto) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validateToken(ctx.GetHeader("token"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}

		var req productos.Producto
		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		v := validator.New()
		err = v.Struct(req)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				mensaje := "el campo " + err.Field() + " es requerido"
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, mensaje))
				return
			}
		}

		p, err := prod.service.Store(req)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

func (prod *Producto) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validateToken(ctx.GetHeader("token"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}
		var existe bool
		productos, err := prod.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		idProd, _ := strconv.Atoi(ctx.Param("id"))
		if idProd != 0 {
			for _, p := range productos {
				if p.Id == idProd {
					ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
					return
				}
			}
			existe = false
		}
		if !existe {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "el producto no existe"))
		}
	}
}

func (p *Producto) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validateToken(ctx.GetHeader("token"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}
		idProd, _ := strconv.Atoi(ctx.Param("id"))
		var req request
		ctx.ShouldBindJSON(&req)
		prod, err := p.service.Update(idProd, req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(http.StatusOK, prod, ""))
	}

}

func (p *Producto) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validateToken(ctx.GetHeader("token"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}
		prodId, _ := strconv.Atoi(ctx.Param("id"))
		var req request
		ctx.ShouldBindJSON(&req)
		err = p.service.Delete(prodId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(200, "elemento eliminado correctamente")
	}
}

func (p *Producto) UpdateNamePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validateToken(ctx.GetHeader("token"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}

		var req request
		ctx.ShouldBindJSON(&req)

		if req.Nombre == "" {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "no se definió un nombre"))
			return
		}

		if req.Precio <= 0.0 {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "no se definió un precio válido"))
			return
		}

		prod, err := p.service.UpdateNamePrice(ctx, req.Nombre, req.Precio)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, prod, ""))
	}
}

func validateToken(token string) error {
	if token != os.Getenv("TOKEN") {
		return errors.New("el token es inválido")
	}
	return nil
}

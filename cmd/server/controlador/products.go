package controlador

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/matiasdestefano/go-web/internal/productos"
)

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
		//Validacion del Token
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		listaProductos, err := prod.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		}
		ctx.JSON(http.StatusOK, listaProductos)
	}
}

func (prod *Producto) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		var req productos.Producto
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}

		v := validator.New()
		err = v.Struct(req)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				mensaje := "el campo " + err.Field() + " es requerido"
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": mensaje,
				})
				return
			}
		}

		p, err := prod.service.Store(req)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, p)
	}
}

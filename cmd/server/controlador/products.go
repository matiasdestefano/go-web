package controlador

import (
	"github.com/gin-gonic/gin"
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

func (c *Producto) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		producto, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(400, err)
		}
		ctx.JSON(200, producto)
	}
}

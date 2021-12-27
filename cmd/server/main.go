package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matiasdestefano/go-web/cmd/server/controlador"
	"github.com/matiasdestefano/go-web/internal/productos"
)

func main() {

	productoRepository := productos.NewRepository()
	productoService := productos.NewService(productoRepository)
	productoController := controlador.NewProducto(productoService)

	r := gin.Default()
	productoRouter := r.Group("/productos")
	productoRouter.GET("/", productoController.GetAll())
	productoRouter.GET("/:id", productoController.GetByID())
	productoRouter.POST("/", productoController.Store())

	r.Run() // localhost:8080
}

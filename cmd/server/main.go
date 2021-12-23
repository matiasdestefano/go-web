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
	pr := r.Group("/productos")
	pr.GET("/", productoController.GetAll())

	r.Run() // localhost:8080
}

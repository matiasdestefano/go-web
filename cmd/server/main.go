package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matiasdestefano/go-web/cmd/server/controlador"
	"github.com/matiasdestefano/go-web/internal/productos"
	"github.com/matiasdestefano/go-web/pkg/store"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	db := store.New(store.FileType, "./products.json")
	productoRepository := productos.NewRepository(db)
	productoService := productos.NewService(productoRepository)
	productoHandler := controlador.NewProducto(productoService)

	r := gin.Default()
	productoRouter := r.Group("/productos")
	productoRouter.GET("/", productoHandler.GetAll())
	productoRouter.GET("/:id", productoHandler.GetByID())
	productoRouter.POST("/", productoHandler.Store())
	productoRouter.PUT("/:id", productoHandler.Update())
	productoRouter.DELETE("/:id", productoHandler.Delete())
	productoRouter.PATCH("/:id", productoHandler.UpdateNamePrice())

	r.Run() // localhost:8080
}

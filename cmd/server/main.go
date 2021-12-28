package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matiasdestefano/go-web/cmd/server/controlador"
	"github.com/matiasdestefano/go-web/docs"
	"github.com/matiasdestefano/go-web/internal/productos"
	"github.com/matiasdestefano/go-web/pkg/store"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//@title MELI Bootcamp API
//@version 1.0
//@description This API handles MELI productos
//@termsOfServices https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

//@contact.name API Support
//@contact.url https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

//@license.name Apache 2.0
//@license.url http://www.apache.org/licenses/LICENSE-2.0.html
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	productoRouter := r.Group("/productos")
	productoRouter.GET("/", productoHandler.GetAll())
	productoRouter.GET("/:id", productoHandler.GetByID())
	productoRouter.POST("/", productoHandler.Store())
	productoRouter.PUT("/:id", productoHandler.Update())
	productoRouter.DELETE("/:id", productoHandler.Delete())
	productoRouter.PATCH("/:id", productoHandler.UpdateNamePrice())

	r.Run() // localhost:8080
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type producto struct {
	Id              int     `json:"id"`
	Nombre          string  `json:"nombre" validate:"required"`
	Color           string  `json:"color" validate:"required"`
	Precio          float64 `json:"precio" validate:"required"`
	Stock           int     `json:"stock" validate:"required"`
	Codigo          string  `json:"codigo" validate:"required"`
	Publicado       bool    `json:"publicado" validate:"required"`
	FechaDeCreacion string  `json:"fechaDeCreacion" validate:"required"`
}

var productos []producto

func main() {

	productos = obtenerProductos()
	router := gin.Default()

	router.GET("/user/:name", ShowName)

	productosRouter := router.Group("/productos")
	productosRouter.GET("/", GetAll)
	productosRouter.GET("/:id", handlerBuscarProductoID)
	productosRouter.POST("/", Guardar())

	router.Run() // localhost:8080
}

func GetAll(c *gin.Context) {
	prodID, _ := strconv.Atoi(c.Query("id"))
	if prodID != 0 {
		for _, p := range productos {
			if p.Id == prodID {
				c.JSON(http.StatusOK, gin.H{
					"producto": p,
				})
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"productos": productos,
	})
}

func ShowName(c *gin.Context) {
	name := c.Param("name")
	texto := fmt.Sprintf("Hola, %s", name)
	c.JSON(http.StatusOK, gin.H{"message": texto})
}

func handlerBuscarProductoID(context *gin.Context) {
	var existe bool
	productos := obtenerProductos()
	idProd, _ := strconv.Atoi(context.Param("id"))
	if idProd != 0 {
		for _, p := range productos {
			if p.Id == idProd {
				context.JSON(http.StatusOK, gin.H{
					"producto": p,
				})
			}
		}
		existe = false
	}
	if !existe {
		context.JSON(http.StatusNotFound, "El producto no existe.")
	}
}

func obtenerProductos() []producto {
	productos := []producto{}
	jsonData, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	errJson := json.Unmarshal(jsonData, &productos)
	if errJson != nil {
		fmt.Println(errJson.Error())
		os.Exit(1)
	}

	return productos
}

func Guardar() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//VALIDAR TOKEN
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}

		var req producto
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
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

		lastID := getLastProductID()
		req.Id = lastID + 1
		//TODO: Deberia agregar dicho producto a products.json
		productos = append(productos, req)
	}
}

func getLastProductID() int {
	return productos[len(productos)-1].Id
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type producto struct {
	Id              string  `json:"id"`
	Nombre          string  `json:"nombre"`
	Color           string  `json:"color"`
	Precio          float64 `json:"precio"`
	Stock           int     `json:"stock"`
	Codigo          string  `json:"codigo"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fechaDeCreacion"`
}

func main() {
	router := gin.Default()

	router.GET("/user/:name", ShowName)

	router.GET("/productos", GetAll)
	router.GET("/productos/:id", handlerBuscarProductoID)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetAll(c *gin.Context) {
	productos := obtenerProductos()
	prodID := c.Query("id")
	if prodID != "" {
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
	idProd := context.Param("id")
	if idProd != "" {
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

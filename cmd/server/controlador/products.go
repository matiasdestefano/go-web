package controlador

import (
	"log"
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

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /productos [get]
func (prod *Producto) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		listaProductos, err := prod.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, listaProductos, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /productos [post]
func (prod *Producto) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req productos.Producto
		err := ctx.ShouldBindJSON(&req)
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

// GetProductByID godoc
// @Summary Get product by ID
// @Tags Products
// @Description get product by id
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /productos/id [get]
func (prod *Producto) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

// UpdateProduct godoc
// @Summary Updates a product
// @Tags Products
// @Description updates a product
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "product to update"
// @Success 200 {object} web.Response
// @Router /productos/id [put]
func (p *Producto) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

// DeleteProduct godoc
// @Summary Delete product
// @Tags Products
// @Description delete product by id
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /productos/id [delete]
func (p *Producto) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		prodId, _ := strconv.Atoi(ctx.Param("id"))
		var req request
		ctx.ShouldBindJSON(&req)
		err := p.service.Delete(prodId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(200, "elemento eliminado correctamente")
	}
}

// UpdateNameAndPriceOfProduct godoc
// @Summary Updates Name and Price of product
// @Tags Products
// @Description updates name and price of product
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Name and price value to update"
// @Success 200 {object} web.Response
// @Router /productos/id [patch]
func (p *Producto) UpdateNamePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("Por favor establecer el valor de la variable de entorno TOKEN")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "se requiere el token de la api"))
			return
		}

		if token != requiredToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "el token es invalido"))
			return
		}
		ctx.Next()
	}
}

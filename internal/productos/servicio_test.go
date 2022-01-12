package productos

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/matiasdestefano/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceGetAll(t *testing.T) {
	productosTest := []Producto{{

		Id:              1,
		Nombre:          "Test",
		Color:           "Verde",
		Precio:          15000,
		Stock:           100,
		Codigo:          "TST",
		Publicado:       true,
		FechaDeCreacion: "23-01-2022",
	},
		{

			Id:              2,
			Nombre:          "Test 2",
			Color:           "Azul",
			Precio:          10,
			Stock:           100,
			Codigo:          "TST2",
			Publicado:       true,
			FechaDeCreacion: "23-01-2022",
		}}

	dataJson, _ := json.Marshal(productosTest)
	dbMock := store.Mock{
		Data: dataJson,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	repo := NewRepository(&storeStub)
	servicio := NewService(repo)

	resultado, err := servicio.GetAll()
	assert.Equal(t, 2, len(resultado), "should be equal")
	assert.Nil(t, err)
}

func TestServiceStore(t *testing.T) {
	productoTest := []Producto{{
		Nombre:          "Test",
		Color:           "Verde",
		Precio:          15000,
		Stock:           100,
		Codigo:          "TST",
		Publicado:       true,
		FechaDeCreacion: "23-01-2022",
	},
	}

	dataJson, _ := json.Marshal(productoTest)

	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	repo := NewRepository(&storeStub)
	servicio := NewService(repo)

	resultado, err := servicio.Store(productoTest[0])
	productoTest[0].Id = 1
	assert.Equal(t, productoTest[0], resultado, "should be equal")
	assert.Nil(t, err, "should be nil")
}

func TestServiceUpdate(t *testing.T) {

	productoTest := []Producto{{
		Id:              1,
		Nombre:          "Test",
		Color:           "Verde",
		Precio:          15000,
		Stock:           100,
		Codigo:          "TST",
		Publicado:       true,
		FechaDeCreacion: "23-01-2022",
	},
	}

	dataJson, _ := json.Marshal(productoTest)

	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	repo := NewRepository(&storeStub)
	servicio := NewService(repo)

	resultado, err := servicio.Update(1, "Test Update", "Verde", 15000, 99, "TST", true, "23-01-2022")
	assert.Equal(t, "Test Update", resultado.Nombre, "should be equal")
	assert.Nil(t, err)
}

func TestServiceUpdateNamePrice(t *testing.T) {
	productoTest := []Producto{{
		Id:              1,
		Nombre:          "Test",
		Color:           "Verde",
		Precio:          15000,
		Stock:           100,
		Codigo:          "TST",
		Publicado:       true,
		FechaDeCreacion: "23-01-2022",
	},
	}

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	dataJson, _ := json.Marshal(productoTest)

	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	repo := NewRepository(&storeStub)
	servicio := NewService(repo)

	resultado, err := servicio.UpdateNamePrice(c, "Test Update", 7000)
	assert.Equal(t, "Test Update", resultado.Nombre, "should be equal")
	assert.Nil(t, err)
}

func TestServiceDelete(t *testing.T) {

	productoTest := []Producto{{
		Id:              1,
		Nombre:          "Test",
		Color:           "Verde",
		Precio:          15000,
		Stock:           100,
		Codigo:          "TST",
		Publicado:       true,
		FechaDeCreacion: "23-01-2022",
	},
	}

	dataJson, _ := json.Marshal(productoTest)

	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	repo := NewRepository(&storeStub)
	servicio := NewService(repo)

	err := servicio.Delete(1)
	err2 := servicio.Delete(2)

	assert.Nil(t, err)
	assert.NotNil(t, err2, "shouldn't be nil")
}

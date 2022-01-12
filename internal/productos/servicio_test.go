package productos

import (
	"encoding/json"
	"testing"

	"github.com/matiasdestefano/go-web/pkg/store"
	"github.com/stretchr/testify/assert"
)

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

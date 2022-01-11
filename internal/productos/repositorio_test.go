package productos

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const prods = `[
	{
	 "id": 1,
	 "nombre": "Microondas",
	 "color": "blanco",
	 "precio": 30000,
	 "stock": 10,
	 "codigo": "M1CR0",
	 "publicado": false,
	 "fechaDeCreacion": "31-01-2021"
	},
	{
	 "id": 2,
	 "nombre": "Playstation 5",
	 "color": "blanco",
	 "precio": 200000,
	 "stock": 100,
	 "codigo": "PS5",
	 "publicado": true,
	 "fechaDeCreacion": "22-12-2021"
	}
]`

type StubDB struct{}

type MockDB struct {
	producto     string
	readExecuted bool
}

func NewStub() *StubDB {
	return &StubDB{}
}

func NewMock() *MockDB {
	return &MockDB{producto: `[{
		"id": 1,
		"nombre": "Before update",
		"color": "blanco",
		"precio": 30000,
		"stock": 10,
		"codigo": "M1CR0",
		"publicado": false,
		"fechaDeCreacion": "31-01-2021"
	   }]`,
		readExecuted: false}
}

func (s *StubDB) Read(data interface{}) error {
	file := []byte(prods)
	return json.Unmarshal(file, &data)
}
func (s *StubDB) Write(data interface{}) error {
	return nil
}

func (m *MockDB) Read(data interface{}) error {
	m.readExecuted = true
	file := []byte(m.producto)
	return json.Unmarshal(file, &data)
}
func (m *MockDB) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//Arrange
	stub := NewStub()
	r := NewRepository(stub)

	//Act
	lista, err := r.GetAll()

	//Assert
	assert.Equal(t, 2, len(lista), "should be the same")
	assert.Nil(t, err, "should be nil")
}

func TestUpdateNamed(t *testing.T) {
	//Arrange
	mock := NewMock()
	r := NewRepository(mock)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	//Act
	resultado, err := r.UpdateNamePrice(c, "After update", 30000)

	//assert
	assert.True(t, mock.readExecuted, "should be true")
	assert.Equal(t, "After update", resultado.Nombre, "should be equal")
	assert.Nil(t, err, "should be nil")
}

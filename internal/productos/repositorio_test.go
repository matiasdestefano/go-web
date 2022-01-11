package productos

import (
	"encoding/json"
	"testing"

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

func NewStub() *StubDB {
	return &StubDB{}
}

func (s *StubDB) Read(data interface{}) error {
	file := []byte(prods)
	return json.Unmarshal(file, &data)
}
func (s *StubDB) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	stub := NewStub()
	r := NewRepository(stub)
	lista, err := r.GetAll()
	assert.Equal(t, 2, len(lista), "should be the same")
	assert.Nil(t, err, "should be nil")
}

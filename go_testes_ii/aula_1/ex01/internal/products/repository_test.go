package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/anolivei/go_meli_go/tree/main/go_testes_ii/aula_1/ex01/pkg/store"
)

func TestGetAll(t *testing.T) {
	myStubProduct := store.New(store.FileType, "./test.json")
	product := NewRepository(myStubProduct)
	expectedProduct := []Product(nil)

	result, _ := product.GetAll()

	assert.Equal(t, expectedProduct, result)
}

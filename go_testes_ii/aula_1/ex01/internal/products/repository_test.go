package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/anolivei/go_meli_go/tree/main/go_testes_ii/aula_1/ex01/pkg/store"
)

func TestGetAll(t *testing.T) {
	myStubProduct := store.New(store.FileType, "./test.json")
	product := NewRepository(myStubProduct)

	var expectedProduct []Product
	expectedProduct = append(
					expectedProduct,
					Product{1, "microondas", "eletrodomestico", 200, 525.25})

	result, _ := product.GetAll()

	assert.Equal(t, expectedProduct, result)
}

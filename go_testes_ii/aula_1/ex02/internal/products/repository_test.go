package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/anolivei/go_meli_go/tree/main/go_testes_ii/aula_1/ex02/pkg/store"
)

func TestGetAll(t *testing.T) {
	var expectedProduct []Product
	expectedProduct = append(expectedProduct,
					Product{1, "microondas", "eletrodomestico", 200, 525.25})
	myStubProduct := store.New(store.FileType, "./test.json")
	myStubProduct.Write(expectedProduct)

	product := NewRepository(myStubProduct)

	result, _ := product.GetAll()

	assert.Equal(t, expectedProduct, result)
}

func TestUpdateName(t *testing.T) {
	var prod []Product
	prod = append(prod,
				Product{1, "microondas", "eletrodomestico", 200, 525.25})
	beforeUpdate := store.New(store.FileType, "./test.json")
	beforeUpdate.Write(prod)

	productBefore := NewRepository(beforeUpdate)
	afterUpdate, err := productBefore.UpdateName(1, "torradeira")

	var productAfter []Product
	productAfter = append(
					productAfter,
					Product{1, "torradeira", "eletrodomestico", 200, 525.25})
	assert.Equal(t, productAfter[0], afterUpdate)
	assert.True(t, beforeUpdate.ReadWasCalled())
	assert.Nil(t, err)
}

func TestUpdateNameNotFound(t *testing.T) {
	var prod []Product
	prod = append(prod,
				Product{1, "microondas", "eletrodomestico", 200, 525.25})
	beforeUpdate := store.New(store.FileType, "./test.json")
	beforeUpdate.Write(prod)

	productBefore := NewRepository(beforeUpdate)
	afterUpdate, err := productBefore.UpdateName(2, "torradeira")

	var productAfter []Product
	productAfter = append(
					productAfter,
					Product{ID:0, Name:"", Typee:"", Count:0, Price:0})

	assert.Equal(t, productAfter[0], afterUpdate)
	assert.True(t, beforeUpdate.ReadWasCalled())
	assert.NotNil(t, err)
}

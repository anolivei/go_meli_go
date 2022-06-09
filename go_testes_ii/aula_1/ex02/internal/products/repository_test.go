package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/anolivei/go_meli_go/tree/main/go_testes_ii/aula_1/ex02/pkg/store"
)

func TestGetAll(t *testing.T) {
	myStubProduct := store.New(store.FileType, "./test_get_all.json")
	product := NewRepository(myStubProduct)

	var expectedProduct []Product
	expectedProduct = append(
					expectedProduct,
					Product{1, "microondas", "eletrodomestico", 200, 525.25})

	result, _ := product.GetAll()

	assert.Equal(t, expectedProduct, result)
}

func TestUpdateName(t *testing.T) {
	beforeUpdate := store.New(store.FileType, "./test_update_name.json")
	productBefore := NewRepository(beforeUpdate)
	afterUpdate, _ := productBefore.UpdateName(1, "torradeira")

	var expectedProduct []Product
	expectedProduct = append(
					expectedProduct,
					Product{1, "torradeira", "eletrodomestico", 200, 525.25})
	assert.Equal(t, expectedProduct[0], afterUpdate)
	assert.True(t, beforeUpdate.ReadWasCalled())
}

func TestUpdateNameNotFound(t *testing.T) {
	beforeUpdate := store.New(store.FileType, "./test_update_name.json")
	productBefore := NewRepository(beforeUpdate)
	afterUpdate, _ := productBefore.UpdateName(2, "torradeira")

	var expectedProduct []Product
	expectedProduct = append(
					expectedProduct,
					Product{ID:0, Name:"", Typee:"", Count:0, Price:0})
	assert.Equal(t, expectedProduct[0], afterUpdate)
	assert.True(t, beforeUpdate.ReadWasCalled())
}

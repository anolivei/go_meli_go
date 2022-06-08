package calc_test

import (
	"testing"

	"github.com/anolivei/go_meli_go/tree/main/go_testes_i/aula_2/ex01/pkg/calc"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := calc.Sum(5, 5)
	expectedResult := 10
	// Using testing
	/*if result != expectedResult {
		t.Errorf(
			"A função Sum() retornou o resultado %v, mas o esperado é %v",
			result,
			expectedResult)
	}*/
	// Using Testify
	assert.Equal(t, expectedResult, result, "devem ser iguais")
} 

func TestSub(t *testing.T) {
	result := calc.Sub(10, 5)
	expectedResult := 5
	// Using testing
	/*if result != expectedResult {
		t.Errorf(
			"A função Sub() retornou o resultado %v, mas o esperado é %v",
			result,
			expectedResult)
	}*/
	// Using Testify
	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestMultiply(t *testing.T) {
	result := calc.Multiply(5, 5)
	expectedResult := 25
	// Using testing
	/*if result != expectedResult {
		t.Errorf(
			"A função Multiply() retornou o resultado %v, mas o esperado é %v",
			result,
			expectedResult)
	}*/
	// Using Testify
	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestDivision(t *testing.T) {
	result := calc.Division(15, 5)
	expectedResult := 3
	// Using testing
	/*if result != expectedResult {
		t.Errorf(
			"A função Division() retornou o resultado %v, mas o esperado é %v",
			result,
			expectedResult)
	}*/
	// Using Testify
	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

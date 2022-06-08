package calc_test

import (
	"testing"
	"github.com/anolivei/go_meli_go/tree/main/go_testes_i/aula_2/ex01/pkg/calc"
)

func TestSum(t *testing.T) {
	result := calc.Sum(5, 5)
	expectedResult := 10
	if result != expectedResult {
		t.Errorf(
			"A função Sum() retornou o resultado %v, mas o esperado é %v",
			result,
			expectedResult)
	}
	
}

func TestSub(t *testing.T) {
	result := calc.Sub(10, 5)
	expectedResult := 5
	if result != expectedResult {
		t.Errorf(
			"A função Sub() retornou o resultado %v, mas o esperado é %v",
			result,
			expectedResult)
	}
	
}

func TestMultiply(t *testing.T) {
	result := calc.Multiply(5, 5)
	expectedResult := 25
	if result != expectedResult {
		t.Errorf(
			"A função Multiply() retornou o resultado %v, mas o esperado é %v",
			result,
			expectedResult)
	}
	
}

func TestDivision(t *testing.T) {
	result := calc.Division(15, 5)
	expectedResult := 3
	if result != expectedResult {
		t.Errorf(
			"A função Division() retornou o resultado %v, mas o esperado é %v",
			result,
			expectedResult)
	}
	
}

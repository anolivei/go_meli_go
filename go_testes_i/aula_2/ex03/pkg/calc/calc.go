package calc

import "fmt"

// Sub recebe dois números inteiros como parâmetro e retorna
// a soma dos dois
func Sum(a int, b int) int {
	return a + b
}

// Sub recebe dois números inteiros como parâmetro e retorna 
// a subtração dos dois
func Sub(a int, b int) int {
	return a - b
}

// Sub recebe dois números inteiros como parâmetro e retorna 
// a multiplicação dos dois
func Multiply(a int, b int) int {
	return a * b
}

// Sub recebe dois números inteiros como parâmetro e retorna 
// a divisão dos dois
func Division(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("o denominador não pode ser 0")
	}
	return a / b, nil
}

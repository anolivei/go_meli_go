package main

import "fmt"

func main() {
	var idade uint
	var empregado bool
	var tempo float32
	var salario float32

	idade, empregado, tempo, salario = 22, true, 2, 100000

	switch {
	case idade >= 22 && empregado == true && tempo >= 1 && salario >= 100000:
		fmt.Println("Empréstimo concedido sem juros")
	case idade >= 22 && empregado == true && tempo >= 1 && salario < 100000:
		fmt.Println("Empréstimo concedido com juros")
	default:
		fmt.Println("Empréstimo não concedido")
	}
}

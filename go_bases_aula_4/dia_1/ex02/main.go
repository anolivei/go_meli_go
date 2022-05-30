package main

import (
	"errors"
	"fmt"
)

func checkSalary(salario int) (string, error) {
	if salario < 15000 {
		return "", errors.New("erro: O salário digitado não está dentro do valor mínimo")
	} else {
		return "Necessário pagamento de imposto", nil
	}
}

func checkError(salario int) {
	fmt.Println("Salario:", salario)
	str, err := checkSalary(salario)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}

func main() {
	var salario int = 14000
	checkError(salario)
	salario = 15000
	checkError(salario)
	salario = 16000
	checkError(salario)
}

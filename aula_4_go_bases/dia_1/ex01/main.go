package main

import "fmt"

type myError struct {
	msg string
}

func (e myError) Error() (string) {
	return fmt.Sprintln(e.msg)
}

func checkSalary(salario int) (string, error) {
	if salario < 15000 {
		return "", &myError{"erro: O salário digitado não está dentro do valor mínimo"}
	} else {
		return "Necessário pagamento de imposto", nil
	}
}

func checkError(salario int) {
	fmt.Println("Salario:", salario)
	str, err := checkSalary(salario)
	if err != nil {
		fmt.Print(err)
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

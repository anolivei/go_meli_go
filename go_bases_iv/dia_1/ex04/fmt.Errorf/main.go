package main

import (
	"fmt"
)

type funcionario struct {
	horasTrabalhadas int
	valorHora float64
}

const (
	ERRO = "erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês"
)

func calcSalario(horasTrabalhadas int, valorHora float64)(float64, error){
	if horasTrabalhadas < 80 {
		return 0, fmt.Errorf("%s e ele trabalhou somente %d horas",
								ERRO, horasTrabalhadas)
	}
	var salario float64 = float64(horasTrabalhadas) * valorHora
	if salario >= 20000 {
		salario = salario * 0.9
	}
	return salario, nil
}

func checkError(f funcionario) {
	fmt.Printf("horas trabalhadas: %d valor por hora: %.2f\n",
				f.horasTrabalhadas,
				f.valorHora)
	str, err := calcSalario(f.horasTrabalhadas, f.valorHora)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("salario:", str)
}

func main() {
	func1 := funcionario{60, 200}
	checkError(func1)
	func2 := funcionario{80, 200}
	checkError(func2)
	func3 := funcionario{200, 500}
	checkError(func3)
}

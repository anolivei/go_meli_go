package main

import "fmt"

func main() {
	var sobrenome string = "Silva"
	var idade int = 25
	boolean :=  false
	var salario float32 = 4585.90
	var nome string = "Fellipe"

	fmt.Printf("sobrenome: %s idade: %d bool: %t salario: %.2f nome: %s\n",
				sobrenome, idade, boolean, salario, nome)
}
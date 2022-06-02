package main

import "fmt"

func main(){
	var nome string = "Joao"
	var sobrenome string = "das Neves"
	idade := 6
	var licencaParaDirigir bool = true
	var estaturaDaPessoa float32 = 1.75
	quantidadeDeFilhos := 2

	fmt.Printf("nome: %s sobrenome: %s idade: %d licença para dirigir: %t estatura: %.2f filhos: %d\n",
				nome, sobrenome, idade, licencaParaDirigir, estaturaDaPessoa, quantidadeDeFilhos)
}

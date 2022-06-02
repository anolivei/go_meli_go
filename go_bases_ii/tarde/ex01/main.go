package main

import "fmt"

type aluno struct {
	nome			string
	sobrenome		string
	rg				string
	dataAdmissao	string
}

func (a aluno) printDados(){
	fmt.Printf("Nome: %s %s RG: %s Data de admissão: %s\n",
		a.nome, a.sobrenome, a.rg, a.dataAdmissao)
}

func main(){
	a := aluno{"Giovanna", "Miranda Pereira", "12.345.678", "21/10/2022"}
	a.printDados()
}

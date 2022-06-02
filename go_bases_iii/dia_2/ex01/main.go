package main

import "fmt"

type usuario struct {
	nome string
	sobrenome string
	idade int
	email string
	senha string
}

func (u *usuario)atualizarNomeSobrenome(nome string, sobrenome string) {
	u.nome = nome
	u.sobrenome = sobrenome
}

func (u *usuario)atualizarIdade(idade int) {
	u.idade = idade
}

func (u *usuario)atualizarEmail(email string) {
	u.email = email
}

func (u *usuario)atualizarSenha(senha string) {
	u.senha = senha
}

func main() {
	u := usuario{"Giovanna", "Miranda Pereira", 5, "gmp@gmail.com", "123"}
	fmt.Printf("Nome: %s %s, Idade: %d E-mail: %s, Senha: %s\n", 
				u.nome,
				u.sobrenome,
				u.idade,
				u.email,
				u.senha)

	u.atualizarNomeSobrenome("Gio", "Miranda")
	u.atualizarIdade(10)
	u.atualizarEmail("giomiranda@gmail.com")
	u.atualizarSenha("abcd")
	fmt.Printf("Nome: %s %s, Idade: %d E-mail: %s, Senha: %s\n", 
				u.nome,
				u.sobrenome,
				u.idade,
				u.email,
				u.senha)
}

package main

import "fmt"

type usuario struct {
	nome string
	sobrenome string
	email string
	produtos []produto
}

type produto struct {
	nome string
	preco float64
	quantidade int
}

func novoProduto(nome string, preco float64) produto {
	p := produto{nome, preco, 0}
	return p
}

func adicionarProduto(u *usuario, p *produto, quantidade int) {
	p.quantidade = quantidade
	u.produtos = append(u.produtos, *p)
}

func deletarProduto(u *usuario) {
	u.produtos = make([]produto, 0)
}

func main() {
	u := usuario{"Joao", "das Neves", "jdasneves@gmail.com", nil}
	p1 := novoProduto("notebook", 5000)
	adicionarProduto(&u, &p1, 100)

	p2 := novoProduto("geladeira", 3000)
	adicionarProduto(&u, &p2, 100)

	fmt.Println(u.produtos)

	deletarProduto(&u)
	
	fmt.Println(u.produtos)
}

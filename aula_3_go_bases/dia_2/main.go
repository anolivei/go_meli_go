package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type produto struct {
	id int
	preco float64
	quantidade int
}

type CustoProduto interface{
	CalcularCustoProduto() float64
}

type Produtos struct{
	prod []produto
}

type Total interface{
	CalcularTotal() float64
}

func (p produto)CalcularCustoProduto()float64{
	return p.preco * float64(p.quantidade)
}

func (p Produtos)CalcularTotal()float64{
	var total float64 = 0
	for _, produto := range p.prod{
		total += produto.CalcularCustoProduto()
	}
	return total
}

func (p *Produtos)Adicionar(prod produto){
	p.prod = append(p.prod, prod)
}

func ReadFile(p *Produtos){
	fi := "./myFile.csv"
	f, err := os.Open(fi)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := r.ReadString('\n')
	for e == nil {
		s2 := strings.Split(s, ";")
		var prod produto
		prod.id, _ = strconv.Atoi(s2[0])
		prod.preco, _ = strconv.ParseFloat(s2[1], 64)
		prod.quantidade, _ = strconv.Atoi((strings.Split(s2[2], "\n"))[0])
		fmt.Printf("%10d %10.2f %10d\n", prod.id, prod.preco, prod.quantidade)
		p.Adicionar(prod)
		s, e = r.ReadString('\n')
	}
}

func main() {
	var p Produtos
	fmt.Printf("%11s %10s %17s\n", "ID", "Preço", "Quantidade")
	ReadFile(&p)
	fmt.Println("Total:", p.CalcularTotal())
}

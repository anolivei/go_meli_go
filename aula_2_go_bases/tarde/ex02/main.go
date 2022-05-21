package main

import "fmt"

const (
	pequeno = "pequeno"
	medio = "medio"
	grande = "grande"
)

const (
	custoAdcionalP = 1
	custoAdcionalM = 1.03
	custoAdcionalG = 1.06
	custoAdcionalEnvioG = 2500
)

type loja struct{
	prod	[]produto
}

type produto struct{
	tipo	string
	nome	string
	preco	float64
}

type Produto interface{
	CalcularCusto() float64
}

type Ecommerce interface{
	Total() float64
	Adicionar(p produto)
}

func (l loja)Total()float64{
	var total float64 = 0
	for _, produto := range l.prod{
		total += produto.CalcularCusto()
	}
	return total
}

func (l *loja)Adicionar(p produto){
	l.prod = append(l.prod, p)
}

func (p produto) CalcularCusto()(float64){
	switch p.tipo{
	case pequeno:
		return p.preco * custoAdcionalP
	case medio:
		return p.preco * custoAdcionalM
	case grande:
		return p.preco * custoAdcionalG + custoAdcionalEnvioG
	}
	return 0
}

func novoProduto(tipo string, nome string, preco float64)produto{
	var p produto = produto{tipo, nome, preco}
	return p
}

func novaLoja()loja{
	var e loja
	return e
}

func main(){
	notebook := novoProduto(pequeno, "notebook", 5000)
	fmt.Println("Preço notebook:", notebook.CalcularCusto())
	tv := novoProduto(medio, "TV", 3000)
	fmt.Println("Preço TV:", tv.CalcularCusto())
	geladeira := novoProduto(grande, "geladeira", 4500)
	fmt.Println("Preço geladeira:", geladeira.CalcularCusto())

	meli := novaLoja()
	meli.Adicionar(notebook)
	meli.Adicionar(tv)
	meli.Adicionar(geladeira)

	fmt.Println("Total meli:", meli.Total())
}
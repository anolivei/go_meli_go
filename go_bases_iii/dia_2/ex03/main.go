package main

import "fmt"

type produtos struct {
	nome string
	preco float64
	quantidade int
}

type servicos struct {
	nome string
	preco float64
	minTrabalhados int
}

type manutencao struct {
	nome string
	preco float64
}

func somarProdutos(p []produtos, out chan<- float64) {
	var total float64 = 0
	for _, prod := range p {
		total += (prod.preco * float64(prod.quantidade))
	}
	out <- total
}

func somarServicos(s []servicos, out chan<- float64) {
	var total float64 = 0
	for _, serv := range s {
		if serv.minTrabalhados < 30 {
			serv.minTrabalhados = 30
		}
		total += (serv.preco * float64(serv.minTrabalhados / 60))
	}
	out <- total
}

func somarManutencao(m []manutencao, out chan<- float64) {
	var total float64 = 0
	for _, man := range m {
		total += man.preco
	}
	out <- total
}

func main() {
	notebook := produtos{"notebook", 5000, 50}
	microondas := produtos{"microondas", 2000, 100}
	p := []produtos{notebook, microondas}

	instalacao := servicos{"instalacao", 60, 120}
	limpeza := servicos{"limpeza", 55, 180}
	s := []servicos{instalacao, limpeza}

	atualizacao := manutencao{"atualizacao", 101}
	backup := manutencao{"backup", 302}
	m := []manutencao{atualizacao, backup}

	totalChP := make(chan float64)
	totalChS := make(chan float64)
	totalChM := make(chan float64)

	go somarProdutos(p, totalChP)
	p2 := <- totalChP

	go somarServicos(s, totalChS)
	s2 := <- totalChS

	go somarManutencao(m, totalChM)
	m2 := <- totalChM

	fmt.Println(p2 + s2 + m2)
}
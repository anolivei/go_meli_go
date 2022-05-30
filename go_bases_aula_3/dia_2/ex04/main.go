package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	insercao = "insercao"
	grupo = "grupo"
	selecao = "selecao"
)

func orderByIsercao(items []int, out chan<- time.Duration) {
	t1 := time.Now()
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j - 1] > items[j] {
				items[j - 1], items[j] = items[j], items[j - 1]
			}
			j = j - 1
		}
	}
	t2 := time.Now()
	fmt.Println("insercao", len(items), t2.Sub(t1))
	out <- t2.Sub(t1)
}

func orderByGrupo(items []int, out chan<- time.Duration) {
	t1 := time.Now()
	sort.Ints(items)
	t2 := time.Now()
	fmt.Println("grupo", len(items), t2.Sub(t1))
	out <- t2.Sub(t1)
}

func orderBySelecao(items []int, out chan<- time.Duration) {
	t1 := time.Now()
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
	t2 := time.Now()
	fmt.Println("selecao", len(items), t2.Sub(t1))
	out <- t2.Sub(t1)
}

func order(numCem []int, numMil []int, numDezMil []int, tipo string, out chan<- time.Duration) {
	switch tipo {
	case insercao:
		defer
		orderByIsercao(numCem, out)
		defer
		orderByIsercao(numMil, out)
		defer
		orderByIsercao(numDezMil, out)
	case grupo:
		defer
		orderByGrupo(numCem, out)
		defer
		orderByGrupo(numMil, out)
		defer
		orderByGrupo(numDezMil, out)
	case selecao:
		defer
		orderBySelecao(numCem, out)
		defer
		orderBySelecao(numMil, out)
		defer
		orderBySelecao(numDezMil, out)
	}
}

func main() {
	numCem := rand.Perm(100)
	numMil := rand.Perm(1000)
	numDezMil := rand.Perm(10000)
	insercaoCh := make(chan time.Duration)

	go order(numCem, numMil, numDezMil, insercao, insercaoCh)
	<- insercaoCh

	numCem = rand.Perm(100)
	numMil = rand.Perm(1000)
	numDezMil = rand.Perm(10000)
	grupoCh := make(chan time.Duration)

	go order(numCem, numMil, numDezMil, grupo, grupoCh)
	<- grupoCh

	numCem = rand.Perm(100)
	numMil = rand.Perm(1000)
	numDezMil = rand.Perm(10000)
	selecaoCh := make(chan time.Duration)

	go order(numCem, numMil, numDezMil, selecao, selecaoCh)
	<- selecaoCh

	a := <- insercaoCh
	b := <- grupoCh
	c := <- selecaoCh

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

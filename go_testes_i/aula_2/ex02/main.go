package main

import (
	"fmt"
	"github.com/anolivei/go_meli_go/tree/main/go_testes_i/aula_2/ex02/pkg/ordering"
)

func main() {
	num := []int{5, 1, 7, 4, 3, 0}
	s := ordering.SortNums(num)
	fmt.Println(s)
}

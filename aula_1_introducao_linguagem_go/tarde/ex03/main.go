package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		i := os.Args[1]
		j, _ := strconv.Atoi(i)
		if j > 12 || j < 1 {
			fmt.Println("Digite um número entre 1 e 12")
			return
		}
		var meses = map[int]string{
			1: "janeiro",
			2: "fevereiro",
			3: "março",
			4: "abril",
			5: "maio",
			6: "junho",
			7: "julho",
			8: "agosto",
			9: "setembro",
			10: "outubro",
			11: "novembro",
			12: "dezembro"}
			fmt.Println(meses[j])
	} else {
		fmt.Println("Digite um número entre 1 e 12")
	}
}

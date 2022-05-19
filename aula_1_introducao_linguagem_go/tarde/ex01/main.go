package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		word := os.Args[1]
		fmt.Println("o tamanho da palavra é:", len(word))
		for i := 0; i < len(word); i++ {
			fmt.Println(string(word[i]))
		}
	}
	if len(os.Args) != 2 {
		fmt.Println("Digite uma palavra =)")
	}
}

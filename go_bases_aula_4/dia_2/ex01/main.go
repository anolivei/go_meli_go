package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Iniciando...")
	_, err := os.Open("customers.txt")
	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
	fmt.Println("Fim")
}

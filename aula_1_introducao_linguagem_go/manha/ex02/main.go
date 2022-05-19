package main

import "fmt"

func main() {
	var temperatura int8 = 11
	var umidade uint8 = 67
	var pressaoAtmosferica float32 = 1.017

	fmt.Printf("Temperatura: %d Umidade: %d Pressão Atmosférica: %.3f\n",
				temperatura, umidade, pressaoAtmosferica)
}
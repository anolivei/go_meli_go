package main

import "fmt"

func calcTax(salario float32, tax float32) float32{
	return salario * (tax / 100)
}

func main(){
	fmt.Printf("A taxa é %.2f reais\n", calcTax(50000, 17))
	fmt.Printf("A taxa é %.2f dólares\n", calcTax(150000, 10))
}

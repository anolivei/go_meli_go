package main

import "fmt"

/*
** const for hours valor
*/
const (
	A = 3000
	B = 1500
	C = 1000
)

/*
** const for extra hours valor
*/
const (
	AEH = 1.5 * A
	BEH = 1.2 * B
	CEH = 1 * C
)

func calc(hours int, valor int, extraValor float32)float32{
	extraHours := 0
	if hours > 160{
		extraHours = hours - 160
		hours = 160
	}
	return float32(hours * valor) + (float32(extraHours) * extraValor)
}

func calcSalary(category string, hours int)float32{
	var salary float32
	switch category{
	case "A":
		salary = calc(hours, A, AEH)
	case "B":
		salary = calc(hours, B, BEH)
	case "C":
		salary = calc(hours, C, CEH)
	}
	return salary
}

func main(){
	fmt.Printf("Funcionário categoria C: %.2f\n", calcSalary("C", 162))
	fmt.Printf("Funcionário categoria B: %.2f\n", calcSalary("B", 176))
	fmt.Printf("Funcionário categoria A: %.2f\n", calcSalary("A", 172))
}
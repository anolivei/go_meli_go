package main

import "fmt"

func main() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Manuel": 26,
		"Brenda": 19,
		"Dario": 44,
		"Pedro": 30}
	fmt.Println("A idade de Benjamin é ", employees["Benjamin"])

	n := 0
	for _, value := range employees{
		if value >= 21{
			n += 1
		}
	}
	fmt.Println("O número de funcionários com mais de 21 anos é", n)

	fmt.Println(employees)
	employees["Frederico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)
}

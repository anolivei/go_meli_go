package main

import "fmt"

func delEmployee(employees map[string]int, employee string){
	delete(employees, employee)
}

func addEmployee(employees map[string]int, employee string, age int){
	employees[employee] = age
}

func countEmployessByAge(employees map[string]int, age int){
	n := 0
	for _, value := range employees{
		if value >= age{
			n += 1
		}
	}
	fmt.Printf("O número de funcionários com mais de %d anos é %d\n", age, n)
}

func findEmployeeAge(employees map[string]int, employee string){
	fmt.Printf("A idade de %s é %d anos\n", employee, employees[employee])
}

func main() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Manuel": 26,
		"Brenda": 19,
		"Dario": 44,
		"Pedro": 30}

	findEmployeeAge(employees, "Benjamin")

	countEmployessByAge(employees, 21)

	fmt.Println(employees)
	addEmployee(employees, "Frederico", 25)
	fmt.Println(employees)
	delEmployee(employees, "Pedro")
	fmt.Println(employees)
}

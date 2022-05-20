package main

import (
	"errors"
	"fmt"
)

const (
	tarantula = "tarantula"
	hamster = "hamster"
	cachorro = "cachorro"
	gato = "gato"
)

const (
	amountTarantula = 0.15
	amountHamster = 0.25
	amountCachorro = 10.00
	amountGato = 5.00
)

func animalDefault(num float64) float64{
	return 0
}

func animalTarantula(num float64) float64{
	return num * amountTarantula
}

func animalHamster(num float64) float64{
	return num * amountHamster
}

func animalCachorro(num float64) float64{
	return num * amountCachorro
}

func animalGato(num float64) float64{
	return num * amountGato
}

func Animal(animal string)(func(num float64)(float64), error){
	switch animal {
	case tarantula:
		return animalTarantula, nil
	case hamster:
		return animalHamster, nil
	case cachorro:
		return animalCachorro, nil
	case gato:
		return animalGato, nil
	}
	return animalDefault, errors.New("animal não definido")
}

func main(){
	animalTarantula, msg := Animal(tarantula)
	animalHamster, msg := Animal((hamster))
	animalDog, msg := Animal(cachorro)
	animalCat, msg := Animal(gato)
	animalDefault, msg := Animal("default")

	var amount float64
	amount += animalTarantula(10)
	amount += animalHamster(15)
	amount += animalDog(5)
	amount += animalCat(8)
	amount += animalDefault(20)

	fmt.Println(amount, msg)
}
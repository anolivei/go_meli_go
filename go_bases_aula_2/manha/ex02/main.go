package main

import (
	"errors"
	"fmt"
)

func calcMean(grades ...float32)(float32, error){
	var sum float32 = 0
	i := 0
	for _, grade := range grades {
		if grade < 0 {
			return 0, errors.New("notas netagtivas não são aceitas")
		}
		sum += grade
		i++
	}
	return (sum / float32(i)), nil
}

func main(){
	mean, e := calcMean(10, -10, 9.5, 7)
	if e != nil{
		fmt.Println("A média não pode ser calculada porque", e)
	} else {
		fmt.Println("A média é:", mean)
	}
}

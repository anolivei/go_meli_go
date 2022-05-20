package main

import (
	"errors"
	"fmt"
)

const (
	mean = "mean"
	minimum = "minimum"
	maximum = "maximum"
)

func meanFunc(grades ...float32)(float32){
	var sum float32 = 0
	i := 0
	for _, grade := range grades {
		sum += grade
		i++
	}
	return (sum / float32(i))
}

func minFunc(grades ...float32)(float32){
	min := grades[0]
	for _, grade := range grades {
		if min > grade{
			min = grade
		}
	}
	return min
}

func maxFunc(grades ...float32)(float32){
	max := grades[0]
	for _, grade := range grades {
		if max < grade{
			max = grade
		}
	}
	return max
}

func defaultFunc(grades ...float32)(float32){
	return 0
}

func operation (operator string) (func(grades ...float32)(float32), error){
	switch operator {
	case mean:
		return meanFunc, nil
	case minimum:
		return minFunc, nil
	case maximum:
		return maxFunc, nil
	}
	return defaultFunc, errors.New("operador não informado")
}

func main(){
	meanFunc, err := operation(mean)
	minFunc, err := operation(minimum)
	maxFunc, err := operation(maximum)
	defaultFunc, err := operation("default")

	fmt.Println(err)

	meanValue := meanFunc(2, 3, 3, 4, 1, 2, 4, 5)
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	defaultValue := defaultFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Mean: %.2f Min: %.2f Max: %.2f default: %.2f\n",
		meanValue, minValue, maxValue, defaultValue)
}
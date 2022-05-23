package main

import (
	"fmt"
	"os"
	"strconv"
)

type produto struct {
	id int
	preco float64
	quantidade int
}

func ftWriteString(prod produto){
	text := fmt.Sprintf("%d;%.2f;%d\n", prod.id, prod.preco, prod.quantidade)
	f, err := os.OpenFile("./myFile.csv",
						os.O_APPEND|os.O_WRONLY|os.O_CREATE,
						0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) == 4 {
		id, _ := strconv.Atoi(os.Args[1])
		preco, _ := strconv.ParseFloat(os.Args[2], 64)
		quantidade, _ := strconv.Atoi(os.Args[3])
		prod := produto{id, preco, quantidade}
		ftWriteString(prod)
	} else {
		fmt.Println("São necessários três argumentos:")
		fmt.Println("1 - id do produto")
		fmt.Println("2 - preço do produto")
		fmt.Println("3 - quantidade do produto")
	}
}

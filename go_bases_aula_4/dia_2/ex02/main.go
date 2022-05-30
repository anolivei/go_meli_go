package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type cliente struct {
	arquivo int
	nomeCompleto string
	rg string
	telefone string
	endereco string
}

func gerarIdArquivo(c *cliente) {
	low := 1000000000
	hi := 9999999999
	rand.Seed(time.Now().UnixNano())
	c.arquivo =  low + rand.Intn(hi - low)
}

func lerArquivo() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()
	file, err := os.Open("customers.txt")
	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}
	defer file.Close()
}

func cadastrarCliente(c *cliente, nome string, rg string,
						telefone string, endereco string) {
	if nome == "" || rg == "" || telefone == "" || endereco == "" {
		panic("nenhum dos dados do cliente pode ser nulo")
	}
	c.nomeCompleto = nome
	c.rg = rg
	c.telefone = telefone
	c.endereco = endereco
}

func main() {
	var c cliente
	gerarIdArquivo(&c)
	if c.arquivo == 0 {
		panic("o id está zerado")
	}
	lerArquivo()
	cadastrarCliente(&c,
					"João das Neves",
					"123.456-07",
					"(11)91234-5678",
					"Rua dos Bobos, n 0")
	fmt.Println(c)
}

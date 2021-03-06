<p align="center">
<img src="https://raw.githubusercontent.com/egonelbre/gophers/10cc13c5e29555ec23f689dc985c157a8d4692ab/vector/fairy-tale/witch-learning.svg" width="250"/>
</p>

# Bootcamp de Go do Mercado Livre

## Go Bases
| módulo | aula | exercício | assunto |
|--------|------|-----------|---------|
| i | manhã | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_bases_i/manha/ex01) | Imprimindo nome e idade na tela |
| i | manhã | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_bases_i/manha/ex02) | Clima - imprimindo três tipos de variávies na tela |
| i | manhã | [ex03](https://github.com/anolivei/go_meli_go/tree/main/go_bases_i/manha/ex03) | Declaração de variáveis - fazer correções |
| i | manhã | [ex04](https://github.com/anolivei/go_meli_go/tree/main/go_bases_i/manha/ex04) | Tipos de dados - fazer correções |
| i | tarde | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_bases_i/tarde/ex01) | Letras de uma palavra - imprimir quantidade de letras e letras de uma palavra (`len, for e Args`) |
| i | tarde | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_bases_i/tarde/ex02) | Empréstimo - `switch case` para devolver mensagens |
| i | tarde | [ex03](https://github.com/anolivei/go_meli_go/tree/main/go_bases_i/tarde/ex03) | A que mês corresponde? - `map` para meses do ano |
| i | tarde | [ex04](https://github.com/anolivei/go_meli_go/tree/main/go_bases_i/tarde/ex04) | Quantos anos tem - `map` para idades de funcionarios (`for range delete`)|
| ii | manhã | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_bases_ii/manha/ex01) | Impostos de salário (`func`) |
| ii | manhã | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_bases_ii/manha/ex02) | Calcular média - função com número dinâmico de parâmetros: Elipsis (`...`) |
| ii | manhã | [ex03](https://github.com/anolivei/go_meli_go/tree/main/go_bases_ii/manha/ex03) | Calcular salário (`switch case const`) |
| ii | manhã | [ex04](https://github.com/anolivei/go_meli_go/tree/main/go_bases_ii/manha/ex04) | Cálculo de estatísticas - multi retorno |
| ii | manhã | [ex05](https://github.com/anolivei/go_meli_go/tree/main/go_bases_ii/manha/ex05) | Cálculo da quantidade de alimento - retorno de função |
| ii | tarde | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_bases_ii/tarde/ex01) | Registro de estudantes - structs |
| ii | tarde | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_bases_ii/tarde/ex02) | Produtos de e-commerce - structs e interfaces (`append`)|
| iii | dia 01 |[ex01](https://github.com/anolivei/go_meli_go/tree/main/go_bases_iii/dia_1/ex01) | Guardar arquivo (`fmt.Sprintf os.OpenFile os.O_APPEND os.O_WRONLY os.O_CREATE`)|
| iii | dia 01 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_bases_iii/dia_1/ex02) | Ler arquivo (`bufio.NewReader(file) ReadString strings.Split(str1, str2) os.Open(file)`) |
| iii | dia 02 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_bases_iii/dia_2/ex01) | Rede social -  ponteiros |
| iii | dia 02 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_bases_iii/dia_2/ex02) | E-commerce - ponteiros (`make`)|
| iii | dia 02 | [ex03](https://github.com/anolivei/go_meli_go/tree/main/go_bases_iii/dia_2/ex03) | Calcular Preço - chan e go routines (`make(chan type) go func`) |
| iii | dia 02 | [ex04](https://github.com/anolivei/go_meli_go/tree/main/go_bases_iii/dia_2/ex04) | Ordenamento - chan, go routines e defer |
| iv | dia 01 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_bases_iv/dia_1/ex01)| Imposto sobre o salário - tratamento de erros com `Error()` |
| iv | dia 01 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_bases_iv/dia_1/ex02) | Imposto sobre o salário - tratamento de erros com `errors.New()` |
| iv | dia 01 | [ex03](https://github.com/anolivei/go_meli_go/tree/main/go_bases_iv/dia_1/ex03) | Imposto sobre o salário - tratamento de erros com `fmt.Errorf()`|
| iv | dia 01 | [ex04](https://github.com/anolivei/go_meli_go/tree/main/go_bases_iv/dia_1/ex04) | Imposto sobre o salário - tratamento de erros com `errors.New(), fmt.Errorf() e errors.Unwrap()` |
| iv | dia 02 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_bases_iv/dia_2/ex01) | Dados dos clientes - `panic` |
| iv | dia 02 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_bases_iv/dia_2/ex02) | Registrando clientes - `panic, defer e recover` |

## Go Web
| módulo | aula | exercício | assunto |
|--------|------|-----------|---------|
| i | aula 01 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_web_i/aula_1/ex01) | Estruturar um JSON usando [Marshal](https://pkg.go.dev/encoding/json#Marshal) do [json](https://pkg.go.dev/encoding/json) |
| i | aula 01 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_web_i/aula_1/ex02) | "Hello {nome}" método GET na API utilizando [gin](https://github.com/gin-gonic/gin) |
| i | aula 01 | [ex03](https://github.com/anolivei/go_meli_go/tree/main/go_web_i/aula_1/ex03) | Listar Entidade (produtos/usuários/transações) usando [Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) do json e criando o handler GetAll |
| i | aula 02 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_web_i/aula_2/ex01) | Filter usando [Query](https://github.com/gin-gonic/gin#another-example-query--post-form) do gin |
| i | aula 02 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_web_i/aula_2/ex02) | Endpoint /theme/:id usando [Param](https://github.com/gin-gonic/gin#parameters-in-path) do gin |
| ii | aula 01 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_web_ii/aula_1/ex01) | Criação de entidade (struct de produtos/usuários/transações), métodos GET e POST e id automático |
| ii | aula 01 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_web_ii/aula_1/ex02) | Validação de campo utilizando a lib [validator](https://github.com/go-playground/validator) |
| ii | aula 01 | [ex03](https://github.com/anolivei/go_meli_go/tree/main/go_web_ii/aula_1/ex03) | Validação de Token |
| ii | aula 02 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_web_ii/aula_2/ex01) | Arquitetura para APIs em Go - internal (service.go e repository.go) |
| ii | aula 02 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_web_ii/aula_2/ex02) | Arquitetura para APIs em Go - server (main.go e handler/controller) |
| iii | aula 01 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_web_iii/aula_1/ex01) | Gerar o método PUT |
| iii | aula 01 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_web_iii/aula_1/ex02) | Gerar o método DELETE |
| iii | aula 01 | [ex03](https://github.com/anolivei/go_meli_go/tree/main/go_web_iii/aula_1/ex03) | Gerar o método PATCH |
| iii | aula 02 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_web_iii/aula_2/ex01) | configuração do .env (variáveis de ambiente na API) utilizando a lib [godotenv](https://github.com/joho/godotenv) |
| iii | aula 02 | [ex02 e ex03](https://github.com/anolivei/go_meli_go/tree/main/go_web_iii/aula_2/ex02_and_ex03) | guardar e ler informações de um arquivo .json (banco de dados em disco)  |
| iv | aula 01 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_web_iv/aula_1/ex01) | respostas genéricas (web package com response.go) para devolver respostas em json com status code e erro/json de retorno |
| iv | aula 02 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_web_iv/aula_2/ex01) | Documentação da API utilizando a lib [swagger](https://github.com/swaggo/gin-swagger) |

## Go Testes
| módulo | aula | exercício | assunto |
|--------|------|-----------|---------|
| i | aula 01 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_testes_i/aula_1/ex01) | Caixa branca e a caixa preta |
| i | aula 01 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_testes_i/aula_1/ex02) | Teste funcional |
| i | aula 01 | [ex03](https://github.com/anolivei/go_meli_go/tree/main/go_testes_i/aula_1/ex03) | Teste de integração |
| i | aula 01 | [ex04](https://github.com/anolivei/go_meli_go/tree/main/go_testes_i/aula_1/ex04) | Dimensões prioritárias de qualidade no MELI |
| i | aula 02 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_testes_i/aula_2/ex01) | Teste Unitário - Função Sub usando o package [testing](https://pkg.go.dev/testing) e [testify](https://github.com/stretchr/testify) |
| i | aula 02 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_testes_i/aula_2/ex02) | Teste Unitário - Função de Ordenar |
| i | aula 02 | [ex03](https://github.com/anolivei/go_meli_go/tree/main/go_testes_i/aula_2/ex03) | Teste Unitário - Função de Dividir |
| ii | aula 01 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_testes_ii/aula_1/ex01) | Testes Unitário na API - GetAll() |
| ii | aula 01 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_testes_ii/aula_1/ex02) | Testes Unitário na API - UpdateName() |
| ii | aula 02 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_testes_ii/aula_2/ex01) | Testes de integração na API - Service/Repo/Db Update() |
| ii | aula 02 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_testes_ii/aula_2/ex02) | Testes de integração na API - Service/Repo/Db Delete() |
| iii | aula 01 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_testes_iii/aula_1/ex01) | Melhorar o código com [Golangci-lint](https://golangci-lint.run/) |
| iii | aula 01 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_testes_iii/aula_1/ex02) | Medir a cobertura (`go test -v --cover ./...`)
| iii | aula 01 | [ex03](https://github.com/anolivei/go_meli_go/tree/main/go_testes_iii/aula_1/ex03) | Aumentar a cobertura |
| iii | aula 02 | [ex01](https://github.com/anolivei/go_meli_go/tree/main/go_testes_iii/aula_2/ex01) | Functional Testing Update() |
| iii | aula 02 | [ex02](https://github.com/anolivei/go_meli_go/tree/main/go_testes_iii/aula_2/ex02) | Functional Testing Delete() |
| iii | aula 02 | [ex03](https://github.com/anolivei/go_meli_go/tree/main/go_testes_iii/aula_2/ex03) | Realizar TDD |

## Go Database
| módulo | aula | assunto |
|--------|------|---------|
| i | [aula 01](https://github.com/anolivei/go_meli_go/tree/main/go_db_i/aula_1) | Banco de dados SQL - CAP (Consistência, Disponibilidade e Tolerância de particionamento) e DER (Diagrama de Relacionamento de Entidade) |
| i | [aula 02](https://github.com/anolivei/go_meli_go/tree/main/go_db_i/aula_2) | Banco de dados SQL - `SELECT, WHERE, ORDER BY, ASC, DESC, AND, OR, LIKE, NOT LIKE, BETWEEN, LIMIT, OFFSET, DISTINCT, COUNT, MIN, MAX, SUM, AVG` |
| ii | [aula 01](https://github.com/anolivei/go_meli_go/tree/main/go_db_ii/aula_1) | Banco de dados SQL - `INNER JOIN, LEFT JOIN, RIGHT JOIN, GROUP BY, HAVING` e subconsultas|
| ii | [aula 02](https://github.com/anolivei/go_meli_go/tree/main/go_db_ii/aula_2) | Banco de dados SQL - `INSERT UPDATE DELETE TEMPORARY TABLE, INDEX` |
| iii | [aula 01](https://github.com/anolivei/go_meli_go/tree/main/go_db_iii/aula_1) | Banco de dados NoSQL - mongodb |
| iii | [aula 02](https://github.com/anolivei/go_meli_go/tree/main/go_db_iii/aula_2) | Banco de dados NoSQL - mongodb |

## [Go Implementação db](https://github.com/anolivei/go_meli_go/tree/main/go_implementacao_db)

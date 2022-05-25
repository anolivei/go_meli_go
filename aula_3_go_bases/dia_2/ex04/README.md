# Exercício 4 - Ordenamento

Uma empresa de sistemas precisa analisar que algoritmos de ordenamento
utilizar para seus serviços.

Para eles é necessário instanciar 3 arrays com valores aleatórios não ordenados
- uma matriz de inteiros com 100 valores
- uma matriz de inteiros com 1000 valores
- uma matriz de inteiros com 10.000 valores

Para instanciar as variáveis, utilize o rand:

```go
package main

import (
"math/rand"
)

func main() {
variavel := rand.Perm(100)
variave2 := rand.Perm(1000)
variave3 := rand.Perm(10000)

}
````

Cada um deve ser ordenado por:
- [Inserção] (https://www.golangprograms.com/golang-program-for-implementation-of-insertionsort.html)
- [Grupo]
- [Seleção](https://www.golangprograms.com/golang-program-for-implementation-of-selection-sort.html)

Uma rotina para cada execução de classificação.

Tenho que esperar terminar a ordenação de 100 números para continuar
com a ordenação de 1000 e depois a ordenação de 10000.

Por fim, devo medir o tempo de cada um e mostrar o resultado na tela,
para saber qual ordenação ficou melhor para cada arranjo.
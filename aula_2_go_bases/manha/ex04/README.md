Os professores de uma universidade na Colômbia precisam calcular algumas
estatísticas de notas dos alunos de um curso, sendo necessário calcular
os valores mínimo, máximo e médio de suas notas.

Será necessário criar uma função que indique que tipo de cálculo deve ser
realizado (mínimo, máximo ou média) e que retorna outra função (e uma mensagem
caso o cálculo não esteja definido) que pode ser passado uma quantidade N de
inteiros e retorne o cálculo que foi indicado na função anterior.

Exemplo:

```go
const (
minimum = "minimum"
average = "average"
maximum = "maximum"
)

...
minhaFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

...
minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
```
# Exercício 4 - Imposto sobre o salário #4

Vamos fazer com que nosso programa seja um pouco mais complexo e útil.

1. Desenvolva as funções necessárias para permitir que a empresa calcule:

Salário mensal de um funcionário segundo a quantidade de horas trabalhadas.

- A função receberá as horas trabalhadas no mês e o valor da hora como
parâmetro.
- Esta função deve retornar mais de um valor (salário calculado e erro).
- No caso de o salário mensal ser igual ou superior a R$ 20.000, 10% devem ser
descontados como imposto.
- Se o número de horas mensais inseridas for menor que 80 ou um número
negativo, a função deverá retornar um erro. Deve indicar "erro: o trabalhador
não pode ter trabalhado menos de 80 horas por mês".

2. Desenvolva o código necessário para cumprir as funcionalidades solicitadas,
usando “errors.New()”, “fmt.Errorf()” e “errors.Unwrap()”. Não esqueça de
realizar as validações dos retornos de erro em sua função “main()”.
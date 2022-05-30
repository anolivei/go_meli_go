# Exercício 1 - Criar Entidade

A funcionalidade para criar a entidade deve ser implementada. Se isso
acontecer, os seguintes passos devem ser seguidos:

1. Crie um endpoint por meio de POST que receba a entidade.

2. Você deve ter um array da entidade na memória (no nível global), no qual
todas as requisições que são feitas devem ser salvas.

3. No momento de fazer a solicitação, o ID deve ser gerado. Para gerar o ID,
devemos procurar o ID do último registro gerado, incrementá-lo em 1 e
atribuí-lo ao nosso novo registro (sem ter uma variável global do último ID).

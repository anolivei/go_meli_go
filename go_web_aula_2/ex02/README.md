# Exercício 2 - Get one endpoint

Gere um novo endpoint que nos permita buscar um único resultado do
array de temas.

Usando parâmetros de caminho o endpoint deve ser /theme/:id (lembre-se que
o tema sempre tem que ser plural). Uma vez que o id é recebido, ele retorna
a posição correspondente.

1. Gere uma nova rota.

2. Gera um manipulador para a rota criada.

3. Dentro do manipulador, procure o item que você precisa.

4. Retorna o item de acordo com o id.

Se você não encontrou nenhum elemento com esse id retorne como código de
resposta 404.
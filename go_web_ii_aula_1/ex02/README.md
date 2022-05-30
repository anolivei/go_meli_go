# Exercício 2 - Validação de campo

As validações dos campos devem ser implementadas no momento do envio do pedido,
para isso devem ser seguidos os seguintes passos:

1. Todos os campos enviados na solicitação devem ser validados, todos os
campos são obrigatórios

2. Caso algum campo não esteja completo, um código de erro 400 deve ser retornado
com a mensagem “campo %s é obrigatório”.
(Em %s deve ir o nome do campo que não está completo).

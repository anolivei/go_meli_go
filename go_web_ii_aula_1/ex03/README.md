# Exercício 3 - Validar Token

Para adicionar segurança à aplicação, o pedido deve ser enviado com um token,
para isso devem ser seguidos os seguintes passos:

1. No momento do envio da solicitação, deve ser validado que um token é enviado

2. Esse token deve ser validado em nosso código (o token pode ser codificado
permanentemente).

3. Caso o token enviado não esteja correto, devemos retornar um erro 401 e uma
mensagem que "você não tem permissão para fazer a solicitação solicitada".

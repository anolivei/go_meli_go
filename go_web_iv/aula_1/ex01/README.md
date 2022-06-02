# Exercício 1 - Manipulando respostas genéricas

É necessário implementar um gerenciamento de resposta genérico para enviar
sempre o mesmo formato nas solicitações. Para conseguir isso, as seguintes
etapas devem ser executadas:

1. Construa o webpack dentro do diretório pkg.

2. Faça a estrutura Response com os capos: código, dados e erro:

a. code terá o código de retorno.

b. os dados terão a estrutura enviada pela aplicação (caso não haja erro).

c. error terá o erro recebido em formato de texto (caso haja erro).

3. Desenvolva uma função que receba o código como um inteiro, os dados como uma
interface e o erro como uma string.

4. A função deve retornar com base no código, se for uma resposta com os dados ou
com o erro.

5. Implemente esta função em todos os retornos dos controllers, antes de enviar a
resposta ao cliente, a função deve gerar a estrutura que definimos.

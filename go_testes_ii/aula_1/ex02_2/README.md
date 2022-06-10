# Exercício 2 - Teste Unitário - UpdateName()

Teste de Design UpdateName, onde é validado que a resposta retornada está
correta para a atualização do nome de um produto específico.

E verifica-se também que o método “Read” do Storage é realmente utilizado para
pesquisar o produto.

Para isto:
1. Crie um mock de Storage, este mock deve conter em seus dados um produto
específico cujo nome pode ser “Before Update”.

2. O método Read do Mock deve conter uma lógica que permita verificar se o
referido método foi invocado. Pode ser através de um booleano como visto
na classe.

3. Para dar o teste como OK, deve-se validar que ao invocar o método Repository
UpdateName, com o id do produto mockado e com o novo nome "After Update", ele
realmente faz a atualização. Também deve ser validado que o método Read foi
executado durante o teste.

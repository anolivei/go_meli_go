# Exercício 1 - Service/Repo/Db Update()

Projete um teste que teste na camada de serviço, o método ou função Update().

Para conseguir isso você deve:
1. Crie um mock de Storage, o mock deve conter em seus dados um produto com as
especificações desejadas.

2. O método Read do Mock deve conter uma lógica que permita verificar se o
referido método foi invocado.

3. Para dar o teste como OK, deve-se validar que ao invocar o método
Service Update(), retorne o produto com o mesmo ID e os dados atualizados.
Verifique também se Read() do Repositório foi executado durante o teste.
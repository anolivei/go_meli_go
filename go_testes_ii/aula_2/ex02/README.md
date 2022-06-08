# Exercício 2 - Service/Repo/Db Delete()

Projete um teste que teste na camada de serviço, o método ou função Delete().
A remoção correta de um produto deve ser testada, e o erro quando o produto não
existe.

Para conseguir isso você pode:

1. Crie um mock de Storage, o mock deve conter em seus dados um produto
com as especificações desejadas.

2. Execute o teste com dois IDs de produto diferentes, sendo um deles um ID
que não existe no Storage Mock.

3. Para dar o teste como OK, deve-se validar que o produto excluído não existe
mais no Storage após Delete().
Além disso, quando você tentar excluir um produto inexistente, deve receber o
erro correspondente.
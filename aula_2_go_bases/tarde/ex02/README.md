# Exercício 2 - Produtos de e-commerce

Diversas lojas de e-commerce precisam realizar funcionalidades no Go para
gerenciar produtos e devolver o valor do preço total.

As empresas têm 3 tipos de produtos:
- Pequeno, Médio e Grande.

Existem custos adicionais para manter o produto no armazém da loja e
custos de envio.

Lista de custos adicionais:
- Pequeno: O custo do produto (sem custo adicional)
- Médio: O custo do produto + 3% pela disponibilidade no estoque
- Grande: O custo do produto + 6% pela disponibilidade no estoque + um custo
adicional pelo envio de $2500.

## Requisitos:
- Criar uma estrutura “loja” que guarde uma lista de produtos.
- Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
- Criar uma interface “Produto” que possua o método “CalcularCusto”
- Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.
- Será necessário uma função “novoProduto” que receba o tipo de produto, seu
nome e preço, e devolva um Produto.
- Será necessário uma função “novaLoja” que retorne um Ecommerce.

### Interface Produto:
- Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o custo adicional segundo o tipo de produto.

### Interface Ecommerce:
- Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com
base no custo total dos produtos + o adicional citado anteriormente (caso a
categoria tenha).
- Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto
e adicioná-lo a lista da loja.
//a. O arquivo repository.go deve ser criado
package products

import "fmt"

//b. A estrutura da entidade deve ser criada
type Product struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
	Typee	string	`json:"type"`
	Count	int		`json:"count"`
	Price	float64	`json:"price"`
}

//c. Variáveis globais devem ser criadas para armazenar as entidades
var ps []Product
var lastID int

//d. A interface Repository deve ser gerada com todos os seus métodos
type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name string, typee string, count int,
		price float64) (Product, error)
	LastID() (int, error)
	Update(id int, name string, typee string, count int,
		price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) (error)
}

//e. A estrutura do repositório deve ser gerada
type repository struct {}

//f. Deve ser gerada uma função que retorne o Repositório
func NewRepository() Repository {
	return &repository{}
}

/*
** g. Todos os métodos correspondentes às operações a serem executadas
** (GetAll, Store, etc.) devem ser implementados.
*/

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, name string, typee string, count int,
	price float64) (Product, error) {
	p := Product{id, name, typee, count, price}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}

func (r *repository) Update(id int, name string, typee string, count int,
	price float64) (Product, error) {
	p := Product{id, name, typee, count, price}
	for i:= range ps {
		if ps[i].ID == id {
			ps[i] = p
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	for i:= range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			p := ps[i]
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}

func (r *repository) Delete(id int) (error) {
	for i:= range ps {
		if ps[i].ID == id {
			ps = append(ps[:i], ps[i + 1:]...)
			return nil
		}
	}
	return fmt.Errorf("produto %d não encontrado", id)
}

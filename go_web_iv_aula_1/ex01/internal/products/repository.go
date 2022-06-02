package products

import (
	"fmt"
	"github.com/anolivei/go_meli_go/tree/main/go_web_iv_aula_1/ex01/pkg/store"
)

type Product struct {
	ID		int		`json:"id"`
	Name	string	`json:"name"`
	Typee	string	`json:"type"`
	Count	int		`json:"count"`
	Price	float64	`json:"price"`
}

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

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps) - 1].ID, nil
}

func (r *repository) Store(id int, name string, typee string, count int,
	price float64) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{id, name, typee, count, price}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) Update(id int, name string, typee string, count int,
	price float64) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{id, name, typee, count, price}
	for i:= range ps {
		if ps[i].ID == id {
			ps[i] = p
			if err := r.db.Write(ps); err != nil {
				return Product{}, err
			}
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	for i:= range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			p := ps[i]
			if err := r.db.Write(ps); err != nil {
				return Product{}, err
			}
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("produto %d não encontrado", id)
}

func (r *repository) Delete(id int) (error) {
	var ps []Product
	r.db.Read(&ps)
	for i:= range ps {
		if ps[i].ID == id {
			ps = append(ps[:i], ps[i + 1:]...)
			if err := r.db.Write(ps); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("produto %d não encontrado", id)
}

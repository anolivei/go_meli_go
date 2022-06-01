//a. O arquivo service.go deve ser criado.
package products

//b. A interface Service com todos os seus métodos deve ser gerada.
type Service interface {
	GetAll() ([]Product, error)
	Store(name string, typee string, count int, price float64) (Product, error)
	Update(id int, name string, typee string, count int,
		price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) (error)
}

//c. A estrutura de serviço que contém o repositório deve ser gerada.
type service struct {
	repository Repository
}

//d. Deve ser gerada uma função que retorne o Serviço.
func NewService(r Repository) Service {
	return &service{repository: r}
}

/*
** e. Todos os métodos correspondentes às operações a serem executadas
**  (GetAll, Create, etc.) devem ser implementados.
*/
func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(name string, typee string, count int,
	price float64) (Product, error) {
	lastId, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}
	lastId++
	product, err := s.repository.Store(lastId, name, typee, count, price)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Update(id int, name string, typee string, count int,
	price float64) (Product, error) {
	product, err := s.repository.Update(id, name, typee, count, price)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) UpdateName(id int, name string) (Product, error) {
	product, err := s.repository.UpdateName(id, name)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Delete(id int) (error) {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

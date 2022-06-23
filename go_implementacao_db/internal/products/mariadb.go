package products

import (
	"database/sql"
)

const (
	sqlGetAll = "SELECT * FROM products"

	sqlStore = "INSERT INTO products (name, type, count, price) VALUES (?, ?, ?, ?)"

	sqlLastID = "SELECT MAX(id) as last_id FROM products"

	sqlUpdate = "UPDATE products SET name=?, type=?, count=?, price=? WHERE id=?"

	sqlUpdateName = "UPDATE products SET name=? WHERE id=?"

	sqlDelete = "DELETE FROM products WHERE id=?"
)

type mariaDBRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) Repository {
	return mariaDBRepository{db: db}
}

func (m mariaDBRepository) GetAll() ([]Product, error) {
	products := []Product{}

	rows, err := m.db.Query(sqlGetAll)
	if err != nil {
		return products, err
	}

	defer rows.Close() // Impedir vazamento de memória

	for rows.Next() {
		var product Product

		err := rows.Scan(&product.ID, &product.Name, &product.Typee, &product.Count, &product.Price)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (m mariaDBRepository) Store(
	id int,
	name,
	typee string,
	count int,
	price float64,
) (Product, error) {
	product := Product{
		Name:  name,
		Typee:  typee,
		Count: count,
		Price: price,
	}

	stmt, err := m.db.Prepare(sqlStore)
	if err != nil {
		return product, err
	}

	defer stmt.Close() // Impedir vazamento de memória

	res, err := stmt.Exec(&product.Name, &product.Typee, &product.Count, &product.Price)
	if err != nil {
		return product, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return product, err
	}

	product.ID = int(lastID)

	return product, nil
}

func (m mariaDBRepository) LastID() (int, error) {
	var maxCount int

	row := m.db.QueryRow(sqlLastID)

	err := row.Scan(&maxCount)
	if err != nil {
		return 0, err
	}

	return maxCount, nil
}

func (m mariaDBRepository) Update(
	id int,
	name, productType string,
	count int,
	price float64,
) (Product, error) {
	product := Product{
		ID:    id,
		Name:  name,
		Typee:  productType,
		Count: count,
		Price: price,
	}

	stmt, err := m.db.Prepare(sqlUpdate)
	if err != nil {
		return product, err
	}

	defer stmt.Close() // Impedir vazamento de memória

	_, err = stmt.Exec(
		&product.Name,
		&product.Typee,
		&product.Count,
		&product.Price,
		&product.ID,
	)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (m mariaDBRepository) UpdateName(id int, name string) (Product, error) {
	product := Product{ID: id, Name: name}

	stmt, err := m.db.Prepare(sqlUpdateName)
	if err != nil {
		return product, err
	}

	defer stmt.Close() // Impedir vazamento de memória

	_, err = stmt.Exec(&product.Name, &product.ID)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (m mariaDBRepository) Delete(id int) error {
	stmt, err := m.db.Prepare(sqlDelete)
	if err != nil {
		return err
	}

	defer stmt.Close() // Impedir vazamento de memória

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

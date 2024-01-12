package repository

import "github.com/agudelozca/productsgoweb11/internal"

// NewProductMap returns a new NewProductMap instance
func NewProductMap(db map[int]internal.Product, lastId int) *ProductMap {
	// default config / values
	// ...

	return &ProductMap{
		db:     db,
		lastId: lastId,
	}
}

// ProductMap is a struct that represents a Product repository
type ProductMap struct {
	// db is a map that represents a database
	// - key: id of the product
	// - value: product
	db map[int]internal.Product
	// lastId is the last id of the product
	lastId int
}

func (m *ProductMap) GetProducts() ([]internal.Product, error) {
	var products []internal.Product

	for _, product := range m.db {
		products = append(products, product)
	}

	return products, nil
}

func (m *ProductMap) Save(product *internal.Product) (err error) {
	// validate Product (consistency)
	// - title: unique
	if err = m.ValidateProductCodeValue((*product).CodeValue); err != nil {
		return
	}

	// autoincrement
	// - increment id
	(*m).lastId++
	// - set id
	(*product).ID = (*m).lastId

	// store product
	(*m).db[(*product).ID] = *product

	return
}

func (m *ProductMap) GetProductByID(id int) (product internal.Product, err error) {
	product, ok := m.db[id]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}

	return
}

func (m *ProductMap) ValidateProductCodeValue(code string) (err error) {
	// validate product (consistency)
	// - title: unique
	for _, v := range (*m).db {
		if v.CodeValue == code {
			return internal.ErrProductCodeValueAlreadyExists
		}
	}

	return
}

func (m *ProductMap) Update(product *internal.Product) (err error) {

	_, ok := m.db[(*product).ID]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}

	m.db[(*product).ID] = *product
	return
}

func (m *ProductMap) Delete(id int) (err error) {
	_, ok := m.db[id]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}

	delete(m.db, id)

	return
}

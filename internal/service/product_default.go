package service

import (
	"fmt"
	"time"

	"github.com/agudelozca/productsgoweb11/internal"
)

// NewProductDefault creates a new instance of a product service
func NewProductDefault(rp internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}
}

// ProductDefault is a struct that represents the default implementation of a product service
type ProductDefault struct {
	// rp is a product repository
	rp internal.ProductRepository
	// external services
	// ... (weather api, etc.)
}

// Save saves a product
func (m *ProductDefault) Save(product *internal.Product) (err error) {
	// external services
	// ...

	// business logic
	// - validate required fields
	if err = ValidateProduct(product); err != nil {
		return
	}

	// save product
	err = m.rp.Save(product)
	if err != nil {
		switch err {
		case internal.ErrProductCodeValueAlreadyExists:
			err = fmt.Errorf("%w: title", internal.ErrProductAlreadyExists)
		}
		return
	}

	return
}

func (m *ProductDefault) GetProducts() ([]internal.Product, error) {
	// Obtener todos los productos
	products, err := m.rp.GetProducts()
	if err != nil {
		switch err {
		// Puedes manejar diferentes errores específicos aquí si es necesario
		default:
			return nil, fmt.Errorf("error obteniendo todos los productos: %w", err)
		}
	}

	return products, nil
}

// GetByID returns a product by id
func (m *ProductDefault) GetProductByID(id int) (product internal.Product, err error) {
	// get product
	product, err = m.rp.GetProductByID(id)
	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		}
		return
	}
	return
}

func ValidateProduct(product *internal.Product) (err error) {
	// - validate required fields
	if (*product).Name == "" {
		return fmt.Errorf("%w: name", internal.ErrFieldRequired)
	}
	if (*product).Quantity == 0 {
		return fmt.Errorf("%w: quantity", internal.ErrFieldRequired)
	}

	// - validate data types
	if (*product).Price < 0 {
		return fmt.Errorf("%w: price", internal.ErrFieldTypeMismatch)
	}

	// - validate code_value uniqueness
	// Validate code_value uniqueness
	if !IsCodeValueUnique((*product).CodeValue) {
		return fmt.Errorf("%w: code_value", internal.ErrFieldNotUnique)
	}

	// - validate date format and values
	dateFormat := "01/02/2006"
	expirationTime, err := time.Parse(dateFormat, product.Expiration)
	if err != nil {
		return fmt.Errorf("%w: expiration", internal.ErrDateFormat)
	}

	// Validate expiration year
	if expirationTime.Year() < 0 {
		return fmt.Errorf("%w: expirationYear", internal.ErrFieldInvalid)
	}

	// Validate expiration month (1-12)
	if expirationTime.Month() < 1 || expirationTime.Month() > 12 {
		return fmt.Errorf("%w: expirationMonth", internal.ErrFieldInvalid)
	}

	// Validate expiration day (1-31)
	if expirationTime.Day() < 1 || expirationTime.Day() > 31 {
		return fmt.Errorf("%w: expirationDay", internal.ErrFieldInvalid)
	}

	return nil
}

// Update updates a product
func (m *ProductDefault) Update(product *internal.Product) (err error) {
	// validate
	if err = ValidateProduct(product); err != nil {
		return
	}

	// update product
	err = m.rp.Update(product)
	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		}
		return
	}
	return
}

func (m *ProductDefault) Delete(id int) (err error) {
	// delete product
	err = m.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrProductNotFound:
			err = fmt.Errorf("%w: id", internal.ErrProductNotFound)
		}
		return
	}
	return
}

func IsCodeValueUnique(codeValue string) bool {
	var products []internal.Product
	for _, product := range products {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}

package internal

import (
	"errors"
)

var (
	ErrFieldRequired        = errors.New("field required")
	ErrFieldQuality         = errors.New("field quality")
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrFieldTypeMismatch    = errors.New("field type mismatch")
	ErrFieldNotUnique       = errors.New("field not unique")
	ErrDateFormat           = errors.New("date format not valid")
	ErrDateValues           = errors.New("date value not valid")
	ErrFieldInvalid         = errors.New("field invalid")
)

// ProductService is an interface that represents a product service
// - business logic
// - validation
// - external services (e.g. apis, databases, etc.)
type ProductService interface {
	// Save saves a product
	Save(product *Product) (err error)
	//GetProducts returns a list of products from the repository
	GetProducts() ([]Product, error)
	// GetProductByID returns a product by id from the repository
	GetProductByID(id int) (product Product, err error)
	// Update updates a movie in the repository
	Update(product *Product) (err error)
	// Delete deletes a movie from the repository
	Delete(id int) (err error)
}

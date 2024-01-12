package internal

import "errors"

var (
	// ErrProducCodeValueAlreadyExists is the error returned when a movie title already exists
	ErrProductCodeValueAlreadyExists = errors.New("product code value already exists")
	ErrProductNotFound               = errors.New("product not found")
)

// MovieRepository is an interface that represents a movie repository
// - database
// - cache
type ProductRepository interface {
	// Save saves a product to the repository
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

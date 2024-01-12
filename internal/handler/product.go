package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/agudelozca/productsgoweb11/internal"
	"github.com/agudelozca/productsgoweb11/platform/web/request"
	"github.com/agudelozca/productsgoweb11/platform/web/response"
	"github.com/go-chi/chi/v5"
)

// DefaultProducts is an implementation with handlers for the Products storage
type DefaultProducts struct {
	// sv is a product service
	sv internal.ProductService
}

// NewDefaultProducts returns a new DefaultProducts instance
func NewDefaultProducts(sv internal.ProductService) *DefaultProducts {
	return &DefaultProducts{
		sv: sv,
	}
}

// NewControllerProducts returns a new ControllerProducts
func NewControllerProducts(storage map[int]*internal.Product, lastId int) *ControllerProducts {
	return &ControllerProducts{storage: storage, lastId: lastId}
}

// ControllerProducts is a struct that contains the storage of products
type ControllerProducts struct {
	storage map[int]*internal.Product
	lastId  int
}

// ProductJSON is the product in JSON format
type ProductJSON struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// BodyRequestProductJSON is the body of the request for a product in JSON format
type BodyRequestProductJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// PingHandler responde con "pong" y un código de estado 200 OK
func (d *DefaultProducts) PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}

// GetProducts devuelve la lista de todos los productos
// GetAll returns all products
func (d *DefaultProducts) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// process
		// - get all products
		products, err := d.sv.GetProducts()
		if err != nil {
			response.Text(w, http.StatusInternalServerError, "internal server error")
			return
		}

		// response
		// - serialize list of ProductJSON
		var data []ProductJSON
		for _, product := range products {
			data = append(data, ProductJSON{
				ID:          product.ID,
				Name:        product.Name,
				Quantity:    product.Quantity,
				CodeValue:   product.CodeValue,
				IsPublished: product.IsPublished,
				Expiration:  product.Expiration,
				Price:       product.Price,
			})
		}

		// - response
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"message": "products found",
			"data":    data,
		})
	}
}

// GetByID returns a movie by id
func (d *DefaultProducts) GetProductByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - get movie
		product, err := d.sv.GetProductByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize MovieJSON
		data := ProductJSON{
			ID:          product.ID,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}
		// - response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product found",
			"data":    data,
		})
	}
}

// Create creates a new product
// func (d *DefaultMovies) Create(w http.ResponseWriter, r *http.Request) {

// }
func (d *DefaultProducts) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var body BodyRequestProductJSON
		if err := request.JSON(r, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// process
		// - serialize internal.Product
		product := internal.Product{
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}
		// - save product
		if err := d.sv.Save(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrFieldQuality):
				// w.Header().Set("Content-Type", "text/plain")
				// w.WriteHeader(http.StatusBadRequest)
				// w.Write([]byte("invalid body"))
				response.Text(w, http.StatusBadRequest, "invalid body")
			case errors.Is(err, internal.ErrProductAlreadyExists):
				response.Text(w, http.StatusConflict, "product already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize ProductJSON
		data := ProductJSON{
			ID:          product.ID,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusCreated)
		// json.NewEncoder(w).Encode(map[string]any{
		// 	"message": "movie created",
		// 	"data": data,
		// })
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "product created",
			"data":    data,
		})
	}
}

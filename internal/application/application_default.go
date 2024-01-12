package application

import (
	//"log"
	"net/http"

	"github.com/agudelozca/productsgoweb11/internal"
	"github.com/agudelozca/productsgoweb11/internal/handler"
	"github.com/agudelozca/productsgoweb11/internal/repository"
	"github.com/agudelozca/productsgoweb11/internal/service"
	//"github.com/agudelozca/productsgoweb11/internal/storage"
	"github.com/go-chi/chi/v5"
)

type DefaultHTTP struct {
	// addr is the address of the http server
	addr string
}

// NewDefaultHTTP creates a new instance of a default http server
func NewDefaultHTTP(addr string) *DefaultHTTP {
	// default config / values
	// ...

	return &DefaultHTTP{
		addr: addr,
	}
}

// Run runs the http server
func (h *DefaultHTTP) Run() (err error) {

	// initialize dependencies
	// dependencies
	/*db, err := storage.LoaderProducts("./docs/db/json/products.json")
	if err != nil {
		log.Println(err)
		return
	}*/

	// - repository
	rp := repository.NewProductMap(make(map[int]internal.Product), 0)
	// - service
	sv := service.NewProductDefault(rp)
	// - handler
	hd := handler.NewDefaultProducts(sv)
	// - router
	rt := chi.NewRouter()
	//   endpoints
	rt.Get("/ping", hd.PingHandler())

	// with subrouter
	rt.Route("/products", func(rt chi.Router) {
		rt.Get("/", hd.GetProducts())
		rt.Get("/{id}", hd.GetProductByID())
		//rt.Get("/search", hd.SearchProductsByPrice())
		rt.Post("/", hd.Create())
	})

	// run http server
	err = http.ListenAndServe(h.addr, rt)
	return
}

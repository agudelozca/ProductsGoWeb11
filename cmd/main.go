package main

import (
	"fmt"
	"github.com/agudelozca/productsgoweb11/internal/application"
)

func main() {
	// env
	// ...

	// app
	// - config
	app := application.NewDefaultHTTP(":8080")
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}

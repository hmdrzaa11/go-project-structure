package routes

import (
	"net/http"

	"github.com/hmdrzaa11/example-go-api/pkg/handlers"
	"github.com/hmdrzaa11/example-go-api/pkg/kernel"
)

// LoadRoutes attach all the REST api routes to the router
func LoadRoutes(app *kernel.Application) {
	serviceRouter := app.Router.Methods(http.MethodGet).Subrouter() //to handle all "GET" verbs
	ph := handlers.NewProductHandler(app)

	//attach your routes
	serviceRouter.HandleFunc("/", ph.HelloWorld).Name("api:products")
}

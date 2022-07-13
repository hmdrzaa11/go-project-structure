package routes

import (
	"net/http"

	"github.com/hmdrzaa11/example-go-api/pkg/domains"
	"github.com/hmdrzaa11/example-go-api/pkg/handlers"
	"github.com/hmdrzaa11/example-go-api/pkg/kernel"
	"github.com/hmdrzaa11/example-go-api/pkg/services"
)

// LoadRoutes attach all the REST api routes to the router
func LoadRoutes(app *kernel.Application) {
	usersService := services.NewDefaultUserService(domains.NewUserRepository(app.DB))
	ph := handlers.NewProductHandler(app, usersService) //now we pass the service into handler

	app.Router.HandleFunc("/users", ph.HelloWorld).Methods(http.MethodGet)
}

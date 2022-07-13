package handlers

import (
	"net/http"

	"github.com/hmdrzaa11/example-go-api/pkg/kernel"
	"github.com/hmdrzaa11/example-go-api/pkg/services"
)

type Hello struct {
	app        *kernel.Application
	userServie services.UserService
}

func NewProductHandler(app *kernel.Application, userService services.UserService) *Hello {
	return &Hello{app, userService}
}

func (h *Hello) HelloWorld(w http.ResponseWriter, r *http.Request) {
	msg := struct {
		Message string `json:"message"`
	}{
		Message: "Hello World!",
	}
	h.app.Respond(w, http.StatusOK, msg)
}

package handlers

import (
	"net/http"

	"github.com/hmdrzaa11/example-go-api/pkg/kernel"
)

type Hello struct {
	app *kernel.Application
}

func NewProductHandler(app *kernel.Application) *Hello {
	return &Hello{app}
}

func (h *Hello) HelloWorld(w http.ResponseWriter, r *http.Request) {
	msg := struct {
		Message string `json:"message"`
	}{
		Message: "Hello World!",
	}
	h.app.Respond(w, http.StatusOK, msg)
}

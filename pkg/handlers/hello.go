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

func (p *Hello) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", p.app.Config.Http.Content)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from server!"))
}

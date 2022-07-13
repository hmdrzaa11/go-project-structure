package handlers

import (
	"net/http"
	"sync"

	"github.com/hmdrzaa11/example-go-api/pkg/apis"
)

type PhotosHandler struct {
	service *apis.Photos
}

func NewPhotoHandler(s *apis.Photos) *PhotosHandler {
	return &PhotosHandler{s}
}

func (ph *PhotosHandler) GetAllPhots(w http.ResponseWriter, r *http.Request) {
	//want to make a concurrent request
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ph.service.GetAllPhotos("https://jsonplaceholder.typeicode.com/photos", &wg)
	}()
	wg.Wait()

}

package apis

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Photos struct {
	client *http.Client
}

func NewPhotosApi() *Photos {
	return &Photos{
		client: &http.Client{
			Timeout: time.Second * 5, //only have 5 second to finsh the request
		},
	}
}

func (p *Photos) GetAllPhotos(url string, wg *sync.WaitGroup) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	fmt.Println(string(bs))
	wg.Done()
	return nil
}

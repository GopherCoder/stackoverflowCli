package infrastructure

import (
	"io/ioutil"

	"github.com/parnurzeal/gorequest"
)

func ResponseStorager(url string) ([]byte, error) {
	request := gorequest.New()
	response, _, _ := request.Get(url).End()
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

package mysportsfeeds

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseAddress = "https://www.mysportsfeeds.com/api/feed/pull"
	cacheFolder = "data"

	//Format x
	Format = "json"
)

//Client x
type Client interface {
	Request(subURL string, target interface{}) (int, error)
}

type webClient struct {
	auth      string
	webClient *http.Client
}

//NewWebClient x
func NewWebClient(username, password string) Client {
	r, _ := http.NewRequest("GET", "temp", nil)
	r.SetBasicAuth(username, password)
	auth := r.Header.Get("Authorization")

	c := &webClient{auth, &http.Client{}}
	return c
}

func (wc *webClient) Request(subURL string, target interface{}) (int, error) {
	fullURL := fmt.Sprintf("%s/%s", baseAddress, subURL)
	log.Println(fullURL)

	request, err := http.NewRequest("GET", fullURL, nil)
	request.Header.Set("Authorization", wc.auth)
	request.Header.Set("Accept-Encoding", "gzip")
	response, err := wc.webClient.Do(request)
	if err != nil {
		return response.StatusCode, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return response.StatusCode, errors.New("Not ok")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, err
	}

	json.Unmarshal(body, target)
	return response.StatusCode, nil
}

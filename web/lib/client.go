package lib

import (
	"io/ioutil"
	"net/http"
)

// HttpClient
type HttpClient struct {
	ApiHostName string
}

// NewHttpClient creates a new HTTP client
func NewHttpClient(apiHostName string) *HttpClient {
	return &HttpClient{
		ApiHostName: apiHostName,
	}
}

// Get does a get request to an endpoint
func (client *HttpClient) Get(url string) ([]byte, error) {

	cl := &http.Client{}

	req, err := http.NewRequest("GET", client.ApiHostName+url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-api-version", "2")

	if err != nil {
		return nil, err
	}

	res, err := cl.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

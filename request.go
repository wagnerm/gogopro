package gopro

import (
	"fmt"
	"net/http"
)

type APIRequester struct {
	URL       string
	BasicAuth *BasicAuth
	Client    *http.Client
}

func (ar *APIRequester) executeRequest(Method string, Endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(Method, ar.URL+Endpoint, nil)
	if err != nil {
		return nil, err
	}

	response, err := ar.Client.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(response)
	return response, nil
}

func (ar *APIRequester) get(Endpoint string) (*http.Response, error) {
	// Need to do something with payload
	resp, err := ar.executeRequest("GET", Endpoint)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

package gogopro

import (
	"fmt"
	"net/http"
	"strconv"
)

type APIRequester struct {
	URL       string
	BasicAuth *BasicAuth
	Client    *http.Client
}

func (ar *APIRequester) doRequest(Method string, RequestURL string) (*http.Response, error) {
	req, err := http.NewRequest(Method, RequestURL+"?"+fmt.Sprintf("t=%s", ar.BasicAuth.Password), nil)
	if err != nil {
		return nil, err
	}

	response, err := ar.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (ar *APIRequester) get(Endpoint string) (*http.Response, error) {
	resp, err := ar.doRequest("GET", ar.URL+Endpoint)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ar *APIRequester) getWithPort(Endpoint string, Port int) (*http.Response, error) {
	resp, err := ar.doRequest("GET", ar.URL+Endpoint+":"+strconv.Itoa(Port))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

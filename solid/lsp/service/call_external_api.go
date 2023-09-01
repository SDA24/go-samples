package call_external_api

import (
	"io"
	"net/http"
)

type Iapi interface {
	GetData() ([]byte, error)
}

type example_api struct {
	url string
}

func (e example_api) GetData() ([]byte, error) {
	resp, err := http.Get(e.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Request.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type another_example_api struct {
	url string
}

func (a another_example_api) GetData() ([]byte, error) {
	resp, err := http.Get(a.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type ApiFactory struct{}

func (a ApiFactory) CreateRepository(apiURL string) Iapi {
	if apiURL == "https://api.example.com/data" {
		return example_api{url: apiURL}
	} else if apiURL == "https://api.anotherexample.com/data" {
		return another_example_api{url: apiURL}
	}
	return nil
}

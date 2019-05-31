package rest

import (
	"fmt"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
)

// ClientInterface represents General HTTP Client to Continuum MS.
type ClientInterface interface {
	Get(url string) (resp *resty.Response, err error)
	Post(url string, body interface{}) (resp *resty.Response, err error)
	Put(url string, body interface{}) (resp *resty.Response, err error)
}

// clientType is a real ClientInterface implementation.
type clientType struct {
	Name string
	rest resty.Client
}

// NewClient creates and initializes a REST client to interact with different Continuum microservices.
func NewClient(name, baseURL string) (ClientInterface, error) {
	errMsg := fmt.Sprintf("failed to create %s client", name)

	_, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	client := new(clientType)

	client.Name = name
	rc := resty.New()
	rc.SetHostURL(baseURL)
	rc.RetryCount = 5
	rc.RetryWaitTime = 10 * time.Second
	rc.RetryMaxWaitTime = 10 * time.Second
	rc.SetHeader("Content-Type", "application/json")
	client.rest = *rc

	return client, nil
}

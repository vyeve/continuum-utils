package rest

import (
	"log"
	"net/http"

	"gopkg.in/resty.v1"
)

// execute is a general function that makes request
func (c clientType) execute(method, path string, body interface{}) (*resty.Response, error) {
	req := c.rest.R()

	req.SetBody(body)

	resp, err := req.Execute(method, path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// fmt.Println("URL:", resp.Request.URL)
	return resp, nil
}

// Get returns response from URL on GET operation
func (c clientType) Get(path string) (*resty.Response, error) {
	return c.execute(http.MethodGet, path, nil)
}

// Post returns response from URL on POST operation
func (c clientType) Post(path string, body interface{}) (resp *resty.Response, err error) {
	return c.execute(http.MethodPost, path, body)
}

// Put returns response from URL on PUT operation
func (c clientType) Put(path string, body interface{}) (resp *resty.Response, err error) {
	return c.execute(http.MethodPut, path, body)
}

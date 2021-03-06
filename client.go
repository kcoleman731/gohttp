package gohttp

import (
	"fmt"
	"net/http"
)

const (
	// GET is a constant for the HTTP GET method.
	GET = "GET"

	// POST is a constant for the HTTP POST method.
	POST = "POST"

	// DELETE is a constant for the HTTP DELETE method.
	DELETE = "DELETE"

	// PATCH is a constant for the HTTP PATCH method.
	PATCH = "PATCH"
)

// Client models an HTTP client
type Client struct {

	// BaseURL models the base URL to be used for all requests issues by
	// the Client.
	BaseURL string

	// Headers model the global headers to be used for all requests issued
	// by the client.
	Headers http.Header

	// Client is the underlying `http.Client` object that is used to issue
	// requests.
	Client *http.Client
}

// NewClient instantiates a new instance of a gohttp.Client.
func NewClient(baseURL string, headers http.Header) Client {
	if headers == nil {
		headers = http.Header{}
	}
	return Client{baseURL, headers, &http.Client{}}
}

func (c *Client) SetHeader(header string, value string) {
	c.Headers.Add(header, value)
}

// Execute executes the HTTP request described with the given `gohttp.Request`.
func (c *Client) Execute(r *Request) (*Response, error) {
	switch r.Method {
	case GET:
		return c.Get(r.URL)
	case POST:
		return c.Post(r.URL, r.Body)
	case DELETE:
		return c.Delete(r.URL)
	case PATCH:
		return c.Patch(r.URL, r.Body)
	}
	return nil, nil
}

// Get performs an HTTP GET request with the supplied URL string.
func (c *Client) Get(url string) (*Response, error) {
	URL := c.BaseURL + url
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header = c.Headers
	return performRequest(req, c.Client)
}

// Post performs an HTTP POST request with the supplied URL string and
// parameters.
func (c *Client) Post(url string, params interface{}) (*Response, error) {
	jsonData, err := JSONData(params)
	if err != nil {
		return nil, err
	}
	fmt.Printf("JSON Data: %v\n", jsonData)
	URL := c.BaseURL + url
	req, err := http.NewRequest("POST", URL, jsonData)
	if err != nil {
		return nil, err
	}

	req.Header = c.Headers
	return performRequest(req, c.Client)
}

// Delete performs an HTTP DELETE request with the supplied URL string.
func (c *Client) Delete(url string) (*Response, error) {
	URL := c.BaseURL + url
	req, err := http.NewRequest("Delete", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header = c.Headers
	return performRequest(req, c.Client)
}

// Patch performs an HTTP PATCH request with the supplied URL string and
// parameters.
func (c *Client) Patch(url string, params interface{}) (*Response, error) {
	jsonData, err := JSONData(params)
	if err != nil {
		return nil, err
	}

	URL := c.BaseURL + url
	req, err := http.NewRequest(PATCH, URL, jsonData)
	if err != nil {
		return nil, err
	}

	req.Header = c.Headers
	return performRequest(req, c.Client)
}

func performRequest(r *http.Request, c *http.Client) (*Response, error) {
	fmt.Printf("Performing request: %v\n", r)
	resp, err := c.Do(r)
	if err != nil {
		fmt.Printf("Error performing request response: %v\n", err)
		return nil, err
	}

	fmt.Printf("Got response: %v\n", resp)
	if (resp.StatusCode != http.StatusOK) && (resp.StatusCode != http.StatusCreated) && (resp.StatusCode != http.StatusNoContent) {
		return nil, fmt.Errorf("Unprocessable response code encountered: %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	return NewResponse(resp)
}

//------------------------------------------------------------------------------
// @Block Based Requests
//------------------------------------------------------------------------------

// func (c *Client) PostWithBlock(r Request, s func(http.Request, map[string]interface{}), f func(http.Request, error)) {
//
// }

// Package hparser provides a simple http wrapper that returns the usual HTTP
// response with a parsed HTML.
package hparser

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
)

// Client is an interface that will be used by the Get and PostForm functions.
// It allows to use a custom HTTP client with a middleware.
type Client interface {
	Get(url string) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// Default returns a basic http.Client with a cookie jar. It implements the
// Client interface.
func Default() *http.Client {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		panic(err) // Should never happen as the New function doesn't return an err.
	}
	return &http.Client{Jar: jar}
}

// Get issues a GET to the specified URL and parses the response body.
func Get(c Client, url string) (*http.Response, *html.Node, error) {
	r, err := c.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer r.Body.Close()
	p, err := html.Parse(r.Body)
	if err != nil {
		return nil, nil, err
	}
	return r, p, nil
}

// PostForm issues a POST to the specified URL, with data's keys and values
// URL-encoded as the request body. The response body is parsed and closed.
func PostForm(c Client, url string, data url.Values) (*http.Response, *html.Node, error) {
	r, err := c.PostForm(url, data)
	if err != nil {
		return nil, nil, err
	}
	defer r.Body.Close()
	p, err := html.Parse(r.Body)
	if err != nil {
		return nil, nil, err
	}
	return r, p, nil
}

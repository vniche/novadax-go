package client

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

// Client stands for the NOVADAX API http client
type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

// NewClient returns a new instance of NOVADAX API http client
func NewClient() *Client {
	novadaxURL := "https://api.novadax.com"

	if os.Getenv("NOVADAX_API_URL") != "" {
		novadaxURL = os.Getenv("NOVADAX_API_URL")
	}

	baseURL, err := url.ParseRequestURI(novadaxURL)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return &Client{
		httpClient: http.DefaultClient,
		BaseURL:    baseURL,
	}
}

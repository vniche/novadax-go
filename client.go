package novadax

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

// Config stands for the NOVADAX API config
type Config struct {
	AccessKey  string
	PrivateKey string
}

// Client stands for the NOVADAX API HTTP client
type Client struct {
	Config    *Config
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

var novadaxURL = "https://api.novadax.com"

// Default returns a new default instance of NOVADAX API http client
func Default() *Client {
	if os.Getenv("NOVADAX_API_URL") != "" {
		novadaxURL = os.Getenv("NOVADAX_API_URL")
	}

	baseURL, err := url.ParseRequestURI(novadaxURL)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return &Client{
		Config: &Config{
			AccessKey:  os.Getenv("NOVADAX_ACCESS_KEY"),
			PrivateKey: os.Getenv("NOVADAX_SECRET_KEY"),
		},
		httpClient: http.DefaultClient,
		BaseURL:    baseURL,
	}
}

// New returns a new instance of NOVADAX API http client
func New(accessKey string, privateKey string) *Client {
	client := Default()

	client.Config.AccessKey = accessKey
	client.Config.PrivateKey = privateKey

	return client
}

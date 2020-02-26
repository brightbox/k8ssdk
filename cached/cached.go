package cached

import (
	"net/http"

	brightbox "github.com/brightbox/gobrightbox"
)

// Client is a cached brightbox Client
type Client struct {
	brightbox.Client
}

// NewClient creates and returns a cached Client
func NewClient(url string, account string, httpClient *http.Client) (*Client, error) {
	cl, err := brightbox.NewClient(url, account, httpClient)
	return &Client{*cl}, err
}

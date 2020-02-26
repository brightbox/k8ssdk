package cached

import (
	brightbox "github.com/brightbox/gobrightbox"
)

// Client is a cached brightbox Client
type Client struct {
	brightbox.Client
}

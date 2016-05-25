package diffbot

import (
	"net/http"
)

// Config for the AlchemyAPI clients.
type Config struct {
	HTTPClient *http.Client
	Token      string
}

// NewConfig with the given access token.
func NewConfig(token string) *Config {
	return &Config{
		HTTPClient: http.DefaultClient,
		Token:      token,
	}
}

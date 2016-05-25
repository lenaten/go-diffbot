package diffbot

import "github.com/segmentio/go-env"

func client() *Client {
	token := env.MustGet("DIFFBOT_TOKEN")
	return New(NewConfig(token))
}

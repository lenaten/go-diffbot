package diffbot

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client implements a AlchemyAPI client.
type Client struct {
	*Config
}

// New client.
func New(config *Config) *Client {
	c := &Client{Config: config}
	return c
}

// call rpc style endpoint.
func (c *Client) call(api string, in map[string]string) (io.ReadCloser, error) {

	query := url.Values{}
	query.Set("token", c.Token)

	host := "http://api.diffbot.com/v3/"
	for key, value := range in {
		query.Set(key, value)
	}

	url := host + api + "?" + query.Encode()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	r, _, err := c.do(req)
	return r, err
}

// perform the request.
func (c *Client) do(req *http.Request) (io.ReadCloser, int64, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	if res.StatusCode < 400 {
		return res.Body, res.ContentLength, err
	}

	defer res.Body.Close()

	e := &Error{
		Status:     http.StatusText(res.StatusCode),
		StatusCode: res.StatusCode,
	}

	kind := res.Header.Get("Content-Type")

	if strings.Contains(kind, "text/plain") {
		if b, err := ioutil.ReadAll(res.Body); err == nil {
			e.Summary = string(b)
			return nil, 0, e
		}

		return nil, 0, err
	}

	if err := json.NewDecoder(res.Body).Decode(e); err != nil {
		return nil, 0, err
	}

	return nil, 0, e
}

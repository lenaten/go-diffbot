package diffbot

import "encoding/json"

// ArticleInput request output.
type ArticleInput struct {
	URL string `url:"url"`
}

type Object struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

// ArticleOutput request output.
type ArticleOutput struct {
	Objects []Object `json:"objects"`
}

func (c *Client) NewArticleInput(url string) *ArticleInput {
	return &ArticleInput{
		URL: url,
	}
}

// Article returns a ArticleOutput.
func (c *Client) Article(in *ArticleInput) (out *ArticleOutput, err error) {
	// v, _ := query.Values(in)
	body, err := c.call("article", map[string]string{
		"url": in.URL,
	})
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}

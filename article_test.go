package diffbot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArticle(t *testing.T) {
	c := client()
	url := "http://www.dailymail.co.uk/sciencetech/article-2355833/Apples-iPhone-5-hated-handset--majority-people-love-Samsung-Galaxy-S4-study-finds.html"
	in := c.NewArticleInput(url)
	out, err := c.Article(in)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, out.Objects)
	object := out.Objects[0]
	assert.NotEqual(t, "", object.Text)
	assert.NotEqual(t, "", object.Author)
	assert.NotEqual(t, "", object.Title)
}

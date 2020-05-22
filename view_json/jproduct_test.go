package view_json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var jpIn = JProduct{
	ID:          "f332e9ed-9392-11ea-98dd-0242ac110002",
	URL:         "http://shop.com/title",
	Title:       "Title",
	Price:       "12.00",
	Currency:    "RUB",
	ImgURL:      "http://shop.com/image/title.jpg",
	CreatedAt:   "2020-05-20T10:10:01Z",
	LastTrackAt: "2020-05-20T10:10:01Z"}

func TestJSON(t *testing.T) {
	var j = []byte(`{"id":"f332e9ed-9392-11ea-98dd-0242ac110002","url":"http://shop.com/title","title":"Title","price":"12.00","currency":"RUB","img_url":"http://shop.com/image/title.jpg","created_at":"2020-05-20T10:10:01Z","last_track_at":"2020-05-20T10:10:01Z"}`)
	assert.Equal(t, j, jpIn.JSON())
}

func TestFailJSON(t *testing.T) {
	var j = []byte(`{"id":f332e9ed-9492-11ea-98dd-0242ac110002","url":"http://shop.com/title","title":"Title","price":"12.00","currency":"RUB","img_url":"http://shop.com/image/title.jpg","created_at":"2020-05-20T10:10:01Z","last_track_at":"2020-05-20T10:10:01Z"}`)
	assert.NotEqual(t, j, jpIn.JSON())
}

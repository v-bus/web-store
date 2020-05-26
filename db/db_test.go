package db

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"time"
	jproduct "web-store/view_json"
)

var jpIn = jproduct.JProduct{
	ID:          "f332e9ed-9392-11ea-98dd-0242ac110002",
	URL:         "http://shop.com/title",
	Title:       "Title",
	Price:       "12.00",
	Currency:    "RUB",
	ImgURL:      "http://shop.com/image/title.jpg",
	CreatedAt:   "2020-05-20T10:10:01Z",
	LastTrackAt: "2020-05-20T10:10:01Z"}
var dbpIn = Product{
	ProductID: "f332e9ed-9392-11ea-98dd-0242ac110002",
	URL:       "http://shop.com/title",
	Title:     "Title",
	Price:     "12.00",
	Currency:  "RUB",
	ImgURL:    "http://shop.com/image/title.jpg"}

func TestNewDBProductFromJProduct(t *testing.T) {
	dbp := NewDBProductFromJProduct(&jpIn)
	assert.Equal(t, dbp.JProduct(), jpIn)
}

func TestJProduct(t *testing.T) {
	dbpIn.ProductCreatedAt, _ = time.Parse(time.RFC3339, "2020-05-20T10:10:01Z")
	dbpIn.ProductLastTrackAt, _ = time.Parse(time.RFC3339, "2020-05-20T10:10:01Z")
	assert.Equal(t, dbpIn.JProduct(), jpIn)
}

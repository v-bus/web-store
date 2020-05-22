package model

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/currency"

	jproduct "web-store/view_json"

	"github.com/google/uuid"
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

func TestNew(t *testing.T) {
	p := Product{}

	P := New()

	assert.NotEqual(t, p, P)
}

func TestNewProductFromJProduct(t *testing.T) {
	var p = Product{}
	var e error

	p.id, e = uuid.Parse(jpIn.ID)               //product id
	p.url = parseURL(jpIn.URL)                  //url
	p.title = jpIn.Title                        //title
	p.price = parsePrice(jpIn.Price)            //price
	p.currency = parseCurrency(jpIn.Currency)   //currency
	p.imgURL = parseURL(jpIn.ImgURL)            //image URL
	p.createdAt = parseTime(jpIn.CreatedAt)     //create_at
	p.lastTrackAt = parseTime(jpIn.LastTrackAt) //last_track_at
	if e != nil {
		assert.Error(t, e)
	}

	assert.Equal(t, NewProductFromJProduct(&jpIn).JSON(), jpIn.JSON())
}

func TestParseURL(t *testing.T) {
	var u = "http://shop.com/title"
	ua := parseURL(u)

	assert.Equal(t, u, ua)
}

func TestFailParseURL(t *testing.T) {
	var u = "345354345435:345354345:#45345345435:#$%3453454"
	ua := parseURL(u)

	assert.NotEqual(t, u, ua)
}

func TestParsePrice(t *testing.T) {
	var p = "12.00"
	pa := parsePrice(p)
	pr, err := decimal.NewFromString(p)
	assert.Equal(t, pr, pa)
	assert.NoError(t, err)
}

func TestFailParsePrice(t *testing.T) {
	var p = "12.00.12.12"
	pa := parsePrice(p)
	pr, err := decimal.NewFromString(p)
	assert.NotEqual(t, pr, pa)
	assert.Error(t, err)
}
func TestParseCurrency(t *testing.T) {
	var c = "EUR"
	ca := parseCurrency(c)
	cr, err := currency.ParseISO(c)
	assert.Equal(t, cr, ca)
	assert.NoError(t, err)
}
func TestFailParseCurrency(t *testing.T) {
	var c = "ird"
	ca := parseCurrency(c)
	cr, err := currency.ParseISO(c)
	assert.NotEqual(t, cr, ca)
	assert.Error(t, err)
}

func TestParseTime(t *testing.T) {
	var tm = "2020-05-20T10:10:01Z"
	tms := parseTime(tm)
	tmr, err := time.Parse(time.RFC3339, tm)
	assert.Equal(t, tmr, tms)
	assert.NoError(t, err)
}
func TestFailParseTime(t *testing.T) {
	var tm = "2020-05-2sdsdd10:01Z"
	tms := parseTime(tm)
	tmr, err := time.Parse(time.RFC3339, tm)
	assert.NotEqual(t, tmr, tms)
	assert.Error(t, err)
}

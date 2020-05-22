package model

import (
	"net/url"

	"time"

	jproduct "web-store/view_json"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/currency"
)

type (
	product struct {
		id          uuid.UUID
		url         string
		title       string
		price       decimal.Decimal
		currency    currency.Unit
		imgURL      string
		createdAt   time.Time
		lastTrackAt time.Time
	}
	// Product struct
	// id          uuid.UUID       |
	// url         string          |	"http://bing.com/search?q=dotnet" correct uri
	// title       string          |	"Title of the Product"
	// price       decimal.Decimal |	"12.11"
	// currency    currency.Unit   |	"EUR"
	// imgURL      string          |	"http://bing.com/search?q=dotnet" coorct uri
	// createdAt   time.Time       |	"2016-11-10T07:28:27.038Z"
	// lastTrackAt time.Time       |	"2016-11-10T07:28:27.038Z"
	Product product
)

// Creates new product struct
// only id, createdAt, last_track_at
func New() Product {
	var e error
	var p Product
	//generate random UUID Version 1
	p.id, e = uuid.NewUUID()
	if e != nil {
		log.Warnln("product New method UUID generator was broken with ", e)
	}
	p.createdAt = time.Now()
	p.lastTrackAt = time.Now()
	return p
}

func NewProductFromJProduct(jp *jproduct.JProduct) Product {
	var p product
	var e error

	p.id, e = uuid.Parse(jp.ID)               //product id
	p.url = parseURL(jp.URL)                  //url
	p.title = jp.Title                        //title
	p.price = parsePrice(jp.Price)            //price
	p.currency = parseCurrency(jp.Currency)   //currency
	p.imgURL = parseURL(jp.ImgURL)            //image URL
	p.createdAt = parseTime(jp.CreatedAt)     //create_at
	p.lastTrackAt = parseTime(jp.LastTrackAt) //last_track_at
	if e != nil {
		log.Warning(e)
	}
	return Product(p)
}

// parses url from json "url" or "img_url" element
// returns good url or clean url
func parseURL(jURL string) string {
	var s string
	u, err := url.Parse(jURL)
	if err != nil {
		log.Warning(err)
		s = new(url.URL).String() // clean url string
	} else {
		s = u.String() // good url string
	}
	return s
}

// parses price from json "price" to decimal "price"
// returns decimal price or Zero decimal price
func parsePrice(jPrice string) decimal.Decimal {
	price, err := decimal.NewFromString(jPrice)
	if err != nil {
		log.Warning(err)
		return decimal.Zero
	}
	return price

}

// parses currency from json "currency" field to currency.Unit "currency"
// return currency.Unit or currensy.TRY (Turkish lira)
func parseCurrency(jCurr string) currency.Unit {
	cur, err := currency.ParseISO(jCurr)
	if err != nil {
		log.Warning(err)
		return currency.TRY
	}
	return cur

}

// parse time from "create_at" and "last_track_at" fileds
// to time type RFC3339 of createAt and lastTrackAt 2016-11-10T07:28:27.038Z
// return time or time.Now (if error)
func parseTime(jTime string) time.Time {
	t, err := time.Parse(time.RFC3339, jTime)
	if err != nil {
		log.Warning(err)
		return time.Now()
	}
	return t
}

//returns JSON Product  JSON Marshalled Product
func (p Product) JSON() []byte {
	var jp jproduct.JProduct
	jp.ID = p.id.String()
	jp.URL = p.url
	jp.Title = p.title
	jp.Price = p.price.StringFixed(2)
	jp.Currency = p.currency.String()
	jp.ImgURL = p.imgURL
	jp.CreatedAt = p.createdAt.Format(time.RFC3339)
	jp.LastTrackAt = p.lastTrackAt.Format(time.RFC3339)

	return jp.JSON()
}

func (p *Product) ID() string {
	return p.id.String()
}

func (p *Product) URL() string {
	return p.url
}

func (p *Product) Title() string {
	return p.title
}

func (p *Product) Price() string {
	return p.price.StringFixed(2)
}

func (p *Product) Currency() string {
	return p.currency.String()
}

func (p *Product) ImgURL() string {
	return p.imgURL
}

func (p *Product) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Product) LastTrackAt() time.Time {
	return p.lastTrackAt
}

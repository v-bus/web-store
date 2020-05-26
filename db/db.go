package db

import (
	"time"

	product "web-store/model"
	jproduct "web-store/view_json"

	"github.com/jinzhu/gorm"
)

// Product describe GORM object to put to SQLite3 DB
type Product struct {
	gorm.Model
	ProductID          string    `gorm:"type:TEXT;UNIQUE_INDEX;NOT NULL"`
	URL                string    `gorm:"type:TEXT"`
	Title              string    `gorm:"type:TEXT"`
	Price              string    `gorm:"type:TEXT"`
	Currency           string    `gorm:"type:TEXT;size:3"`
	ImgURL             string    `gorm:"type:TEXT"`
	ProductCreatedAt   time.Time `gorm:"type:datetime"`
	ProductLastTrackAt time.Time `gorm:"type:datetime"`
	RawLastCall        time.Time `gorm:"type:datetime"`
}

// Users describe User accounts
type Users struct {
	gorm.Model
	User string `gorm:"type:TEXT"`
	Role string `gorm:"type:TEXT"`
}

//NewDBProductFromJProduct creates db.Product from view_json.Product
func NewDBProductFromJProduct(jp *jproduct.JProduct) Product {
	var db Product
	p := product.NewProductFromJProduct(jp)
	db.ProductID = p.ID()
	db.URL = p.URL()
	db.Title = p.Title()
	db.Price = p.Price()
	db.Currency = p.Currency()
	db.ImgURL = p.ImgURL()
	db.ProductCreatedAt = p.CreatedAt()
	db.ProductLastTrackAt = p.LastTrackAt()
	db.RawLastCall = time.Now()
	return db
}

//JProduct return view_json.JProduct object
func (db *Product) JProduct() jproduct.JProduct {
	var jp jproduct.JProduct

	jp.ID = db.ProductID
	jp.URL = db.URL
	jp.Title = db.Title
	jp.Price = db.Price
	jp.Currency = db.Currency
	jp.ImgURL = db.ImgURL
	jp.CreatedAt = db.ProductCreatedAt.Format(time.RFC3339)
	jp.LastTrackAt = db.ProductLastTrackAt.Format(time.RFC3339)
	db.RawLastCall = time.Now()
	return jp
}

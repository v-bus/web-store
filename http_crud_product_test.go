package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	dbproduct "web-store/db"
	"web-store/view_json"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/currency"
)

func TestMain(m *testing.M) {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	// calling method as a field, instruct the logger
	log.SetReportCaller(true)
	var err error
	db, err = gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()
	db.LogMode(false)
	// Migrate the schema
	if err := db.AutoMigrate(&dbproduct.DBProduct{}).Error; err != nil {
		log.Fatal(err)
	}
	exitVal := m.Run()

	os.Exit(exitVal)
}
func TestCreateProduct(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/product")

	// Assertions
	if assert.NoError(t, createProduct(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var resp = view_json.JProduct{}
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.ID)
		assert.Empty(t, resp.URL)
		assert.Empty(t, resp.Title)
		d, e := decimal.NewFromString("0.00")
		assert.NoError(t, e)
		assert.Equal(t, resp.Price, d.StringFixed(2))
		assert.Equal(t, resp.Currency, currency.TRY.String())
		assert.Empty(t, resp.ImgURL)
		assert.NotEmpty(t, resp.CreatedAt)
		assert.NotEmpty(t, resp.LastTrackAt)
	}
}
func TestFailCreateProduct(t *testing.T) {
	// Setup
	e := echo.New()
	db.DropTable(&dbproduct.DBProduct{})
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/product")

	// Assertions
	if assert.NoError(t, createProduct(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}

	db.AutoMigrate(&dbproduct.DBProduct{})
}

// deleteDBRecord to tier down records
func deleteDBRecord(t *testing.T, id string) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/product")
	c.SetParamNames("id")
	c.SetParamValues(id)

	log.Trace("TestDeleteProduct ID of product - ", fmt.Sprintf("/product/%s", id))
	if assert.NoError(t, deleteProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body.Bytes())
		log.Trace("TestDeleteProduct request - ", rec.Body.String())
		var jp struct {
			ID, Status, Reason string
		}
		if err := json.Unmarshal(rec.Body.Bytes(), &jp); err != nil {
			log.Trace("TestDeleteProduct json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
		}
		assert.Equal(t, jp.ID, id)
		assert.Equal(t, jp.Status, "deleted")
	}
}
func TestGetProduct(t *testing.T) {
	//Before Get we should create
	TestCreateProduct(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//find in DB first created product element
	var dbp []dbproduct.DBProduct
	if err := db.First(&dbp).Error; err == nil {
		c.SetPath("/product")
		c.SetParamNames("id")
		c.SetParamValues(dbp[0].ProductID)
		log.Trace("TestGetProduct sais - ", fmt.Sprintf("/product/%s", dbp[0].ProductID))
		if assert.NoError(t, getProduct(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotEmpty(t, rec.Body.Bytes())
			log.Trace("TestGetProduct sais - ", rec.Body.String())
			var jp view_json.JProduct
			if err := json.Unmarshal(rec.Body.Bytes(), &jp); err != nil {
				log.Trace("TestGetProduct json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
			}
			assert.Equal(t, jp.ID, dbp[0].ProductID)
		}
	}
	//tier down
	deleteDBRecord(t, dbp[0].ProductID)
}
func TestFailGetProduct(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/product")
	c.SetParamNames("id")
	c.SetParamValues("asdfghjkhgw65734254iuylukgdsfg")

	if assert.NoError(t, getProduct(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotEmpty(t, rec.Body.Bytes())
		log.Trace("TestFailGetProduct sais - ", rec.Body.String())
		var jp struct {
			ID, Status, Reason string
		}
		if err := json.Unmarshal(rec.Body.Bytes(), &jp); err != nil {
			log.Trace("TestFailGetProduct json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
		}
		assert.Equal(t, jp.Status, "error")
		assert.Equal(t, jp.Reason, "No such Product ID")
	}

}

func TestGetAllProducts(t *testing.T) {
	const dbdeep = 50
	for i := 0; i < dbdeep; i++ {
		TestCreateProduct(t)
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/product/all")
	var jps []view_json.JProduct
	if assert.NoError(t, getAllProducts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEmpty(t, rec.Body.Bytes())

		if err := json.Unmarshal(rec.Body.Bytes(), &jps); err != nil {
			log.Trace("TestGetAllProducts []view_json.JProduct sais - ", err.Error())
		}
		assert.NotEmpty(t, jps)
		for i, v := range jps {
			log.Trace("TestGetAllProducts returned array has ", i, " with ", v)
		}
		assert.Equal(t, len(jps), dbdeep)
	}

	//tier down
	for _, v := range jps {
		deleteDBRecord(t, v.ID)
	}
}

func TestFailGetAllProductsWithDB(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/product/all")
	if assert.NoError(t, getAllProducts(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotEmpty(t, rec.Body.Bytes())
		var jp struct {
			ID, Status, Reason string
		}
		if err := json.Unmarshal(rec.Body.Bytes(), &jp); err != nil {
			log.Trace("TestFailGetAllProductsWithDB json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
		}
		assert.Equal(t, jp.Status, "error")
		assert.Equal(t, jp.Reason, "No products")

	}
}

var jpIn = view_json.JProduct{
	ID:          "f332e9ed-9392-11ea-98dd-0242ac110002",
	URL:         "http://shop.com/title",
	Title:       "Title",
	Price:       "12.00",
	Currency:    "RUB",
	ImgURL:      "http://shop.com/image/title.jpg",
	CreatedAt:   "2020-05-20T10:10:01Z",
	LastTrackAt: "2020-05-20T10:10:01Z"}

func TestUpdateProduct(t *testing.T) {
	//Setup
	TestCreateProduct(t)

	//find in DB first created product element
	var dbp []dbproduct.DBProduct
	if err := db.First(&dbp).Error; err == nil {
		e := echo.New()

		jpIn.ID = dbp[0].ProductID
		log.Trace("TestUpdateProduct jpIn - ", jpIn)
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(jpIn.JSON()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("User", "admin")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/product")
		c.SetParamNames("id")
		c.SetParamValues(dbp[0].ProductID)

		log.Trace("TestUpdateProduct ID of product - ", fmt.Sprintf("/product/%s", dbp[0].ProductID))
		if assert.NoError(t, updateProduct(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotEmpty(t, rec.Body.Bytes())
			log.Trace("TestUpdateProduct request - ", rec.Body.String())
			var jp struct {
				ID, Status string
			}
			if err := json.Unmarshal(rec.Body.Bytes(), &jp); err != nil {
				log.Trace("TestUpdateProduct json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
			}
			assert.Equal(t, jp.ID, dbp[0].ProductID)
			assert.Equal(t, jp.Status, "success")
		}
		req = httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(jpIn.JSON()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recG := httptest.NewRecorder()
		cG := e.NewContext(req, recG)
		cG.SetPath("/product")
		cG.SetParamNames("id")
		cG.SetParamValues(dbp[0].ProductID)
		if assert.NoError(t, getProduct(cG)) {
			assert.Equal(t, http.StatusOK, recG.Code)
			assert.NotEmpty(t, recG.Body.Bytes())
			log.Trace("TestUpdateProduct in GetProduct request - ", recG.Body.String())
			var jp view_json.JProduct
			if err := json.Unmarshal(recG.Body.Bytes(), &jp); err != nil {
				log.Trace("TestUpdateProduct in getProduct json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
			}

			assert.Equal(t, jpIn.ID, jp.ID)
			assert.Equal(t, jpIn.URL, jp.URL)
			assert.Equal(t, jpIn.Price, jp.Price)
			assert.Equal(t, jpIn.Currency, jp.Currency)
			assert.Equal(t, jpIn.ImgURL, jp.ImgURL)
			assert.Equal(t, jpIn.Title, jp.Title)
		}
	}
	//tier down
	deleteDBRecord(t, dbp[0].ProductID)
}
func TestFailUpdateProductNoHeader(t *testing.T) {
	//Setup
	TestCreateProduct(t)

	//find in DB first created product element
	var dbp []dbproduct.DBProduct
	if err := db.First(&dbp).Error; err == nil {
		e := echo.New()

		jpIn.ID = dbp[0].ProductID
		log.Trace("TestUpdateProduct jpIn - ", jpIn)
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(jpIn.JSON()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/product")
		c.SetParamNames("id")
		c.SetParamValues(dbp[0].ProductID)

		log.Trace("TestUpdateProduct ID of product - ", fmt.Sprintf("/product/%s", dbp[0].ProductID))
		if assert.NoError(t, updateProduct(c)) {
			assert.Equal(t, http.StatusForbidden, rec.Code)
			assert.NotEmpty(t, rec.Body.Bytes())
			log.Trace("TestUpdateProduct request - ", rec.Body.String())
			var jp struct {
				ID, Status, Reason string
			}
			if err := json.Unmarshal(rec.Body.Bytes(), &jp); err != nil {
				log.Trace("TestUpdateProduct json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
			}
			assert.Equal(t, jp.ID, dbp[0].ProductID)
			assert.Equal(t, jp.Status, "error")
			assert.Equal(t, jp.Reason, "No user specified")
		}
		req = httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(jpIn.JSON()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recG := httptest.NewRecorder()
		cG := e.NewContext(req, recG)
		cG.SetPath("/product")
		cG.SetParamNames("id")
		cG.SetParamValues(dbp[0].ProductID)
		if assert.NoError(t, getProduct(cG)) {
			assert.Equal(t, http.StatusOK, recG.Code)
			assert.NotEmpty(t, recG.Body.Bytes())
			log.Trace("TestUpdateProduct in GetProduct request - ", recG.Body.String())
			var jp view_json.JProduct
			if err := json.Unmarshal(recG.Body.Bytes(), &jp); err != nil {
				log.Trace("TestUpdateProduct in getProduct json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
			}

			assert.Equal(t, jpIn.ID, jp.ID)
			assert.NotEqual(t, jpIn.URL, jp.URL)
			assert.NotEqual(t, jpIn.Price, jp.Price)
			assert.NotEqual(t, jpIn.Currency, jp.Currency)
			assert.NotEqual(t, jpIn.ImgURL, jp.ImgURL)
			assert.NotEqual(t, jpIn.Title, jp.Title)
		}
	}
	//tier down
	deleteDBRecord(t, dbp[0].ProductID)
}
func TestFailUpdateProductNoAdmin(t *testing.T) {
	//Setup
	TestCreateProduct(t)

	//find in DB first created product element
	var dbp []dbproduct.DBProduct
	if err := db.First(&dbp).Error; err == nil {
		e := echo.New()

		jpIn.ID = dbp[0].ProductID
		log.Trace("TestUpdateProduct jpIn - ", jpIn)
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(jpIn.JSON()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("User", "noadmin")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/product")
		c.SetParamNames("id")
		c.SetParamValues(dbp[0].ProductID)

		log.Trace("TestUpdateProduct ID of product - ", fmt.Sprintf("/product/%s", dbp[0].ProductID))
		if assert.NoError(t, updateProduct(c)) {
			assert.Equal(t, http.StatusForbidden, rec.Code)
			assert.NotEmpty(t, rec.Body.Bytes())
			log.Trace("TestUpdateProduct request - ", rec.Body.String())
			var jp struct {
				ID, Status, Reason string
			}
			if err := json.Unmarshal(rec.Body.Bytes(), &jp); err != nil {
				log.Trace("TestUpdateProduct json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
			}
			assert.Equal(t, jp.ID, dbp[0].ProductID)
			assert.Equal(t, jp.Status, "error")
			assert.Equal(t, jp.Reason, "Only admin can update product")
		}
		req = httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(jpIn.JSON()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recG := httptest.NewRecorder()
		cG := e.NewContext(req, recG)
		cG.SetPath("/product")
		cG.SetParamNames("id")
		cG.SetParamValues(dbp[0].ProductID)
		if assert.NoError(t, getProduct(cG)) {
			assert.Equal(t, http.StatusOK, recG.Code)
			assert.NotEmpty(t, recG.Body.Bytes())
			log.Trace("TestUpdateProduct in GetProduct request - ", recG.Body.String())
			var jp view_json.JProduct
			if err := json.Unmarshal(recG.Body.Bytes(), &jp); err != nil {
				log.Trace("TestUpdateProduct in getProduct json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
			}

			assert.Equal(t, jpIn.ID, jp.ID)
			assert.NotEqual(t, jpIn.URL, jp.URL)
			assert.NotEqual(t, jpIn.Price, jp.Price)
			assert.NotEqual(t, jpIn.Currency, jp.Currency)
			assert.NotEqual(t, jpIn.ImgURL, jp.ImgURL)
			assert.NotEqual(t, jpIn.Title, jp.Title)
		}
	}
	//tier down
	deleteDBRecord(t, dbp[0].ProductID)
}
func TestFailUpdateProductNoID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(nil))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("User", "admin")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ID := "234t5yhtgferghng"
	c.SetPath("/product/:id")
	c.SetParamNames("id")
	c.SetParamValues(ID)

	if assert.NoError(t, updateProduct(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotEmpty(t, rec.Body.Bytes())
		log.Trace("TestUpdateProduct request - ", rec.Body.String())
		var jp struct {
			ID, Status, Reason string
		}
		if err := json.Unmarshal(rec.Body.Bytes(), &jp); err != nil {
			log.Trace("TestUpdateProduct json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
		}
		assert.Equal(t, jp.ID, ID)
		assert.Equal(t, jp.Status, "error")
		assert.Equal(t, jp.Reason, "No such Product ID")
	}

}

func TestDeleteProduct(t *testing.T) {
	//Setup
	TestCreateProduct(t)

	//find in DB first created product element
	var dbp []dbproduct.DBProduct
	if err := db.First(&dbp).Error; err == nil {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewReader(nil))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/product")
		c.SetParamNames("id")
		c.SetParamValues(dbp[0].ProductID)

		log.Trace("TestDeleteProduct ID of product - ", fmt.Sprintf("/product/%s", dbp[0].ProductID))
		if assert.NoError(t, deleteProduct(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotEmpty(t, rec.Body.Bytes())
			log.Trace("TestDeleteProduct request - ", rec.Body.String())
			var jp struct {
				ID, Status, Reason string
			}
			if err := json.Unmarshal(rec.Body.Bytes(), &jp); err != nil {
				log.Trace("TestDeleteProduct json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
			}
			assert.Equal(t, jp.ID, dbp[0].ProductID)
			assert.Equal(t, jp.Status, "deleted")
		}
		req = httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(nil))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		recG := httptest.NewRecorder()
		cG := e.NewContext(req, recG)
		cG.SetPath("/product")
		cG.SetParamNames("id")
		cG.SetParamValues(dbp[0].ProductID)
		if assert.NoError(t, getProduct(cG)) {
			assert.Equal(t, http.StatusNotFound, recG.Code)
			assert.NotEmpty(t, recG.Body.Bytes())
			log.Trace("TestDeleteProduct in GetProduct request - ", recG.Body.String())
			var jp struct {
				ID, Status, Reason string
			}
			if err := json.Unmarshal(recG.Body.Bytes(), &jp); err != nil {
				log.Trace("TestDeleteProduct json.Unmarshal(rec.Body.Bytes()) sais error - ", err.Error())
			}
			assert.Equal(t, jp.ID, dbp[0].ProductID)
			assert.Equal(t, jp.Status, "error")
			assert.Equal(t, jp.Reason, "No such Product ID")
		}
	}
}

package main

import (
	"net/http"
	"time"

	dbproduct "web-store/db"
	product "web-store/model"
	jproduct "web-store/view_json"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type (
	// error message struct
	errorMsg struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		Reason string `json:"reason"`
	}
	// succes message
	scsRsp struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}
)

//----------
// Handlers
//----------
/**
 *
 * @api {POST} /product Create product
 * @apiName create_product
 * @apiGroup Products
 * @apiVersion  0.1.0
 *
 * @apiParam {String} name description
 * @apiParam {string} url URL to product page
 * @apiParam {string} title Product title
 * @apiParam {string} price Product price
 * @apiParam {string} currency Price currency (is global for current user seccion)
 * @apiParam {string} img_url URL to image of the product
 *
 * @apiSuccess (Product information) {string} id identifier of product
 * @apiSuccess (Product information) {string} url URL to product page
 * @apiSuccess (Product information) {string} title Product title
 * @apiSuccess (Product information) {string} price Product price
 * @apiSuccess (Product information) {string} currency Price currency (is global for current user seccion)
 * @apiSuccess (Product information) {string} img_url URL to image of the product
 * @apiSuccess (Product information) {string} created_at date and time of product SCU record creation Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
 * @apiSuccess (Product information) {string} last_buy_at date and time of the last order with this product Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
 *
 * @apiParamExample  {type} Request-Example:
 * {
 *	 "url":"http://shop.com/ipad",
 *	 "title":"IPad",
 *	 "price":"12.00",
 *	 "currency":"RUB",
 *	 "img_url":"http://shop.com/images/ipad.jpg"
 *	}
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *	 "id":"0ca17c65-98f8-11ea-9b20-04d4c44d69ba",
 *	 "url":"",
 *	 "title":"",
 *	 "price":"0.00",
 *	 "currency":"TRY",
 *	 "img_url":"",
 *	 "created_at":"2020-05-18T14:09:31+03:00",
 *	 "last_track_at":"2020-05-18T14:09:31+03:00"
 *	}
 *
 * @apiExample {curl} cURL Example  usage:
 *    curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" -d '{"url":"http://shop.com/ipad","title":"IPad","price":"12.00","currency":"RUB","img_url":"http://shop.com/images/ipad.jpg"}' "http://api.example.com/product"
 *
 * @apiUse ReqHeaders
 * @apiDescription Create new product
 */
func createProduct(c echo.Context) error {
	p := product.New()

	// Trace
	log.Traceln("Generated product id ", p.ID())
	log.Traceln("Generated product url ", p.URL())
	log.Traceln("Generated product title ", p.Title())
	log.Traceln("Generated product price ", p.Price())
	log.Traceln("Generated product currency ", p.Currency())
	log.Traceln("Generated product img_url", p.ImgURL())
	log.Traceln("Generated product createdAt", p.CreatedAt())
	log.Traceln("Generated product lastTrackAt", p.LastTrackAt())
	//jProduct init before binding
	jp := &jproduct.JProduct{
		ID:          p.ID(),
		CreatedAt:   p.CreatedAt().Format(time.RFC3339),
		LastTrackAt: p.LastTrackAt().Format(time.RFC3339),
	}
	// Trace
	log.Traceln("New jProduct id ", jp.ID)
	log.Traceln("New jProduct  url ", jp.URL)
	log.Traceln("New jProduct  title ", jp.Title)
	log.Traceln("New jProduct price ", jp.Price)
	log.Traceln("New jProduct currency ", jp.Currency)
	log.Traceln("New jProduct img_url", jp.ImgURL)
	log.Traceln("New jProduct createdAt", jp.CreatedAt)
	log.Traceln("New jProduct lastTrackAt", jp.LastTrackAt)
	defer recovery(c)
	// Bind jProduct
	if err := c.Bind(jp); err != nil {
		log.Warnln("getProduct returns ", err)
		f := errorMsg{
			ID:     p.ID(),
			Status: "fail",
			Reason: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, f)
	}
	// Trace
	log.Traceln("Binded jProduct id ", jp.ID)
	log.Traceln("Binded jProduct  url ", jp.URL)
	log.Traceln("Binded jProduct  title ", jp.Title)
	log.Traceln("Binded jProduct price ", jp.Price)
	log.Traceln("Binded jProduct currency ", jp.Currency)
	log.Traceln("Binded jProduct img_url", jp.ImgURL)
	log.Traceln("Binded jProduct createdAt", jp.CreatedAt)
	log.Traceln("Binded jProduct lasstTrackAt", jp.LastTrackAt)

	// Create product from jProduct
	dbp := dbproduct.NewDBProductFromJProduct(jp)
	// Trace
	log.Traceln("product before DB create id ", dbp.ProductID)
	log.Traceln("product before DB create url ", dbp.URL)
	log.Traceln("product before DB create title ", dbp.Title)
	log.Traceln("product before DB create price ", dbp.Price)
	log.Traceln("product before DB create currency ", dbp.Currency)
	log.Traceln("product before DB create img_url", dbp.ImgURL)
	log.Traceln("product before DB create createdAt", dbp.ProductCreatedAt)
	log.Traceln("product before DB create lasstTrackAt", dbp.ProductLastTrackAt)
	log.Traceln("DB ", db)
	defer recovery(c)
	dbp.DeletedAt = nil
	dbp.RawLastCall = time.Now()
	// Create DB record
	if err := db.Create(&dbp).GetErrors(); len(err) > 0 {
		for _, emsg := range err {
			log.Warning(emsg)
			f := errorMsg{
				ID:     p.ID(),
				Status: "fail",
				Reason: emsg.Error(),
			}
			return c.JSON(http.StatusInternalServerError, f)
		}
	}
	return c.JSON(http.StatusCreated, dbp.JProduct())
}

/**
 *
 * @api {GET} /product/:id Product information
 * @apiName product_item
 * @apiGroup Products
 * @apiVersion  0.1.0
 *
 *
 * @apiParam  {id} ID unique ID of an item in store
 *
 * @apiSuccess (Product information) {string} id identifier of product
 * @apiSuccess (Product information) {string} url URL to product page
 * @apiSuccess (Product information) {string} title Product title
 * @apiSuccess (Product information) {string} price Product price
 * @apiSuccess (Product information) {string} currency Price currency (is global for current user seccion)
 * @apiSuccess (Product information) {string} img_url URL to image of the product
 * @apiSuccess (Product information) {string} created_at date and time of product SCU record creation Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
 * @apiSuccess (Product information) {string} last_buy_at date and time of the last order with this product Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "id": "f332e9ed-9392-11ea-98dd-0242ac110002",
 *   "url": "http://shop.com",
 *   "title": "Title product",
 *   "price": "12.00",
 *   "currency": "RUB",
 *   "img_url": "http://shop.com",
 *   "created_at": "2020-05-11T14:23:14Z",
 *   "last_track_at": "2020-05-11T14:23:14Z"
 * }
 * @apiExample {curl} cURL Example  usage:
 *    curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" "http://api.example.com/product/f332e9ed-9392-11ea-98dd-0242ac110002"
 *
 * @apiUse ReqHeaders
 * @apiDescription Detailed Product Information
 */

func getProduct(c echo.Context) error {

	id := c.Param("id")

	e := errorMsg{
		ID:     id,
		Status: "error",
		Reason: "No such Product ID",
	}
	dp := new(dbproduct.Product)
	defer recovery(c)
	//SELECT
	if err := db.Where("product_id = ?", id).First(dp).Error; err != nil {
		return c.JSON(http.StatusNotFound, e)
	}

	return c.JSON(http.StatusOK, dp.JProduct())

}

/**
 *
 * @api {GET} /product/all Products list
 * @apiName product_list
 * @apiGroup Products
 * @apiVersion  0.1.0
 *
 * @apiSuccess (Product information) {string} id identifier of product
 * @apiSuccess (Product information) {string} url URL to product page
 * @apiSuccess (Product information) {string} title Product title
 * @apiSuccess (Product information) {string} price Product price
 * @apiSuccess (Product information) {string} currency Price currency (is global for current user seccion)
 * @apiSuccess (Product information) {string} img_url URL to image of the product
 * @apiSuccess (Product information) {string} created_at date and time of product SCU record creation Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
 * @apiSuccess (Product information) {string} last_buy_at date and time of the last order with this product Format: ISO 8601 YYYY-MM-DDTHH:MM:SSZ
 *
 * @apiSuccessExample {type} Success-Response:
 * [  {
 *  "id": "f2e19ada-9392-11ea-98dd-0242ac110002",
 *  "url": "http://shop.com",
 *  "title": "Title product",
 *  "price": "0.00",
 *  "currency": "TRY",
 *  "img_url": "http://shop.com",
 *  "created_at": "2020-05-11T14:23:13Z",
 *   "last_track_at": "2020-05-11T14:23:13Z"
 *  },
 *  {
 *  "id": "f332e9ed-9392-11ea-98dd-0242ac110002",
 *  "url": "http://shop.com",
 *  "title": "Title product",
 *  "price": "0.00",
 *  "currency": "TRY",
 *  "img_url": "http://shop.com",
 *  "created_at": "2020-05-11T14:23:14Z",
 *  "last_track_at": "2020-05-11T14:23:14Z"
 *  }]
 *
 * @apiExample {curl} cURL Example  usage:
 *    curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer $TOKEN" "http://api.example.com/product/all"
 *
 * @apiUse ReqHeaders
 * @apiDescription Product list
 */
func getAllProducts(c echo.Context) error {

	e := errorMsg{
		ID:     "all",
		Status: "error",
		Reason: "No products",
	}
	defer recovery(c)

	var res []dbproduct.Product
	if err := db.Find(&res).Error; err == nil {
		log.Traceln("getAllProduct result DBProduct{} array is ", res)
		log.Traceln("getAllProduct result DBProduct{} array length is ", len(res))
		var resj []jproduct.JProduct
		for _, el := range res {
			resj = append(resj, el.JProduct())
		}
		log.Traceln("getAllProduct result json array is ", resj)
		log.Traceln("getAllProduct result json array length is ", len(resj))
		if len(resj) > 0 {
			return c.JSON(http.StatusOK, resj)
		}

	} else {
		log.Warnln("WebStore DB getAllProducts said  - ", err.Error())
		return c.JSON(http.StatusInternalServerError, e)
	}
	return c.JSON(http.StatusNotFound, e)
}

/**
 *
 * @api {PUT} /product/:id Update Product Information
 * @apiName product_update
 * @apiGroup Products
 * @apiVersion  0.1.0
 *
 * @apiParam {url} url URL to product page
 * @apiParam {string} title Product title
 * @apiParam {money} price Product price
 * @apiParam {string} currency Price currency (is global for current user seccion)
 * @apiParam {url} img_url URL to image of the product
 *
 * @apiSuccess (200) {longint} id Product ID
 * @apiSuccess (200) {string} status Result Responce `(success|fail|error|auth_error)`.
 *
 * @apiParamExample  {type} Request-Example:
 * {
 *  "id": "f332e9ed-9392-11ea-98dd-0242ac110002",
 *  "url": "http://shop.com",
 *  "title": "Title product",
 *  "price": "0.00",
 *  "currency": "TRY",
 *  "img_url": "http://shop.com"
 * }
 *
 * @apiSuccessExample {json} Success-Response:
 *
 * {"id":"edf2f8a2-9392-11ea-98dd-0242ac110002","status":"success"}
 *
 * @apiErrorExample {json} Error-Response: Failed means application errors
 * {
 *      "id" : 34567876544
 *      "status": "fail"
 * }
 *
 * @apiErrorExample {json} Error-Response: Error means application errors
 *
 * {"id":"edf2f8a2-9392-11ea-98dd-0242ac11000","status":"error","reason":"No such Product ID"}
 *
 * @apiErrorExample {json} Error-Response: No_User means no user was set in Header
 * {
 *	 "id":"f31fd9f3-98ce-11ea-ab73-0242ac110002",
 *   "status":"error",
 *   "reason":"No user specified"
 * }
 * @apiErrorExample {json} Error-Response: Auth_Error means administrator authentication errors
 * {
 *	 "id":"f31fd9f3-98ce-11ea-ab73-0242ac110002",
 *	 "status":"error",
 *	 "reason":"Only admin can update product"
 * }
 * @apiExample {curl} cURL Example  usage:
 *    curl -v -X PUT -H "User: admin" -H "Authorization: Bearer 1234567890" -H "Content-Type: application/json" -d '{"url":"http://shop.com/my","title":"Title product","price":"0.00","currency":"TRY","img_url":"http://shop.com"}' "http://api.example.com/product/edf2f8a2-9392-11ea-98dd-0242ac110002"
 *
 * @apiUse ReqHeaders
 * @apiDescription Update Product information. PUT `id` and new information structure.  Admin user only can update information.
 */
func updateProduct(c echo.Context) error {
	if header := c.Request().Header; len(header["User"]) <= 0 {
		log.Warnln("Header User not found")
		return c.JSON(http.StatusForbidden, errorMsg{ID: c.Param("id"), Status: "error", Reason: "No user specified"})
	} else if a := getUsersByRole("admin"); len(a) > 0 && header["User"][0] != a[0] {
		log.Warnln("User is not admin")
		return c.JSON(http.StatusForbidden, errorMsg{ID: c.Param("id"), Status: "error", Reason: "Only admin can update product"})
	}
	jp := new(jproduct.JProduct)
	id := c.Param("id")
	jp.ID = id
	if err := c.Bind(jp); err != nil {

		log.Warnln("updateProduct returns ", err)
		f := errorMsg{
			ID:     jp.ID,
			Status: "fail",
			Reason: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, f)
	}

	dp := new(dbproduct.Product)
	defer recovery(c)
	if err := db.Where("product_id == ?", id).First(dp).Error; err != nil {
		log.Warnln("updateProduct DB said ", err)
		e := errorMsg{
			ID:     id,
			Status: "error",
			Reason: "No such Product ID",
		}
		return c.JSON(http.StatusNotFound, e)
	}

	jdp := dbproduct.NewDBProductFromJProduct(jp)
	jdp.RawLastCall = time.Now()
	if err := db.Model(&dbproduct.Product{}).Omit(
		"product_id",
		"created_at",
		"deleted_at",
		"product_created_at",
		"product_last_track_at").Update(jdp).Error; err != nil {
		f := errorMsg{
			ID:     id,
			Status: "fail",
			Reason: "WebStore DB fail 1",
		}
		return c.JSON(http.StatusInternalServerError, f)
	}

	return c.JSON(http.StatusOK, scsRsp{ID: id, Status: "success"})
}

/**
 *
 * @api {DELETE} /product/:id Delete product
 * @apiName delete_product
 * @apiGroup Products
 * @apiVersion  0.1.0
 *
 *
 * @apiParam  {String} id Product ID
 *
 * @apiSuccess (Status) {string} id Product ID
 * @apiSuccess (Status) {string} status Delete status should be "deleted"
 *
 *
 * @apiSuccessExample {type} Success-Response:
 * {
 *	 "id":"0ca17c65-98f8-11ea-9b20-04d4c44d69ba",
 *	 "status":"deleted"
 *	}
 *
 * @apiErrorExample {json} Error-Response: No product in store
 * {
 *	 "id":"f31fd9f3-98ce-11ea-ab73-0242ac110002",
 *	 "status":"error",
 *	 "reason":"No such Product ID"
 * }
 * @apiExample {curl} cURL Example  usage:
 *    curl -v -X DELETE -H "Authorization: Bearer 1234567890" -H "Content-Type: application/json" "http://api.example.com/product/edf2f8a2-9392-11ea-98dd-0242ac110002"
 *
 * @apiUse ReqHeaders
 * @apiDescription Update Product information. PUT `id` and new information structure.  Admin user only can update information.
 */
func deleteProduct(c echo.Context) error {
	id := c.Param("id")
	defer recovery(c)
	dp := new(dbproduct.Product)
	if err := db.Where("product_id == ?", id).First(dp).Error; err != nil {
		e := errorMsg{
			ID:     id,
			Status: "error",
			Reason: "No such Product ID",
		}
		return c.JSON(http.StatusNotFound, e)
	}

	if err := db.Delete(dp).Error; err != nil {
		log.Warnln("deleteProduct DB said - ", err)
		f := errorMsg{
			ID:     id,
			Status: "fail",
			Reason: "WebStore DB fail 2",
		}
		return c.JSON(http.StatusOK, f)
	}
	return c.JSON(http.StatusOK, scsRsp{ID: id, Status: "deleted"})
}

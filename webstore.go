package main

import (
	"crypto/subtle"
	"net/http"
	"os"
	"runtime/debug"
	dbproduct "web-store/db"
	"web-store/defines"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB
var web *echo.Echo
var version = "undefined"

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(defines.Log_level)

	// calling method as a field, instruct the logger
	log.SetReportCaller(true)
	log.Infoln("App version is ", version)
}
func main() {
	var err error
	db, err = gorm.Open("sqlite3", defines.DBFileName)
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()
	db.LogMode(defines.Logging)
	// Migrate the schema

	log.Trace("Check Users table ...")
	if db.HasTable(&dbproduct.Users{}) {
		dbu := dbproduct.Users{}
		log.Trace("Try create admin admin ...")
		if err := db.FirstOrCreate(&dbu, dbproduct.Users{User: "admin", Role: "admin"}).Error; err != nil {
			log.Warn(err)
		}
		log.Trace("admin admin was created ...")
	} else {
		log.Trace("Try to AutoMigrate Users schema ...")
		if err := db.AutoMigrate(&dbproduct.Users{}).Error; err != nil {
			log.Fatal(err)
		} else {
			log.Trace("Try to create admin admin Users record in DB ...")
			if err := db.Create(&dbproduct.Users{User: "admin", Role: "admin"}).Error; err != nil {
				log.Fatal(err)
			}
			log.Trace("admin admin was created success")
		}
		log.Trace("Schema Users was migrated success")
	}
	if !db.HasTable(&dbproduct.Product{}) {
		log.Trace("Try to AutoMigrate Products schema...")
		if err := db.AutoMigrate(&dbproduct.Product{}).Error; err != nil {
			log.Fatal(err)
		}
		log.Trace("Schema Products was created OK")
	}
	// init WebServer instance
	web = echo.New()
	// Middleware
	web.Use(middleware.Logger())
	web.Use(middleware.Recover())

	// Routes
	web.POST("/product", createProduct)
	web.GET("/product/:id", getProduct)
	web.GET("/product/all", getAllProducts)
	web.PUT("/product/:id", updateProduct)
	web.DELETE("/product/:id", deleteProduct)

	web.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to WebStore API!</h1>
		`)
	})
	web.Use(middleware.KeyAuth(func(token string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(token), []byte("1234567890")) == 1 {
			return true, nil
		}
		return false, nil
	}))
	// Start server
	web.Logger.Fatal(web.Start(":8080"))
}
func recovery(c echo.Context) {
	if r := recover(); r != nil {
		log.Warnln("recovery runtimerror: ", r, string(debug.Stack()))
		e := errorMsg{
			ID:     uuid.Nil.String(),
			Status: "fail",
			Reason: "WebStore Internal server error",
		}
		_ = c.JSON(http.StatusInternalServerError, e)
	}
}

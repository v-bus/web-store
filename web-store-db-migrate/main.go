package main

import (
	"fmt"
	"os"
	"time"

	dbproduct "web-store/db"
	"web-store/defines"

	"github.com/hashicorp/go-version"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

var appVersion = "undifined"

const migrateVersion = "0.1.0"

//DBMigrateVersion  store data of last migrations
type DBMigrateVersion struct {
	gorm.Model
	Version        string    `gorm:"type:TEXT"`
	LastUpdateTime time.Time `gorm:"type:datetime"`
}

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
	log.Infoln("App version is ", appVersion)
}

//dbOpen opens DB connection
//defer db.Close() see in main() func
func dbOpen() {
	var err error
	db, err = gorm.Open("sqlite3", defines.DBFileName)
	if err != nil {
		panic("failed to connect database")
	}
}

//migrateUsers add db.DBProduct table to DB
func migrateUsers() {
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
}
func migrateProducts() {
	if !db.HasTable(&dbproduct.Product{}) {
		log.Trace("Try to AutoMigrate Products schema...")
		if err := db.AutoMigrate(&dbproduct.Product{}).Error; err != nil {
			log.Fatal(err)
		}
		log.Trace("Schema Products was created OK")
	}
	log.Trace("No need to create dbproduct.Product schema it's already exists")
}

//addNewVersion add new DBMigrateVersion record to DB
func addNewVersion() {
	v, _ := version.NewVersion(migrateVersion)
	var dbmv = DBMigrateVersion{
		Version:        v.String(),
		LastUpdateTime: time.Now(),
	}
	if err := db.Create(&dbmv).Error; err != nil {
		log.Fatal(err)
	}
}

//migrateDBMigrateVersion add DBMigrateVersion table to DB
// and new DBMigrateVersion record
func migrateDBMigrateVersion() {
	if !db.HasTable(&DBMigrateVersion{}) {
		if err := db.AutoMigrate(&DBMigrateVersion{}).Error; err != nil {
			log.Fatal(err)
		}
	}
	addNewVersion()
}

//DBProduct version v0.0.0 of product table
type DBProduct struct {
	gorm.Model
	ProductID          string    `gorm:"type:TEXT;UNIQUE_INDEX;NOT NULL"`
	URL                string    `gorm:"type:TEXT"`
	Title              string    `gorm:"type:TEXT"`
	Price              string    `gorm:"type:TEXT"`
	Currency           string    `gorm:"type:TEXT;size:3"`
	ImgURL             string    `gorm:"type:TEXT"`
	ProductCreatedAt   time.Time `gorm:"type:datetime"`
	ProductLastTrackAt time.Time `gorm:"type:datetime"`
}

//dbProductExists returns true if product table v0.0.0 exists
func dbProductExists() bool {
	return db.HasTable(DBProduct{})
}

//copy all records from product table v0.0.0 to product table v0.1.0
func copyDBProductToProduct() []error {
	var err []error
	if dbProductExists() {
		var dbp = []DBProduct{}
		if e := db.Find(&dbp).Error; e != nil {
			err = append(err, e)
			log.Warn(e)
		} else if db.HasTable(dbproduct.Product{}) {
			for _, rec := range dbp {
				var p dbproduct.Product
				p.Currency = rec.Currency
				p.ImgURL = rec.ImgURL
				p.Price = rec.Price
				p.ProductCreatedAt = rec.ProductCreatedAt
				p.ProductLastTrackAt = rec.ProductLastTrackAt
				p.ProductID = rec.ProductID
				p.Title = rec.Title
				p.URL = rec.URL
				p.RawLastCall = time.Now()
				if e := db.Create(&p).Error; e != nil {
					err = append(err, e)
					log.Warn(e)
				}
			}
		} else {
			err = append(err, fmt.Errorf("No table product was found"))
			log.Warn("No table product was found")
		}
	}
	return err
}
func main() {

	dbOpen()
	defer db.Close()
	db.LogMode(defines.Logging)

	// Migrate the schema

	migrateUsers()
	migrateProducts()

	if !db.HasTable(DBMigrateVersion{}) {
		copyDBProductToProduct()
	}
	migrateDBMigrateVersion()
}

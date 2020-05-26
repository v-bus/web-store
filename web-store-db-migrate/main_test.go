package main

import (
	"fmt"
	"os"
	"testing"
	"time"
	dbproduct "web-store/db"
	"web-store/defines"

	"github.com/google/uuid"
	"github.com/hashicorp/go-version"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	os.Remove(defines.DBFileName)
	main()
	os.Remove(defines.DBFileName)
}

func Test_dbOpen(t *testing.T) {
	//setup
	os.Remove(defines.DBFileName)
	dbOpen()
	assert.FileExists(t, defines.DBFileName)

	//teardown
	os.Remove(defines.DBFileName)
}

func Test_migrateUsers(t *testing.T) {
	//setup
	os.Remove(defines.DBFileName)

	t.Run("Test AutoMigrate", func(t *testing.T) {
		//setup
		dbOpen()
		//test
		migrateUsers()
		assert.True(t, db.HasTable(&dbproduct.Users{}))

		var dbfu = dbproduct.Users{}
		if err := db.First(&dbfu, "user = 'admin'").Error; err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, dbfu.User, "admin")
		assert.Equal(t, dbfu.Role, "admin")

	})
	t.Run("Test FirstOrCreate", func(t *testing.T) {
		//test
		migrateUsers()
		var dbfu = []dbproduct.Users{}
		if err := db.Find(&dbfu, "user = 'admin'").Error; err != nil {
			log.Fatal(err)
		}
		assert.True(t, len(dbfu) == 1)
		//teardown
		os.Remove(defines.DBFileName)
	})
	t.Run("Test HasTable", func(t *testing.T) {
		//setup
		dbOpen()
		if err := db.AutoMigrate(&dbproduct.Users{}).Error; err != nil {
			log.Fatal(err)
		}
		migrateUsers()
		var dbfu = dbproduct.Users{}
		if err := db.First(&dbfu, "user = 'admin'").Error; err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, dbfu.User, "admin")
		assert.Equal(t, dbfu.Role, "admin")
	})
	//teardown
	os.Remove(defines.DBFileName)
}

func Test_migrateProducts(t *testing.T) {
	//setup
	dbOpen()
	//test
	migrateProducts()
	assert.True(t, db.HasTable(dbproduct.Product{}))
	//teardown
	os.Remove(defines.DBFileName)
}

func Test_migrateDBMigrateVersion(t *testing.T) {
	//setup
	dbOpen()
	//test
	migrateDBMigrateVersion()
	var dbfv = DBMigrateVersion{}
	if err := db.First(&dbfv).Error; err != nil {
		log.Fatal(err)
	}
	assert.NotEmpty(t, dbfv)
	v, _ := version.NewVersion(migrateVersion)
	assert.Equal(t, dbfv.Version, v.String())
	//teardown
	os.Remove(defines.DBFileName)
}

func Test_dbProductExists(t *testing.T) {
	//setup
	dbOpen()

	if err := db.AutoMigrate(&DBProduct{}).Error; err != nil {
		log.Fatal(err)
	}
	assert.True(t, dbProductExists())
	//teardown
	os.Remove(defines.DBFileName)
}

func Test_copyDBProductToProduct(t *testing.T) {
	//setup
	dbOpen()
	const dbpCnt = 50
	if err := db.AutoMigrate(&DBProduct{}).Error; err != nil {
		log.Fatal(err)
	}
	assert.True(t, dbProductExists())

	for i := 0; i < dbpCnt; i++ {
		var dbp = DBProduct{
			ProductID:          uuid.New().String(),
			ImgURL:             fmt.Sprintf("https://shop.com/image/title_%d", i),
			URL:                fmt.Sprintf("https://shop.com/title_%d", i),
			Price:              fmt.Sprintf("%d.%d", i, i),
			Currency:           "RUB",
			Title:              fmt.Sprintf("Title_%d", i),
			ProductCreatedAt:   time.Now(),
			ProductLastTrackAt: time.Now(),
		}
		if err := db.Create(&dbp).Error; err != nil {
			log.Fatal(err)
		}
	}

	migrateProducts()

	//test
	copyDBProductToProduct()
	var dbpt = []DBProduct{}
	var pt = []dbproduct.Product{}
	if err := db.Find(&dbpt).Error; err != nil {
		log.Fatal(err)
	} else if err := db.Find(&pt).Error; err != nil {
		log.Fatal(err)
	} else {
		for i, pte := range pt {
			dbpte := dbpt[i]
			assert.Equal(t, pte.ProductID, dbpte.ProductID)
			assert.Equal(t, pte.ImgURL, dbpte.ImgURL)
			assert.Equal(t, pte.URL, dbpte.URL)
			assert.Equal(t, pte.Title, dbpte.Title)
			assert.Equal(t, pte.Price, dbpte.Price)
			assert.Equal(t, pte.Currency, dbpte.Currency)
			assert.Equal(t, pte.ProductCreatedAt, dbpte.ProductCreatedAt)
			assert.Equal(t, pte.ProductLastTrackAt, dbpte.ProductLastTrackAt)
			log.Warn("assert equal after copyDBProductToProduct ", i)
		}
	}
	//teardown
	os.Remove(defines.DBFileName)
}

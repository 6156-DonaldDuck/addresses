package db

import (
	"fmt"
	"time"

	"github.com/6156-DonaldDuck/addresses/pkg/config"
	"github.com/6156-DonaldDuck/addresses/pkg/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbConn *gorm.DB

func init() {
	host := config.Configuration.Mysql.Host
	port := config.Configuration.Mysql.Port
	username := config.Configuration.Mysql.Username
	password := config.Configuration.Mysql.Password
	databaseName := config.Configuration.Mysql.DatabaseName
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, databaseName)
	DbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("[db.init] error occurred while creating database connection, err=%v\n", err)
	}

	createAddressTable()
}

func createAddressTable() {
	tableName := "addresses"
	if !DbConn.Migrator().HasTable(tableName) {
		log.Infof("[db.createTables] table %v not found, creating new one\n", tableName)
		if err := DbConn.Migrator().CreateTable(&model.Address{}); err != nil {
			log.Errorf("[db.createTables] error occurred while creating table %v, err=%v\n", tableName, err)
		}

		// insert test data
		testData := model.Address{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
			},
			StreetName1: "200 W 116 St",
			StreetName2: "",
			City:        "New York",
			Region:      "NY",
			CountryCode: "1",
			PostalCode:  "10031",
			UserId:      1,
		}
		result := DbConn.Create(&testData)
		if result.Error != nil {
			log.Errorf("[db.createTables] error occurred while inserting test data, err=%v\n", result.Error)
		} else {
			log.Infof("[db.createTables] successfully inserted test data, rows affected=%v\n", result.RowsAffected)
		}
	}
}

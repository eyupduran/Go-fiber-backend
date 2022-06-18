package database

import (
	"log"
	"os"

	"github.com/eyupduran/fiber-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}
//MySql connecting
const DNS = "root:392743.eE@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the database succesfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	//TODO:Add Migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	Database = DbInstance{
		Db: db,
	}
}
//Other connection
//--------------------------------------------
// var DB *gorm.DB
// var err error

// const DNS = "root:392743.eE@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

// func ConnectDb() {
// 	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Cannot connect to Database")
// 	}
// 	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
// }

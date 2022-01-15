package database

import (
	"log"
	"os"
	"testing"

	"git.xenonstack.com/util/continuous-security-backend/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	os.Remove(os.Getenv("HOME") + "/account-testing.db")
	db, err := gorm.Open("sqlite3", os.Getenv("HOME")+"/account-testing.db")
	if err != nil {
		log.Println(err)
		log.Println("Exit")
		os.Exit(1)
	}
	config.DB = db
	//create database
	CreateDatabase()
	//create table
	CreateDBTablesIfNotExists()

	requestinfo := RequestInfo{}
	requestinfo.Email = "test@xenonstack.com"
	requestinfo.ID = 1

	db.Create(&requestinfo)

	requestinfo2 := RequestInfo{}
	requestinfo2.Email = "testing@xenonstack.com"
	requestinfo2.ID = 2

	db.Create(&requestinfo2)
}

func TestSaveRow(t *testing.T) {
	mapd := make(map[string]interface{})
	mapd["secure"] = "true"
	mapd["description"] = "SSL is available for this Site"
	SaveRow(mapd, "", "", "", true)
}

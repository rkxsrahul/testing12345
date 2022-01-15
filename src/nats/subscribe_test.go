package nats

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"git.xenonstack.com/util/continuous-security-backend/config"
	"git.xenonstack.com/util/continuous-security-backend/src/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nats-io/nats.go"
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

	//create table
	database.CreateDBTablesIfNotExists()
}

func TestSaveToDatabase(t *testing.T) {
	mapd := make(map[string]interface{})
	var (
		testingdata  nats.Msg
		result       Result
		testingdata2 nats.Msg
	)

	mapd["test"] = "testing"
	result.Data = mapd
	result.Header = "ugtwed"
	result.Method = "oeih"
	result.Status = true
	//covert data into byte formate
	bytedata, _ := json.Marshal(result)
	testingdata.Data = bytedata
	//test save git result
	saveGitResult(&testingdata)
	//test save web result
	saveToDatabase(&testingdata)

	//test save git result with empty data
	saveToDatabase(&testingdata2)
	//test save web result with empty data
	saveGitResult(&testingdata2)
}

package hubspot

import (
	"log"
	"os"
	"testing"

	"git.xenonstack.com/util/continuous-security-backend/config"
	"git.xenonstack.com/util/continuous-security-backend/src/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	config.ConfigurationWithToml("../../example.toml")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	os.Remove(os.Getenv("HOME") + "/account-testing.db")
	db, err := gorm.Open("sqlite3", os.Getenv("HOME")+"/account-testing.db")
	if err != nil {
		log.Println(err)
		log.Println("Exit")
		os.Exit(1)
	}
	config.DB = db

	database.CreateDBTablesIfNotExists()

	info := database.RequestInfo{}
	info.Email = "testing@xenonstack.com"
	info.Status = "GitScan"

	db.Create(&info)

	info2 := database.RequestInfo{}
	info2.Email = "testing@xenonstack.com"
	info2.Status = "WebsiteScan"

	db.Create(&info2)

}
func TestSaveRow(t *testing.T) {
	HubspotSubmission("testing@xenonstack.com", "test", "xenonstack.com", "www.xenonstack.com", "GitScan")
	HubspotSubmission("testing@xenonstack.com", "test", "xenonstack.com", "www.xenonstack.com", "WebsiteScan")
}

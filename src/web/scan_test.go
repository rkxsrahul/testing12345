package web

import (
	"testing"

	"git.xenonstack.com/util/continuous-security-backend/src/database"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func TestDelete(t *testing.T) {
	delete("test1254766567978")
}

func TestGitScan(t *testing.T) {

	data := database.RequestInfo{}
	GitScan(data, "javascript")
	GitScan(data, "python")
	GitScan(data, "rust")
	GitScan(data, "golang")
	GitScan(data, "ruby")
}

func TestScan(t *testing.T) {
	data := URL{}
	data.URL = "https://www.xenonstack.com"
	data.Email = "test22@xenonstack.com"
	Scan(data, "", "")
	data = URL{}
	data.URL = "rjhhrb"
	Scan(data, "", "")
}

package web

import (
	"log"
	"testing"

	"git.xenonstack.com/util/continuous-security-backend/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestFetchFromDatabase(t *testing.T) {
	config.ConfigurationWithToml("/home/xs109-rahkum/go/src/git.xenonstack.com/metasecure-ai/web-and-domain-scanning-service/example.toml")
	dbConfig := config.DBConfig()
	// connecting db using connection string
	db, err := gorm.Open("postgres", dbConfig)
	log.Println(err)
	config.DB = db
	defer db.Close()

	chanstream := make(chan interface{})

	// mapd := make(map[string]interface{})

	go FetchFromDatabase("c7funfv7dscce6g0h810", chanstream)
	// c.JSON(200, <-chanstream)
	a := 0
	for {
		if _, ok := <-chanstream; ok {
			log.Println("hello", a)
			a = a + 1
		} else {
			break
		}
	}

	defer ReconnectDatabase()

}

func TestFetchGitResult(t *testing.T) {
	config.ConfigurationWithToml("/home/xs109-rahkum/go/src/git.xenonstack.com/metasecure-ai/web-and-domain-scanning-service/example.toml")
	dbConfig := config.DBConfig()
	// connecting db using connection string
	db, err := gorm.Open("postgres", dbConfig)
	log.Println(err)
	config.DB = db
	defer db.Close()
	chanstream := make(chan interface{})
	go FetchGitResult("c7em5gn7dsc5s6vnckn0", chanstream)
	a := 0
	for {
		if _, ok := <-chanstream; ok {
			log.Println("hello", a)
			a = a + 1
		} else {
			break
		}
	}
	defer ReconnectDatabase()
}

package database

import (
	"testing"

	"git.xenonstack.com/util/continuous-security-backend/config"
)

func TestCreateDBTablesIfNotExists(t *testing.T) {
	CreateDBTablesIfNotExists()
}

func TestCreateDatabase(t *testing.T) {
	config.Conf.Database.Name = "database"
	CreateDatabase()
}

func TestSetup(t *testing.T) {
	Setup()
}

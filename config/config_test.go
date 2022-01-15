package config

import "testing"

func TestConfigurationWithEnv(t *testing.T) {
	ConfigurationWithEnv()
}

func TestConfigurationWithToml(t *testing.T) {
	ConfigurationWithToml("../example.toml")
	ConfigurationWithToml("")
}

func TestSetConfig(t *testing.T) {
	SetConfig()
}

func TestDBConfig(t *testing.T) {
	TomlFile = "../example.toml"
	DBConfig()
	TomlFile = ""
	SetConfig()
}

package nats

import (
	"testing"

	"git.xenonstack.com/util/continuous-security-backend/config"
)

func TestInitConnection(t *testing.T) {
	config.ConfigurationWithToml("../../example.toml")
	InitConnection()
}

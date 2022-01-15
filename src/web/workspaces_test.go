package web

import (
	"testing"

	"git.xenonstack.com/util/continuous-security-backend/config"
)

func TestCheckWorkspaceUser(t *testing.T) {
	config.ConfigurationWithToml("/home/xs109-rahkum/go/src/git.xenonstack.com/metasecure-ai/web-and-domain-scanning-service/example.toml")
	checkWorkspaceUser("rahul@xenonstack.com", "rahul-default-254", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQxODE3ODE0LCJpZCI6MjU0LCJuYW1lIjoiUmFodWwgS3VtYXIiLCJvcmlnX2lhdCI6MTY0MTgxNjAxNCwic3lzX3JvbGUiOiJ1c2VyIn0.AKkQpLaxmReg_nL_uzkXLv22Re1eaRIfUdpa_aDZLfI")
	checkWorkspaceUser("rahul@xenonstack.com", "rahul-default-254", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQxODE3ODE0LCJpZCI6MjU0LCJuYW1lIjoiUmFodWwgS3VtYXIiLCJvcmlnX2lhdCI6MTY0MTgxNjAxNCwic3")

}

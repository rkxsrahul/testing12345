package web

import (
	"git.xenonstack.com/util/continuous-security-backend/config"
	"git.xenonstack.com/util/continuous-security-backend/src/database"
	"git.xenonstack.com/util/continuous-security-backend/src/method"
	"github.com/gosimple/slug"
)

func WorkspaceNameUpdate(email, id string) error {
	db := config.DB
	name := slug.Make(method.ProjectNamebyEmail(email) + "-" + id)
	return db.Model(database.RequestInfo{}).Where("email=?", email).Update("workspace", name).Error
}

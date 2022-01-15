package mail

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func TestEmailTemplate(t *testing.T) {

	mapd := make(map[string]interface{})
	mapd["test"] = "test"
	mapd["xenon"] = "xenon"
	EmailTemplate("/home/xs109-rahkum/go/src/git.xenonstack.com/akirastack/continuous-security-backend/templates/user-notification.tmpl", mapd)
	EmailTemplate("/go/src/git.xenonstack.com/akirastack/continuous-security-backend/templates/user-notification.tmpl", mapd)
}

func TestSendConfirmationMail(t *testing.T) {
	SendConfirmationMail("testing@xenonstack.com", "firstname", "lastname", "www.xenonstack.com")
}

func TestSendMail(t *testing.T) {
	tmplPath, sub, images := ReadToml("userNotification")
	SendMail("testing@xenonstack.com", sub, tmplPath, images)
}

func TestReadToml(t *testing.T) {
	ReadToml("")
	ReadToml("testingTesting")
	ReadToml("testing")
	ReadToml("testingxenon")
}

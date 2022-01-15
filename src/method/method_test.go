package method

import "testing"

func TestHeaderFileSlug(t *testing.T) {
	HeaderFileSlug("test")
}

func TestProjectNamebyEmail(t *testing.T) {
	ProjectNamebyEmail("test@xenonstack.com")
	ProjectNamebyEmail("")
}

func TestCheckURL(t *testing.T) {
	CheckURL("www.xenonstack.com")
	CheckURL("xenonstack")
	CheckURL("https://www.xenonstack.com")
	CheckURL("https://xenonstack")
	CheckURL("www.reoihgfoefje.com")
}

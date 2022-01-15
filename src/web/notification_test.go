package web

import "testing"

func TestUserNotification(t *testing.T) {
	info := URL{}
	UserNotification(info)
}

func TestSupportNotification(t *testing.T) {
	info := URL{}
	SupportNotification(info)

}

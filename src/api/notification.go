package api

import (
	"log"

	"git.xenonstack.com/util/continuous-security-backend/src/web"
	"github.com/gin-gonic/gin"
)

// Notification is an api handler for notificaiton for email send
func Notification(c *gin.Context) {

	var data web.URL
	if err := c.BindJSON(&data); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Please pass url of website",
		})
		return
	}

	web.UserNotification(data)
	web.SupportNotification(data)

	c.JSON(200, gin.H{
		"error":   false,
		"message": "Email has been sent successfully to " + data.Email + "",
	})
}

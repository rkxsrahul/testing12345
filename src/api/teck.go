package api

import (
	"git.xenonstack.com/util/continuous-security-backend/config"
	"git.xenonstack.com/util/continuous-security-backend/src/database"
	"github.com/gin-gonic/gin"
)

// TechStack is an api handler to get results from database related to technology stack
func TechStack(c *gin.Context) {
	uuid := c.Param("id")
	db := config.DB
	scanned := database.ScanResult{}
	db.Where("uuid=? AND command_name=?", uuid, "Technology Stack").Find(&scanned)
	if scanned.CommandName == "" {
		c.JSON(200, gin.H{
			"error": false,
			"stack": scanned,
		})
		return
	}
	c.JSON(200, gin.H{
		"error":   false,
		"stack":   scanned,
		"message": "final result",
	})
}

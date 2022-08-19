package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pi-incident-paging-system/alert"
	"pi-incident-paging-system/webhook_data"
)

func main() {
	r := gin.Default()

	// 1923929
	// 138075000

	r.POST("/opsgenie/alert/created", func(c *gin.Context) {
		var input webhook_data.OpsgenieWebhookCreate

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		alert.SendAlert(input.Alert.Priority, input.Alert.Message, 138075000, []string{
			"1923929",
		})
	})

	r.Run()
}
package main

import (
	"github.com/gin-gonic/gin"
	"pi-incident-paging-system/alert"
)

func main() {
	r := gin.Default()

	// 1923929
	// 138075000

	r.POST("/opsgenie/alert/created", func(c *gin.Context) {
		alert.SendAlert("System 123 Down", 138075000, []string{
			"1923929",
		})
	})

	r.Run()
}
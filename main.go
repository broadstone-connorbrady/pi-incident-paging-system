package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
)

func main() {
	r := gin.Default()

	r.POST("/opsgenie/alert-created", func(c *gin.Context) {
		cmd := exec.Command("sh", "-c", "echo -e \"1923929:Front door buzzer activated\" | ./pocsag -f \"138075000\" -b 3 -r 1200")

		err := cmd.Start()
		fmt.Println("The command is running")
		if err != nil {
			fmt.Println(err)
		}
	})

	// echo -e "1923929:Front door buzzer activated" | sudo ./pocsag -f "138075000" -b 3 -r 1200

	r.Run()
}
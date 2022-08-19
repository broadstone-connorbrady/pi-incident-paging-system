package alert

import (
	"fmt"
	"os/exec"
)

func SendAlert(priority string, message string, frequency int32, addresses []string) {
	for _, address := range addresses{
		command := fmt.Sprintf("echo -e \"%s:%s - %s\" | ./pocsag -f \"%d\" -b 3 -r 1200", address, priority, message, frequency)

		systemCommand := exec.Command("bash", "-c", command)

		go systemCommand.Run()
	}
}
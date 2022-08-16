package alert

import (
	"fmt"
	"os/exec"
)

func SendAlert(message string, frequency int32, addresses []string) {
	for _, address := range addresses{
		command := fmt.Sprintf("echo -e \"%s:%s\" | ./pocsag -f \"%d\" -b 3 -r 1200", address, message, frequency)

		systemCommand := exec.Command("bash", "-c", command)

		systemCommand.Run()
	}
}
package main

import (
	"./state"
	"os"
	"strings"
)

func main() {
	function := os.Args[1]
	serviceStateManager := state.New()

	if strings.EqualFold(function, "add") {
		serviceName := os.Args[2]
		command := os.Args[3]
		serviceStateManager.AddService(serviceName, command)
	} else if strings.EqualFold(function, "start") {
		serviceName := os.Args[2]
		command := serviceStateManager.GetService(serviceName)
		state.Start(command)
	} else if strings.EqualFold(function, "list") {
		serviceStateManager.List()
	} else if strings.EqualFold(function, "delete") {
		serviceStateManager.Remove(os.Args[2])
	} else if strings.EqualFold(function, "delete-all") {
		serviceStateManager.RemoveAll()
	}



	serviceStateManager.Close()

}

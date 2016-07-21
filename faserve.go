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
		serviceStateManager.AddService(os.Args[2], os.Args[3])
	} else if strings.EqualFold(function, "start") {
		command := serviceStateManager.GetService(os.Args[2])
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

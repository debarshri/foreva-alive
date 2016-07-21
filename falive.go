package main

import (
	"./state"
	"os"
)

func main() {
	state.Start(os.Args[1])
}

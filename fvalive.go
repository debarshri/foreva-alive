package main

import (
	"os"
	"os/exec"
	"bufio"
	"fmt"
)

func main() {
	start(os.Args[1])
}

func start(cmd_line string) {
	cmd := exec.Command("sh", "-c", cmd_line)
	stdout, _ := cmd.StdoutPipe()
	in := bufio.NewScanner(stdout)

	err_on_run := cmd.Run()

	for in.Scan() {
		fmt.Println(in.Text())
	}

	if err_on_run != nil{
		start(cmd_line)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	start(os.Args[1])
}

func start(cmd_line string) {

	fmt.Println("Starting app..")

	cmd := exec.Command("sh", "-c", cmd_line)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	in := bufio.NewScanner(stdout)
	err_stream := bufio.NewScanner(stderr)

	go collect(in)
	go collect(err_stream)

	err_on_run := cmd.Run()

	if err_on_run != nil {
		start(cmd_line)
	}
}

func collect(in *bufio.Scanner) {
	for in.Scan() {
		fmt.Println(in.Text())
	}
}

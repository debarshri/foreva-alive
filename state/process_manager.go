package state

import (
	"bufio"
	"log"
	"os/exec"
)

func Start(cmd_line string) {

	cmd, in, err_stream := prepare_command(cmd_line)

	log.Println("[FASERVE]  Starting channels")

	std, err_std := create_channels()

	go collect(in, std)
	go collect(err_stream, err_std)

	log.Println("[FASERVE] Running command")

	err_on_run := cmd.Run()

	log.Println(<-std)
	log.Println(<-err_std)

	if err_on_run != nil {
		Start(cmd_line)
	}
}

func create_channels() (std chan string, err_std chan string) {

	std = make(chan string)
	err_std = make(chan string)

	return std, err_std

}
func prepare_command(cmd_line string) (cmd *exec.Cmd, in *bufio.Scanner, err_stream *bufio.Scanner) {

	cmd = exec.Command("sh", "-c", cmd_line)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	in = bufio.NewScanner(stdout)
	err_stream = bufio.NewScanner(stderr)

	return cmd, in, err_stream

}
func collect(in *bufio.Scanner, out chan string) {
	for in.Scan() {
		out <- in.Text()
	}

	close(out)
}

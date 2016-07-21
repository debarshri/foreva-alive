package main

import "fmt"

func main() {

	c := make(chan string)
	go run("test", c)

	x := <-c // receive from c
	fmt.Println(x)

}

func run(data string, c chan string) {

	fmt.Println(data)

	data = "test2"
	c <- data
}

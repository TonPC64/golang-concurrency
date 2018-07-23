package main

import (
	"fmt"
	"time"
)

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func pinger(c chan<- string) {
	for {
		c <- "ping"
	}
}

func ponger(c chan<- string) {
	for {
		c <- "pong"
	}
}

func main() {
	c := make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	fmt.Scanln()
}

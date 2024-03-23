package main

import (
	"fmt"
	"time"
)

func workerOne(c1 chan<- string) {
	const delay = 500 * time.Millisecond
	for {
		c1 <- fmt.Sprintf("worker one: every %s", delay)
		time.Sleep(delay)
	}
}

func workerTwo(c2 chan<- string) {
	const delay = 2 * time.Second
	for {
		c2 <- fmt.Sprintf("worker two: every %s", delay)
		time.Sleep(delay)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go workerOne(c1)
	go workerTwo(c2)

	ft := time.Now().Add(time.Second * 10)

	for time.Until(ft) >= 0 {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		default:
		}
	}
}

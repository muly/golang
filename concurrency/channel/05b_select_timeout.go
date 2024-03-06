// using select on group of channels

package main

import (
	"fmt"
	"time"
)

var resp chan string

func main() {
	resp = make(chan string) // Note: without this initialiazation the select will only reach the timeout case.

	go send()
	wait()
}

func wait() {
	tick := time.NewTicker(5 * time.Second)

	for {
		select {
		case x := <-resp:
			fmt.Println(x)
			return
		case <-tick.C:
			fmt.Println("timeout")
			return
		}
	}
}

func send() {
	time.Sleep(1 * time.Second) // increase this > timeout value to simulate timeout

	resp <- "hello"
}

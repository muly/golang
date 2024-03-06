package main

import (
	"fmt"
	"time"
)

// sync.Mutex

// var wg sync.WaitGroup

func main() {
	var s chan string
	go someComplexFunc1(s)
	time.Sleep(20 * time.Second) // rudimentary way to wait for the go routine to finish
	fmt.Println(<-s)
}

func someComplexFunc1(s chan string) {
	time.Sleep(10 * time.Second)
	fmt.Println("hello world")
	// lock s
	s <- "hello"
	// unlock s
}

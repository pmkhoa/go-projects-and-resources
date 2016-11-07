package main

import (
	"fmt"
	"time"
)

func sayHello(c chan string, counter int) {
	for i := 0; ; i++ {
		c <- fmt.Sprint("hello", counter)
	}
}

func printHello(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// Declare channel of string
	var c chan string = make(chan string)
	go sayHello(c, 0)
	go sayHello(c, 1)
	go sayHello(c, 2)
	go printHello(c)
	go printHello(c)

	var input string
	fmt.Scanln(&input)
}

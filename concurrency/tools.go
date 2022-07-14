package main

import (
	"fmt"
	"time"
)

func simpleGoRoutine() {
	go helloWorld()
	time.Sleep(1 * time.Millisecond) // If we don't wait at least 1 ml second the main routine finishes before helloWorld routine

	bufferedChannelsCorrectUseExample()
}

func channelsBlockedExample() {
	/*
		Channels
		- are and should be the only way to communicate between go routines
		- They have mainly two responsibilities: communicate different go routines and synchronize them
	*/
	c := make(chan string)
	fmt.Println("sending to the channel")

	c <- "hello" // This blocks the execution. In order to send something to a channel we need a different go routing listening to that channel

	fmt.Println("receiving from the channel")

	greeting := <-c
	fmt.Println("greeting received")

	fmt.Println(greeting)
}

func channelsCorrectUseExample() {
	c := make(chan string)
	fmt.Println("sending to the channel")

	go func(input chan string) {
		input <- "hello"
	}(c)

	fmt.Println("receiving from the channel")

	greeting := <-c // Here the main go routing listens to the channel c

	fmt.Println("greeting received")

	fmt.Println(greeting)
}

func bufferedChannelsCorrectUseExample() {
	/*
		Buffered channels are channels with multiple slots
		These channels block when its capacity is full
	*/
	c := make(chan string, 3)

	go func(input chan string) {
		fmt.Println("sending 1 to the channel")
		input <- "hello1"

		fmt.Println("sending 2 to the channel")
		input <- "hello2"

		fmt.Println("sending 3 to the channel")
		input <- "hello3"
	}(c)

	fmt.Println("receiving from the channel")

	for greeting := range c {
		fmt.Println("greeting received")

		fmt.Println(greeting)
	}
}

func helloWorld() {
	fmt.Println("Hello World")
}

package main

import (
	"fmt"
	"time"
)

func main() {


c:= make(chan string)

go Ponger(c)
go Pinger(c)
go printer(c)

	var input string
	fmt.Scanln(&input)
}

func Pinger(c chan<- string)  {
	for i := 0; i<5; i++{

		c <- "ping"   // this function sends ping on the channel
	}
}

func Ponger(c chan<- string)  {
	for i := 0; i<5; i++{

		c <- "pong"   // this function sends ping on the channel
	}
}

func printer(c <-chan string) {
	for  {
		msg := <- c // recieces message from the channel
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
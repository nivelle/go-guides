package main

import (
	"fmt"
)

func main(){
	message := make(chan string)
	signals := make(chan bool)

	select {
	case msg:=<-message:
		fmt.Println("receive message:",msg)
	default:
		fmt.Println("no message received")
	}

	msg:="hi"
	select {
	case message<-msg:
		fmt.Println("sent message:",msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg:=<-message:
		fmt.Println("received messag:e",msg)
	case sig:=<-signals:
		fmt.Println("received signal:",sig)
	default:
		fmt.Println("no activity")
	}

}
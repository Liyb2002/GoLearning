package main

import "fmt"

func main(){
	msg:= make(chan string)

	go func(){
		msg <- "ping"
	}()

	message:= <- msg
	fmt.Println(message)
}
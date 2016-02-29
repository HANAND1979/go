package main

import (
	"fmt"
	//"time"
)

func main() {
	messages := make(chan string, 5)
	messages <- "alpha"
	messages <- "beta"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	done_words := make(chan bool, 1)
	go words(done_words)
	done_numbers := make(chan bool, 1)
	go numbers(done_numbers)
	<-done_words
	<-done_numbers
}

func words(done chan bool){
	list:=[]string{"a","b","c","d","e"}
	fmt.Println(list)
	done<-true

}
func numbers(done chan bool){
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	done<-true
}

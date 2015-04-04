package main

import "fmt"
import "time"
import "math/rand"

func boring(msg string) <-chan string { //return receive-only channel of string
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c:= make(chan string)
	go func() { for { c <- <- input1 } }()
	go func() { for { c <- <- input2 } }()
	return c
}

func main() {
	ann := boring("Ann: ")
	jon := boring("Jon: ")
	//c := fanIn(boring("Ann: "), boring("Jon: "))
	c := fanIn(ann, jon)
	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
}
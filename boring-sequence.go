// https://www.youtube.com/watch?v=f6kdp27TYZs 20.12

package main

import "fmt"
import "time"
import "math/rand"

type Message struct {
	str string
	wait chan bool
}

func boring(msg string) <-chan Message { //return receive-only channel of string
	waitForIt := make(chan bool)
	c := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			c <- Message{ fmt.Sprintf("%s %d", msg, i), waitForIt }
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<- waitForIt
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c:= make(chan Message)
	go func() { for { c <- <- input1 } }()
	go func() { for { c <- <- input2 } }()
	return c
}

func main() {	
	c := fanIn(boring("Ann: "), boring("Jon: "))
	for i := 0; i < 20; i++ {
		msg1 := <- c; fmt.Println(msg1.str)	
		msg2 := <- c; fmt.Println(msg2.str)	
		msg1.wait <- true	
		msg2.wait <- true
	}
}
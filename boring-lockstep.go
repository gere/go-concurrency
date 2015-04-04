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

func main() {
	ann := boring("Ann: ")
	jon := boring("Jon: ")
	for i := 0; i < 10; i++ {
		fmt.Println(<-jon)
		fmt.Println(<-ann)
	}
}
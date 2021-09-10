package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	errChan := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		c1 <- "one"
		errChan <- "err one"
	}()
	go func() {
		time.Sleep(2 * time.Second)

		c2 <- "two"
		errChan <- "err two"

	}()
	go func() {
		time.Sleep(1 * time.Second)
		c3 <- "three"
		// errChan <- "err three"
	}()

	count := 0
	errorX := ""
	for count < 3 && errorX == "" {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
			count++
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			count++
		case msg3 := <-c3:
			fmt.Println("received", msg3)
			count++
		case err := <-errChan:
			fmt.Println("received", err)
			errorX = err
		}
	}
}

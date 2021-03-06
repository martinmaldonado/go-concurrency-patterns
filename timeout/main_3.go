package main

import (
	"fmt"
	"time"
	"math/rand"
) 

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

	g := GetData("fravega")

	FOR:
	for {
		// Wait for a communication
		select {

		case <- time.After(1000 * time.Millisecond):

			fmt.Println("Timeout")
			break FOR

		case msg := <- g:

			fmt.Println(msg)
		}
	}
}

// Get data returns the channel
func GetData(user string) chan string {
	ch := make(chan string)

	go func() {

		i := 0
		for {

			time.Sleep(time.Duration(750 + rand.Intn(500)) * time.Millisecond)
			ch <- fmt.Sprintf("%s is done (%d)", user, i)
			i++
		}
	} ()

	return ch
}
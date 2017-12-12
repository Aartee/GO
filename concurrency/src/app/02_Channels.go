package main

import "fmt"

func sum(s []int, channel chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	channel <- sum //send sum to channel
}

func main() {
	//Create a new channel with make(chan val-type)
	messages := make(chan string)

	//Send a value into a channel using the channel <- syntax.
	//Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	//The <-channel syntax receives a value from the channel.
	//Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)

	//The example code sums the numbers in a slice, distributing the work between two goroutines.
	//Once both goroutines have completed their computation, it calculates the final result.
	s := []int{7, 2, 8, -9, 4, 0}

	c = make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c //receive from c

	fmt.Println(x, y, x+y)

}

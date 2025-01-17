package main

import "fmt"

func main() {
	//Here we make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	//Because this channel is buffered, we can send these values into
	//the channel without a corresponding concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	//Later we can recieve these two values as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

package channelbuffer

import "fmt"

func channelWithoutBuffer() {
	channel := make(chan string)
	channel <- "first"
	fmt.Println(<-channel)
}

func channelWithBuffer() {
	channel := make(chan string, 1) // 1 adalah buffer
	channel <- "first"
	fmt.Println(<-channel)
}

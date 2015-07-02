package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func whisper(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: whisper <number-of-whispers>")
		return
	}

	count, _ := strconv.Atoi(os.Args[1])

	start := now()
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i < count; i++ {
		right = make(chan int)
		go whisper(left, right)
		left = right
	}

	right <- 1
	res := <-leftmost
	fmt.Println("res:", res)

	stop := now()
	duration := (stop - start)

	fmt.Printf("took %f seconds\n", float64(duration)/1000)
}

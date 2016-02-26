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

func worker(count int, left, right chan int) {
	for i := 0; i < count; i++ {
		left <- 1 + <-right
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: ring <number-of-workers> <messages-count>")
		return
	}

	count, _ := strconv.Atoi(os.Args[1])
	msgCount, _ := strconv.Atoi(os.Args[2])

	for run := 0; run < 30; run++ {
		start := now()
		leftmost := make(chan int, 10)
		right := leftmost
		left := leftmost

		for i := 0; i < count; i++ {
			right = make(chan int, 10)
			go worker(msgCount, left, right)
			left = right
		}

		for n := 0; n < msgCount; n++ {
			right <- n
			<-leftmost
		}

		stop := now()
		duration := (stop - start)

		fmt.Printf("run %d: took %f seconds\n", run, float64(duration)/1000)
	}
}

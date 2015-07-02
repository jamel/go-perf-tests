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

func worker(in chan int, out chan int) {
	for {
		val := <-in
		out <- val
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: chan <millions-of-roundtrips>")
		return
	}

	count, _ := strconv.Atoi(os.Args[1])
	count *= 1000000

	out := make(chan int, 0)
	in := make(chan int, 0)

	start := now()
	go worker(out, in)

	val := 0
	for i := 0; i < count; i++ {
		out <- val
		val = <-in
	}

	stop := now()
	duration := (stop - start)
	ns := (duration * 1000000) / int64(count)

	fmt.Printf("done %dM roundtrips in %f seconds\n",
		count/1000000, float64(duration)/1000)
	fmt.Printf("duration of passing a single message: %d ns\n", ns)
	fmt.Printf("message passes per second: %fM\n",
		float64(1000000000/ns)/1000000)
}

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

func worker() {
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go <millions-of-goroutines>")
		return
	}

	count, _ := strconv.Atoi(os.Args[1])
	count *= 1000000
	start := now()

	for i := 0; i < count; i++ {
		go worker()
	}

	stop := now()
	duration := (stop - start)
	ns := (duration * 1000000) / int64(count)

	fmt.Printf("executed %dM goroutines in %f seconds\n",
		count/1000000, float64(duration)/1000)
	fmt.Printf("duration of one coroutine creation+termination: %d ns\n", ns)
	fmt.Printf("goroutine creations+terminatios per second: %fM\n",
		float64(1000000000/ns)/1000000)
}

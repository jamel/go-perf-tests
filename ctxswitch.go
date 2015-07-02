package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func worker(count int) {
	for i := 0; i < count; i++ {
		runtime.Gosched()
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ctxswitch <millions-of-constext-switches>")
		return
	}

	count, _ := strconv.Atoi(os.Args[1])
	count = count * 1000000 / 2

	start := now()
	go worker(count)

	for i := 0; i < count; i++ {
		runtime.Gosched()
	}

	stop := now()
	duration := (stop - start)
	ns := (duration * 1000000) / (int64(count) * 2)

	fmt.Printf("performed %dM context switches in %f seconds\n",
		count*2/1000000, float64(duration)/1000)
	fmt.Printf("duration of one context switch: %d ns\n", ns)
	fmt.Printf("context switches per second: %fM\n",
		float64(1000000000/ns)/1000000)
}

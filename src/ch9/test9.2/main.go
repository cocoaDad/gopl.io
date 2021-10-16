package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var count = 0

// var open = make(chan struct{})
var A2B, B2A = make(chan struct{}), make(chan struct{})
var mu = sync.RWMutex{}

func trans(c chan struct{}) {
	for {
		<-c
		mu.Lock()
		count++
		c <- struct{}{}
		mu.Unlock()
	}
}

func output() int {
	mu.RLock()
	defer mu.RUnlock()
	return count
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	if err := input.Err(); err != nil {
		fmt.Errorf("%v\n", err)
	}
	if input.Text() == "open" {
		go trans(A2B)
		go trans(B2A)
	}
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	tick := time.NewTicker(time.Second)
	for {
		select {
		// case <- open:
		case <-tick.C:
			fmt.Fprintf(os.Stdout, "After 1s: %d\n", output)
		case <-abort:
			fmt.Println("Launch aborted!")
			tick.Stop()
			close(B2A) // close
			close(A2B)
			return
		}
	}
}

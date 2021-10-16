package main

import (
	"bufio"
	"fmt"
	"os"
)

func test1_1() {
	fmt.Println(os.Args[0])
}
func test1_2() {
	for s, sep := range os.Args {
		fmt.Println(s, sep)
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		switch input.Text() {
		case "1.1":
			test1_1()

		case "1.2":
			test1_2()
		case "break":
			return
		}
	}
}

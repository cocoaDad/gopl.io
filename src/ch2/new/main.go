package main

import "fmt"

func main() {
	new1 := new(struct{})
	new2 := new(struct{})

	new3 := new([0]int)
	new4 := new([0]int)
	fmt.Println(&new1, &new2)
	fmt.Println(&new3, &new4)
}

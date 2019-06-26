package main

import "fmt"

func main() {
	var test [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		test[i] = 100 + i
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("test[%d]=%d \n", j, test[j])
	}
}

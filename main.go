package main

import "fmt"

func main() {
	fmt.Println("Hello this is feature1")
	fmt.Printf("Call func Min(2, 3): %d\n", Min(2, 3))
	fmt.Println("End")
}

func Min(a, b int) int{
	if a < b {
		return a
	}
	return b
}

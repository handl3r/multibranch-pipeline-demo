package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

func Min(a, b int) int{
	if a < b {
		return a
	}
	return b
}

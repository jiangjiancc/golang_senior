package main

import "fmt"

func fibonacci(i int) int  {
	if i < 2 {
		return i
	}
	return fibonacci(i-2) + fibonacci(i-1)
}
func main() {
	fmt.Println(fibonacci(34))
}

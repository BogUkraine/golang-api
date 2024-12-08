package main

import "fmt"

func main() {
	var x int = 10
	var y *int = &x
	var k int = *y

	fmt.Println(y, x, k)
}

package main

import "fmt"

func main() {
	var a [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		a[i] = i + 10
	}
	for j = 0; j < 10; j++ {
		fmt.Println(a[j])
	}

}

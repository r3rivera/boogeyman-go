package main

import "fmt"

func main() {

	items := []int{1, 2, 3, 4}

	for _, v := range items {
		fmt.Println(v)
	}
}

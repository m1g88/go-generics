package main

import "fmt"

func main() {
	fmt.Println(Sum([]int64{1, 2, 3, 4}...))

	fmt.Println(Sum(1, 2, 3, 4))
}

func Sum(n ...int64) int64 {
	var s int64

	for _, v := range n {
		s += v
	}
	return s
}

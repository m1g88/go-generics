package main

import "fmt"

func main() {
	fmt.Println(Sum([]int64{1, 2, 3, 4}...))
	fmt.Println(Sum(int64(1), int64(2), int64(3), int64(4)))
	fmt.Println(Sum(1.2, 2.2, 3.3, 4.4))
}

func Sum[V int64 | float64](n ...V) V {
	var s V

	for _, v := range n {
		s += v
	}
	return s
}

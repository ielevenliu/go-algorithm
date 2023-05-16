package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	s[0] = 1
	s[1] = 2

	ss := s[1:2]
	ss = append(ss, 4, 5)
	fmt.Printf("slice: %+v, len: %d, cap: %d \n", s, len(s), cap(s))
	fmt.Printf("slice: %+v, len: %d, cap: %d \n", ss, len(ss), cap(ss))
}

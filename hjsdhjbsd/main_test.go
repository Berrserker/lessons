package main

import (
	`testing`
)

func TestScale(t *testing.T) {
	const  = iota
	c1 := make(chan int)
	c2 := make(chan int)
	a1 := []int{1, 5, 7, 8, 24, 25, 50}
	a2 := []int{1, 3, 6, 8, 25}
	
	go func() {dump(c1, a1)}()
	go dump(c2, a2)
	
	merge(c1, c2)
}

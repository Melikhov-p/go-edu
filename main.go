package main

import (
	"first-go/mathslice"
	"fmt"
)

func main() {
	slicesum := mathslice.SumSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	fmt.Println(slicesum)
}

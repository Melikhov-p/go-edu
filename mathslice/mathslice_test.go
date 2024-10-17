package mathslice

import (
	"fmt"
	"testing"
)

func TestSumSlice(t *testing.T) {
	_, err := SumSlice([]int{1, 2, 3})

	if err != nil {
		t.Error(err)
	}
}

func TestSumEmptySlice(t *testing.T) {
	_, err := SumSlice([]int{})

	if err != nil {
		t.Error(err)
	}
}

func TestMapSlice(t *testing.T) {
	slc := MapSlice([]int{1, 2, 3}, func(i int) int {
		return i * i
	})
	fmt.Println(slc)
}

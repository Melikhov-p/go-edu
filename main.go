package main

import "fmt"

type Myinter interface {
}

func Mul(a Myinter, b int) {
	switch a.(type) {
	case int:
		fmt.Println(a.(int) * b)
	}
}

func main() {
	Mul(1, 2)
}

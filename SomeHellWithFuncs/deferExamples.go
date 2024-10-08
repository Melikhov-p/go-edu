package SomeHellWithFuncs

import "fmt"

var Global = 5

func testFunc() {
	defer func(check int) {
		Global = check
	}(Global) // defer RollBack(Global)
	Global = 12
	fmt.Println("New Global: ", Global)
}

func RollBack(x int) {
	Global = x
}

func main() {
	testFunc()
	fmt.Println("Old Global: ", Global)
}

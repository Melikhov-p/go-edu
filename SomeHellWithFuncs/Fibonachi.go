package SomeHellWithFuncs

func main() {
	var fiba func() = Fib()

	fiba()
	fiba()
	fiba()
	fiba()
	fiba()
	fiba()
	fiba()
	fiba()
}

func Fib() func() {
	a, b := 0, 1
	return func() {
		println(a, b)
		a, b = b, a+b
	}
}

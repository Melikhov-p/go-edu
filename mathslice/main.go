package mathslice

func SumSlice(s []int) int {
	sum := 0

	for _, v := range s {
		sum += v
	}

	return sum
}

func MapSlice(s []int, op func(int) int) {
	for i := 0; i < len(s); i++ {
		s[i] = op(i)
	}
}

package mathslice

func SumSlice(s []int) (int, error) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	return sum, nil
}

func MapSlice(s []int, op func(int) int) []int {
	for i := 0; i < len(s); i++ {
		s[i] = op(s[i])
	}
	return s
}

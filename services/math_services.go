package services

func Add(a, b int) int {

	return a + b
}

func AddMany(numbers ...int) int {

	res := 0

	for _, v := range numbers {
		res += v
	}

	return res
}

func multiply(a, b int) int {
	return a * b
}

func Multiply(a, b int) int {
	return a * b
}

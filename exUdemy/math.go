package ex

func IsOdd(x int) bool {
	return x&1 != 0
}

func IsEven(x int) bool {
	return !(IsOdd(x))
}

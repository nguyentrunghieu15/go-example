package eximplement

import (
	"testing"
)

func TestSum(t *testing.T) {
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	go sum(c, arr[:5]...)
	go sum(c, arr[5:]...)
	sum := 0
	for i := 0; i < 2; i++ {
		sum += <-c
	}
	if sum != 55 {
		t.Error("Error func Sum")
	}
}

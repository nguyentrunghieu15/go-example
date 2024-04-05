package ex

import (
	"fmt"
	"testing"
)

func TestIsEven(t *testing.T) {
	x := 4
	if !IsEven(x) {
		t.Errorf("Test even error , %v is even but return false", x)
	}
	y := 5
	if IsEven(y) {
		t.Errorf("Test even error , %v is odd but return true", y)
	}
}

func TestIsOdd(t *testing.T) {
	x := 4
	if IsOdd(x) {
		t.Errorf("Test func odd error , %v is even but return true", x)
	}
	y := 5
	if !IsOdd(y) {
		t.Errorf("Test func odd error , %v is odd but return false", y)
	}
}

func TestIsEvenAndIsOdd(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range arr {
		if IsEven(v) {
			fmt.Println(v, "is Even")
		}
		if IsOdd(v) {
			fmt.Println(v, "is Odd")
		}
	}
}

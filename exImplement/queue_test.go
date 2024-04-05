package eximplement

import (
	"testing"
)

func TestPush(t *testing.T) {
	var temp = Queue[int]{}
	arr := []int{1, 2, 3, 4}
	temp.Push(arr...)
	if temp.Size() != 4 {
		t.Error("Error func Push")
	}
	for i, v := range arr {
		if val, e := temp.At(i); e != nil || val != v {
			t.Error("Error func Push")
		}
	}
	temp.Show()
}

func TestPop(t *testing.T) {
	var temp = Queue[int]{}
	e := temp.Pop()
	if e == nil {
		t.Error("Error func Pop")
	}
	arr := []int{1, 2, 3, 4}
	temp.Push(arr...)
	e = temp.Pop()
	if e != nil {
		t.Error("Error func Pop")
	}
	if temp.Size() != 3 {
		t.Error("Error func Pop")
	}
	for i, v := range arr[1:] {
		if val, e := temp.At(i); e != nil || val != v {
			t.Error("Error func Pop")
		}
	}
	temp.Show()
}

func TestFront(t *testing.T) {
	var temp = Queue[int]{}
	var front, e = temp.Front()
	if e == nil {
		t.Error("Error func Front")
	}
	arr := []int{1, 2, 3, 4}
	temp.Push(arr...)
	front, e = temp.Front()
	if e != nil || front != 1 {
		t.Error("Error func Front")
	}
	temp.Show()
}

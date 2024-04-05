// Using For-loop, interfaces, data structure, error handling,name conversation, data structure , go module, pakage import

package eximplement

import (
	"fmt"
	"time"
)

type List interface {
	Show()
	At(index int) (any, error)
	Size() int
	ForEach(fn func(v any))
}

type QueueError struct {
	time string
	log  string
}

const (
	QueueErrorEmptyLog   = "This queue empty"
	QueueErrorOutOfRange = "Index out of range"
)

func (v QueueError) Error() string {
	return fmt.Sprintf("[%v] Error: %v", v.time, v.log)
}

type Queue[V any] struct {
	items []V
	size  int
}

func (v *Queue[V]) Push(val ...V) {
	v.items = append(v.items, val...)
	v.size += len(val)
}

func (v *Queue[V]) Pop() error {
	if v.size == 0 {
		return QueueError{time.Now().Format(time.UnixDate), QueueErrorEmptyLog}
	}
	v.items = v.items[1:]
	v.size--
	return nil
}

func (v *Queue[V]) Front() (any, error) {
	if v.size == 0 {
		return nil, QueueError{time.Now().Format(time.UnixDate), QueueErrorEmptyLog}
	}
	return v.items[0], nil
}

func (v *Queue[V]) Show() {
	fmt.Println(v.items)
}

func (v *Queue[V]) At(index int) (any, error) {
	if index < 0 || index >= v.size {
		return nil, QueueError{time.Now().Format(time.UnixDate), QueueErrorOutOfRange}
	}
	if v.size == 0 {
		return nil, QueueError{time.Now().Format(time.UnixDate), QueueErrorEmptyLog}
	}
	return v.items[index], nil
}

func (v *Queue[V]) Size() int {
	return v.size
}

func (v *Queue[V]) ForEach(fn func(v any)) {
	for i := 0; i < v.size; i++ {
		fn(v.items[i])
	}
}

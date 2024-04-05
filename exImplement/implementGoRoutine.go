// Using mutithread, generic, interface

package eximplement

type Addable interface {
	~int | ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64 | ~float32 | ~float64
}

func sum[V Addable](c chan V, arr ...V) {
	var s V
	for _, v := range arr {
		s += v
	}
	c <- s
}

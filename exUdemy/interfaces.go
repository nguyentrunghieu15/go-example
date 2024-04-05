package ex

import "fmt"

type Triangle struct {
	base   float64
	height float64
}

type Square struct {
	sideLength float64
}

type Shape interface {
	getArea() (float64, error)
}

type ShapeUndefineError struct{}

func (u *ShapeUndefineError) Error() string {
	return "The shape undefined"
}

func (u *Triangle) getArea() (float64, error) {
	if u == nil {
		return 0, &ShapeUndefineError{}
	}
	return u.base * u.height * 0.5, nil
}

func (u *Square) getArea() (float64, error) {
	if u == nil {
		return 0, &ShapeUndefineError{}
	}
	return u.sideLength * u.sideLength, nil
}

func PrintArea(u Shape) {
	v, e := u.getArea()
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("Area is: ", v)
}

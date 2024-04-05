package ex

import (
	"testing"
)

func TestPrintArea(t *testing.T) {
	a := Square{2}
	b := Triangle{2, 3}
	var c *Square
	PrintArea(c)
	PrintArea(&b)
	if v, _ := a.getArea(); v != 4 {
		t.Errorf("Error method getArea, Area of %T %v is %v", a, a, 4)
	}
	if v, _ := b.getArea(); v != 3 {
		t.Errorf("Error method getArea, Area of %T %v is %v", b, b, 3)
	}
}

package mathy

import (
	"fmt"
	"strings"
	"testing"
)

func TestClipInt(t *testing.T) {
	x := ClipInt(-5, -2, 2)
	if x != -2 {
		t.Errorf("Expected -2, got %v", x)
	}

	x = ClipInt(5, -2, 2)
	if x != 2 {
		t.Errorf("Expected 2, got %v", x)
	}

	x = ClipInt(1, -2, 2)
	if x != 1 {
		t.Errorf("Expected 1, got %v", x)
	}
}

func TestPolyReduce(t *testing.T) {
	a := []int{67}
	x := 17

	c := PolyReduce(a, x)
	expected := []int{16, 3}
	for i, v := range c {
		if v != expected[i] {
			err := fmt.Sprint("Expected ", expected[i], " at index ", i)
			t.Errorf(err)
		}
	}
}

func TestInt2Poly(t *testing.T) {
	a := 67
	x := 17

	c := Int2Poly(a, x)
	expected := []int{16, 3}
	for i, v := range c {
		if v != expected[i] {
			err := fmt.Sprint("Expected ", expected[i], " at index ", i)
			t.Errorf(err)
		}
	}

	i := Poly2Int(c, x)
	if i != 67 {
		t.Errorf("Expected 67")
	}
}

func TestPolyReduce2(t *testing.T) {
	a := []int{67, 0, 0, 0, 0}
	x := 17

	c := PolyReduce(a, x)
	expected := []int{16, 3}
	for i, v := range c {
		if v != expected[i] {
			err := fmt.Sprint("Expected ", expected[i], " at index ", i)
			t.Errorf(err)
		}
	}
}

func TestPolyAdd(t *testing.T) {
	// 67 in base 17
	a := []int{16, 3}
	x := 17

	c := PolyReduce(PolyAdd(a, a), x)

	expected := []int{15, 7}
	for i, v := range c {
		if v != expected[i] {
			err := fmt.Sprint("Expected ", expected[i], " at index ", i)
			t.Errorf(err)
		}
	}
}

func TestPolyMul(t *testing.T) {
	a := []int{9, 3} // this is 39 in base 10
	x := 10

	c := PolyReduce(PolyMul(a, a), x)
	expected := []int{1, 2, 5, 1}
	for i, v := range c {
		if v != expected[i] {
			err := fmt.Sprint("Expected ", expected[i], " at index ", i)
			t.Errorf(err)
		}
	}
	//
	//for _, v := range c {
	//	println(v)
	//}
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
func TestPolyDivide(t *testing.T) {
	x := 10456729
	c := Int2Poly(x, 7)
	fmt.Println("[", "-", "]: ", arrayToString(c, ", "))

	d := 5
	for x > 0 {
		x = x / d
		ci := Int2Poly(x/d, 7)
		fmt.Println("[", d, "]: ", arrayToString(ci, ", "))

	}

}

func TestChangeBase(t *testing.T) {
	x := 15
	x_2 := Int2Poly(x, 2)
	x_10 := ChangeBase(x_2, 2, 10)
	y := Poly2Int(x_10, 10)
	fmt.Println(y)
}

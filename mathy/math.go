package mathy

func MaxInt(nums ...int) int {
	maxNum := nums[0]
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}
func MinInt(nums ...int) int {
	minNum := nums[0]
	for _, v := range nums {
		if v < minNum {
			minNum = v
		}
	}
	return minNum
}

func AbsInt(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func SumIntSlice(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

func MultiplyIntSlice(nums []int) int {
	product := 1
	for _, n := range nums {
		product *= n
	}
	return product
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func ClipInt(num int, lo int, hi int) int {
	retVal := Min(Max(num, lo), hi)
	return retVal
}

func Int2Poly(a int, x int) []int {
	// convert a to a base x polynomial
	retVal := []int{}

	for a > 0 {
		m, n := a%x, a/x
		retVal = append(retVal, m)
		a = n
	}

	return retVal
}

func Poly2Int(c []int, x int) int {
	retVal := 0
	x_term := 1
	for _, ci := range c {
		retVal += ci * x_term
		x_term *= x
	}
	return retVal
}

func PolyAdd(a []int, b []int) []int {
	// a+b
	size := Max(len(a), len(b))

	c := make([]int, size)
	for i := 0; i < size; i++ {
		ai := 0
		bi := 0
		if i < len(a) {
			ai = a[i]
		}
		if i < len(b) {
			bi = b[i]
		}
		c[i] = ai + bi
	}

	return c

}

func PolyMul(a []int, b []int) []int {

	if len(a) == 0 || len(b) == 0 {
		return []int{}
	}
	// multiply two polynomials
	c := make([]int, len(a)+len(b))

	for ai := 0; ai < len(a); ai++ {
		for bi := 0; bi < len(b); bi++ {
			ci := ai + bi
			c[ci] += a[ai] * b[bi]
		}
	}
	return c
}

func PolyReduce(a []int, x int) []int {
	// reduces a polynomial whose base is 'x'
	// a[0] = a[0]*x^0
	// a[1] = a[1]*x^1
	// a[2] = a[2]*x^2
	// ...
	// a[n] = a[n]*x^n

	// no digit can be more than the base (x)
	// if a[n] > x
	// then a[n] = a[n] - x
	//      a[n+1] += 1

	c := a

	for n := 0; n < len(c); n++ {
		for c[n] > x {
			cdx := c[n] / x
			cmx := c[n] % x
			c[n] = cmx
			if n == len(c)-1 {
				// last coeff, add a new one
				c = append(c, cdx)
			} else {
				c[n+1] += cdx
			}
		}
	}

	if len(c) >= 1 {
		// trim leading zeros
		for c[len(c)-1] == 0 {
			c = c[:len(c)-1]
		}
	}

	for _, x := range c {
		if x < 0 {
			panic("did you overflow?")
		}
	}

	return c

}

func PolyPow(a []int, n int, base int) []int {
	if n == 0 {
		// a^0 = 1
		return []int{1}
	} else if n == 1 {
		// a^1 = a
		return a
	} else {
		// start with a^1
		ans := a
		// multiply together
		for i := 2; i <= n; i++ {
			ans = PolyReduce(PolyMul(ans, a), base)
		}
		return PolyReduce(ans, base)
	}
}

func ChangeBase(a []int, from int, to int) []int {
	// changes a polynomial of base 'from' to base 'to'
	// original [1,2,3] = 1*from^0 + 2*from^1 + 3*from^2
	// so..

	// put 'from' into 'to'
	b := Int2Poly(from, to)
	c := []int{}
	for i, ai := range a {
		Ai := Int2Poly(ai, to)
		term := PolyPow(b, i, to)
		p := PolyMul(Ai, term)
		c = PolyAdd(c, p)
		c = PolyReduce(c, to)
	}

	for _, x := range c {
		if x < 0 {
			panic("change base verflow")
		}
	}
	return c
}

func PolyDivide(a []int, x int) {
	// divide a by x using integer division (ie, rounding to the nearest x)

}

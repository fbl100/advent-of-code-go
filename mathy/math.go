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

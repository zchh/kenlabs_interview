package s5

import "math"

func quickAdd(y, z, x int) bool {
	for result, add := 0, y; z > 0; z >>= 1 {
		if z&1 > 0 {
			if result < x-add {
				return false
			}
			result += add
		}
		if z != 1 {
			if add < x-add {
				return false
			}
			add += add
		}
	}
	return true
}

func Divide(a, b int) int {
	if a == math.MinInt32 {
		if b == 1 {
			return math.MinInt32
		}
		if b == -1 {
			return math.MaxInt32
		}
	}
	if b == math.MinInt32 {
		if a == math.MinInt32 {
			return 1
		}
		return 0
	}
	if a == 0 {
		return 0
	}

	rev := false
	if a > 0 {
		a = -a
		rev = !rev
	}
	if b > 0 {
		b = -b
		rev = !rev
	}

	ans := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1
		if quickAdd(b, mid, a) {
			ans = mid
			if mid == math.MaxInt32 {
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if rev {
		return -ans
	}
	return ans
}

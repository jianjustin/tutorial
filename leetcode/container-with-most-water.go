package leetcode

func maxArea(height []int) int {
	x, y, result := 0, 0, 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			x = j - i
			y = min(height[i], height[j])
			if result < x*y {
				result = x * y
			}
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

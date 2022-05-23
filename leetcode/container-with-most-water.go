package leetcode

func maxArea(height []int) int {
	x, y, result := 0, len(height)-1, 0

	for x < y {
		item := (y - x) * min(height[x], height[y])
		if item > result {
			result = item
		}

		if height[x] > height[y] {
			y--
		} else {
			x++
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

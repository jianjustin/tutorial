package leetcode

func lengthOfLongestSubstring(s string) int {
	res := make(map[byte]int)
	var count, j = 0, -1

	for i := 0; i < len(s); i++ {
		if i > 0 { //删除前面字符
			delete(res, s[i-1])
		}

		for j+1 < len(s) && res[s[j+1]] == 0 {
			res[s[j+1]] = 1
			j++
		}

		count = max(count, j-i+1)
	}

	return count
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

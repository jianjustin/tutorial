package leetcode

// 计算最终存的钱
func calculate_money(n int) int {
	var result = 0

	var p = n / 7
	var q = n % 7

	result += (28 + (28 + (p-1)*7)) * p / 2
	result += (q+1)*q/2 + p*q

	return result
}

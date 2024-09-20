// https://leetcode.com/problems/fibonacci-number

package MediumAlgo

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i < n; i++ {
		dp[i] = -1
	}
	return calf(n, dp)
}

func calf(n int, dp []int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	s := dp[n-1]
	if s == -1 {
		s = calf(n-1, dp) + calf(n-2, dp)
		dp[n-1] = s
	}
	return s
}

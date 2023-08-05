// https://leetcode.com/problems/climbing-stairs/

package EasyAlgoAndStrings

func climbStairs(n int) int {
    return fibonacci(n)
}

func fibonacci(n int) int {
    a := 1
    b := 1
    for i := 2; i <= n; i++ {
		c := a + b
        b = a
        a = c
	}
    return a
}

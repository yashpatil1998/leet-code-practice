// https://leetcode.com/problems/longest-harmonious-subsequence

package EasyAlgoAndStrings

func findLHS(nums []int) int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	max := 0
	for n, f := range freq {
		if freq[n+1] > 0 {
			m := f + freq[n+1]
			if m > max {
				max = m
			}
		}
	}
	return max
}

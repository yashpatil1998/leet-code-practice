// https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target

package MediumAlgo

import "sort"

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	start, end, ans := 0, len(nums)-1, 0
	for start < end {
		if (nums[start] + nums[end]) < target {
			ans = ans + end - start
			start = start + 1
		} else {
			end = end - 1
		}
	}
	return ans
}

// https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum

package EasyAlgoAndStrings

import "sort"

func maxSubsequence(nums []int, k int) []int {
	type pair struct {
		val int
		idx int
	}
	pq := make([]pair, 0)
	for i, v := range nums {
		pq = append(pq, pair{v, i})
		sort.Slice(pq, func(i, j int) bool {
			return pq[i].val < pq[j].val
		})
		if len(pq) > k {
			pq = pq[1:]
		}
	}
	sort.Slice(pq, func(i, j int) bool {
		return pq[i].idx < pq[j].idx
	})
	res := make([]int, k)
	for i, p := range pq {
		res[i] = nums[p.idx]
	}
	return res
}

// https://leetcode.com/problems/median-of-two-sorted-arrays/

package EasyAlgoAndStrings

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	for i != len(nums1) {
		merged = append(merged, nums1[i])
		i++
	}
	for j != len(nums2) {
		merged = append(merged, nums2[j])
		j++
	}
	if len(merged)%2 == 0 {
		middle := len(merged) / 2
		return float64(merged[middle]+merged[middle-1]) * 0.5
	}
	return float64(merged[len(merged)/2])
}

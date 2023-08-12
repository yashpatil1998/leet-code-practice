// https://leetcode.com/problems/search-in-rotated-sorted-array

package MediumAlgo

func search(nums []int, target int) int {
    start, end := 0, len(nums) - 1

    for start <= end {
        mid := (start + end) / 2

        if nums[mid] == target {
            return mid
        }

        if nums[start] <= nums[mid] { // meaning left half is sorted
            if nums[start] <= target && target < nums[mid] { // meaning target is in this half  
                end = mid - 1 // if target is here then ignore other half of array
            } else {
                start = mid + 1
            }
        } else { // right half is sorted
            if nums[mid] < target && target <= nums[end] {
                start = mid + 1
            } else {
                end = mid - 1
            }
        }
    }

    return -1
}

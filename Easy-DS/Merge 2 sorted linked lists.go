// https://leetcode.com/problems/merge-two-sorted-lists/

package EasyDS

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	resultNode := &ListNode{}
	resultNodeCopy := resultNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			resultNode.Next = list1
			list1 = list1.Next
		} else {
			resultNode.Next = list2
			list2 = list2.Next
		}
		resultNode = resultNode.Next
	}
	if list1 == nil {
		resultNode.Next = list2
	} else {
		resultNode.Next = list1
	}
	return resultNodeCopy.Next
}

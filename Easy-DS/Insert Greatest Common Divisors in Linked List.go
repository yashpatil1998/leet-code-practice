// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/

package EasyDS

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	head.Next = &ListNode{
		Val:  gcd(head.Val, head.Next.Val),
		Next: head.Next,
	}
	head.Next.Next = insertGreatestCommonDivisors(head.Next.Next)
	return head
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	if a == b {
		return a
	}

	if a > b {
		return gcd(a-b, b)
	}
	return gcd(a, b-a)
}

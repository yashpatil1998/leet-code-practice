// https://leetcode.com/problems/add-two-numbers

package MediumDS

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    str1, str2 := []int{}, []int{}
    for l1 != nil {
        str1 = append((str1), int(l1.Val))
        l1 = l1.Next
    }
    for l2 != nil {
        str2 = append((str2), int(l2.Val))
        l2 = l2.Next
    }
    if (len(str1) > len(str2)) {
        str2 = appendZeros(str2, len(str1) - len(str2))
    } else {
        str1 = appendZeros(str1, len(str2) - len(str1))
    }
    str1 = revSlice(str1)
    str2 = revSlice(str2)
    
    carry := 0
    var result []int
    for i := len(str1) - 1; i > -1; i-- {
        currentSum := str1[i] + str2[i] + carry
        result = append(result, currentSum % 10)
        if currentSum > 9 {
            carry = 1
        } else {
            carry = 0
        }

    }
    if carry == 1 {
        result = append(result, 1)
    }
    resultList := &ListNode{
        Val: result[0],
    }
    headResult := resultList
    for i := 1; i < len(result); i++ {
        resultList.Next = &ListNode{}
        resultList = resultList.Next
        resultList.Val = result[i]
    }
    resultList.Next = nil
    return headResult
}

func revSlice(slc []int) []int {
    rev_slc := []int{}
    for i := len(slc) - 1; i>= 0; i-- {
        rev_slc = append(rev_slc, slc[i])
    }
    return rev_slc
}

func appendZeros(slc []int, numOfZ int) []int {
    for i := 0; i < numOfZ; i++ {
        slc = append(slc, 0)
    }
    return slc
}

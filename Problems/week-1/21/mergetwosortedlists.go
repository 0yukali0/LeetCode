/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil || list2 == nil {
        if list1 != nil {
            return list1
        }
        if list2 != nil {
            return list2
        }
        return list1
    }

    var result, tail *ListNode
    if list1.Val <= list2.Val {
        result = list1
        list1 = list1.Next
    } else {
        result = list2
        list2 = list2.Next
    }
    tail = result
    tail.Next = nil

    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }
        tail = tail.Next
        tail.Next = nil
    }

    if list1 != nil {
        tail.Next = list1
    } else {
        tail.Next = list2
    }
    return result
}

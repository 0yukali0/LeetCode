/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    address := make(map[*ListNode]bool, 0)
    for ptr := head; ptr != nil; {
        if ok, _ := address[ptr]; ok {
            return true
        } else {
            address[ptr] = true
        }
        ptr = ptr.Next
    }
    return false
}

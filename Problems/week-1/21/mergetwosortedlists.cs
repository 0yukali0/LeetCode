/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution {
    public ListNode MergeTwoLists(ListNode list1, ListNode list2) {
        if (list1 == null || list2 == null) {
            if (list1 == null) {
                return list2;
            }

            if (list2 == null) {
                return list1;
            }
        }

        ListNode result;
        ListNode last;
        if (list1.val <= list2.val) {
            result = list1;
            list1 = list1.next;
        } else {
            result = list2;
            list2 = list2.next;
        }
        result.next = null;
        last = result;

        while (list1 != null && list2 != null) {
            if (list1.val <= list2.val) {
                last.next = list1;
                list1 = list1.next;
            } else {
                last.next = list2;
                list2 = list2.next;
            }
            last = last.next;
            last.next = null;
        }

        if (list1 != null) {
            last.next = list1;
        }

        if (list2 != null) {
            last.next = list2;
        }

        return result;
    }
}

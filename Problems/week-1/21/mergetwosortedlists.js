/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
var mergeTwoLists = function(list1, list2) {
    if (!list1 || !list2) {
        if (list1) {
            return list1;
        }
        return list2;
    }

    let result = list2;
    if (list1.val <= list2.val) {
        result = list1;
        list1 = list1.next;
    } else {
        list2 = list2.next;
    }
    
    let rs = result;
    while (list1 && list2) {
        if (list1.val <= list2.val) {
            rs.next = list1;
            list1 = list1.next;
        } else {
            rs.next = list2;
            list2 = list2.next;
        }
        rs = rs.next;
        rs.next = null;
    }

    if (list1) {
        rs.next = list1;
    }

    if (list2) {
        rs.next = list2;
    }
    return result;
};

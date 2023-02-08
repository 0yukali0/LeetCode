/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }

    waitforcheck := []*TreeNode{root}
    for i := 0; i < len(waitforcheck); i++ {
        p := waitforcheck[i]
        var tmp *TreeNode
        if p.Left != nil {
            tmp = p.Left
        }
        p.Left = p.Right
        p.Right = tmp

        if p.Left != nil {
            waitforcheck = append(waitforcheck, p.Left)
        }

        if p.Right != nil {
            waitforcheck = append(waitforcheck, p.Right)
        }
    }
    return root
}

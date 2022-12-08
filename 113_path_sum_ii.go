/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "fmt"

func pathSum(root *TreeNode, targetSum int) [][]int {
    var res [][]int
    var path []int
    dfs(root, targetSum, path, 0, &res)
    return res
}

func dfs(
    root *TreeNode,
    targetSum int,
    path []int,
    currSum int,
    res *[][] int){
    
    
    if root == nil {
        return
    }

    if root.Left == nil && root.Right == nil {
        if currSum + root.Val == targetSum {
            path = append(path, root.Val)
            *res = append(*res, path) // append main variable 
            return
        }
    } else {
        c_s := currSum + root.Val
        path = append(path, root.Val)
        dfs(root.Left, targetSum, path, c_s, res)
        dfs(root.Right, targetSum, path, c_s, res)
    }
}

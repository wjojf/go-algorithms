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
    dfs(root, targetSum, &path, 0, &res)
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
            f_p := append(path, root.Val)
            fmt.Println(f_p)
            *res = append(*res, f_p)
            fmt.Println(*res)
            return
        }
    }

    c_s := currSum + root.Val
    p := append(path, root.Val)

    dfs(root.Left, targetSum, p, c_s, res)
    dfs(root.Right, targetSum, p, c_s, res)
}

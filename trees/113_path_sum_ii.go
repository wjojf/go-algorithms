/**
Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.

A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.
 */
package main

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
    res *[][] int) {
    

    if root == nil {
        return
    }

    c_s := currSum + root.Val
    if root.Left == nil && root.Right == nil {
        if c_s == targetSum {
            tmp := make([]int, len(path))
            copy(tmp, path)
            *res = append(*res, append(tmp, root.Val))
            return
        }
    } else {
        dfs(root.Left, targetSum, append(path, root.Val), c_s, res)
        dfs(root.Right, targetSum, append(path, root.Val), c_s, res)
    }
}

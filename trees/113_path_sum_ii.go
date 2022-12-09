/**
Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.

A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.
 */
package main

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
<<<<<<< HEAD:trees/113_path_sum_ii.go
            f_p := append(path, root.Val)
            fmt.Println(f_p)
            *res = append(*res, f_p)
            fmt.Println(*res)
=======
            path = append(path, root.Val)
            *res = append(*res, path) // append main variable 
>>>>>>> f81ce9522cd105418b20db316186f6b8c9674095:113_path_sum_ii.go
            return
        }
    } else {
        c_s := currSum + root.Val
        path = append(path, root.Val)
        dfs(root.Left, targetSum, path, c_s, res)
        dfs(root.Right, targetSum, path, c_s, res)
    }
<<<<<<< HEAD:trees/113_path_sum_ii.go

    c_s := currSum + root.Val
    p := append(path, root.Val)

    dfs(root.Left, targetSum, p, c_s, res)
    dfs(root.Right, targetSum, p, c_s, res)
=======
>>>>>>> f81ce9522cd105418b20db316186f6b8c9674095:113_path_sum_ii.go
}

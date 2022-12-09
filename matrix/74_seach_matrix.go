package main

func searchMatrix(matrix [][]int, target int) bool {
    if target < matrix[0][0] || target > matrix[len(matrix) -1][len(matrix[0]) -1]{
        return false 
    }

    for _, row := range(matrix) {
        if target > row[len(row)-1] {
            continue
        }
        for _, x := range(row) {
            if target == x {
                return true
            }
        }
    }
    return false
}

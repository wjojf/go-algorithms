package main

func isValidSudoku(board [][]byte) bool {
    for i := 0; i < 9; i++{
        for j := 0; j < 9; j++{
            if string(board[i][j]) == "." {
                continue 
            }
            v := check(board, i, j)
            if !v {
                return false
            }
        }
    }
    return true
}

func check(board [][]byte, i,j int) bool{
    // check row
    c := board[i][j]

    for in, el := range(board[i]) {
        if el == c && in != j {
            return false
        }
    }
    
    // check col 
    for r := 0; r < 9; r++ {
        if board[r][j] == c{
            if r != i {
                return false
            }    
        }
    }

    // check square 
    _i := int(i / 3)
    _j := int(j / 3)

    for x := _i * 3; x < (_i + 1) * 3; x++ {
        for y := _j * 3; y < (_j + 1) * 3; y++ {
            if board[x][y] == c && ((x != i) || (y != j)){
                return false
            }
        }
    }

    return true
}

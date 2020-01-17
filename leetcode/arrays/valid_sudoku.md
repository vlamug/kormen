Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

A partially filled sudoku which is valid.

The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

Example 1:
```
Input:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: true
```

Example 2:
```
Input:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being 
    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
```

Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
The given board contain only digits 1-9 and the character '.'.
The given board size is always 9x9.

Solution:

```go
package main

import (
	"fmt"
)

func main() {
	sudokuStr := []string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}
	
	sudokyByte := [][]byte{}
	for i := 0; i < 9; i++ {
		sudokyByte = append(sudokyByte, []byte(sudokuStr[i]))
	}
	
	fmt.Println(isValidSudoku(sudokyByte))
}

func isValidSudoku(board [][]byte) bool {
	var empty struct{}
	
	for i := 0; i < 9; i++ {
		horizontCheckMap := make(map[byte]struct{}, 9)
		verticalCheckMap := make(map[byte]struct{}, 9)
		matrixCheckMap := make(map[byte]struct{}, 9)
		
		for j := 0; j < 9; j++ {
			if board[i][j] != 46 {
				if _, ok := horizontCheckMap[board[i][j]]; ok {
					return false
				}		
				horizontCheckMap[board[i][j]] = empty
			}
			if board[j][i] != 46 {
				if _, ok := verticalCheckMap[board[j][i]]; ok {
					return false
				}		
				verticalCheckMap[board[j][i]] = empty
			}
			
			ii := (i/3)*3 + j/3
			jj := (i%3)*3 + j%3
			if board[ii][jj] != 46 {
				if _, ok := matrixCheckMap[board[ii][jj]]; ok {
					return false
				}		
				matrixCheckMap[board[ii][jj]] = empty
			}	
		}
	}
	return true
}
```

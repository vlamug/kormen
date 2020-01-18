You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Note:

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:

Given input matrix = 
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

Example 2:

Given input matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
], 

rotate the input matrix in-place such that it becomes:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]

Solution:

```go
func rotate(matrix [][]int)  {
	n := len(matrix)-1

	for i := 0; i <= n/2; i++ {
		for j := 0; j < n-2*i; j++ {
			rightII := j+i
			rightJJ := n-i
						
			rightBuf := matrix[rightII][rightJJ]
			matrix[rightII][rightJJ] = matrix[i][j+i]
			
			bottomII := n-i
			bottomJJ := n-j-i
			
			bottomBuf := matrix[bottomII][bottomJJ]
			matrix[bottomII][bottomJJ] = rightBuf
			
			leftII := n-j-i
			leftJJ := i
			
			leftBuf := matrix[leftII][leftJJ]
			matrix[leftII][leftJJ] = bottomBuf
			
			matrix[i][j+i] = leftBuf
		}
	}
}
```

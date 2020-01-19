Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

Solution:

```go
func reverse(x int) int {

    isNegative := false
    if x < 0 {
 	isNegative = true
	x = -1 * x
    }

    res := x%10
    x = x/10
    for x != 0 {
        mod := x%10
        x = x/10

	res = res*10 + mod
	
	if res > 2147483647 {
	    return 0
	}
    }

    if isNegative {
        res = -1 * res
    }

    return res
}
```
